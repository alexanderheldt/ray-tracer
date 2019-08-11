package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/heldtalex/ray-tracer/camera"
	"github.com/heldtalex/ray-tracer/sdf"
	"github.com/heldtalex/ray-tracer/shape"
	"github.com/heldtalex/ray-tracer/vec"
)

var (
	MAX_STEPS        = 1000
	MAX_DISTANCE     = 1000.0
	MIN_HIT_DISTANCE = 0.001
)

// sphereSurfaceNormal estimates the surface normal of a Sphere at point p
func sphereSurfaceNormal(p vec.Vec3, s shape.Sphere) vec.Vec3 {
	epsilon := 0.001

	gradientX := sdf.Sphere(vec.V3(p.X+epsilon, p.Y, p.Z), s) - sdf.Sphere(vec.V3(p.X-epsilon, p.Y, p.Z), s)
	gradientY := sdf.Sphere(vec.V3(p.X, p.Y+epsilon, p.Z), s) - sdf.Sphere(vec.V3(p.X, p.Y-epsilon, p.Z), s)
	gradientZ := sdf.Sphere(vec.V3(p.X, p.Y, p.Z+epsilon), s) - sdf.Sphere(vec.V3(p.X, p.Y, p.Z-epsilon), s)

	return vec.V3(gradientX, gradientY, gradientZ).Unit()
}

func rayMarch(origin, direction vec.Vec3, light vec.Vec3, sphere shape.Sphere) vec.Vec3 {
	distanceFromOrigin := 0.0

	for i := 0; i < MAX_STEPS; i++ {
		currentPosition := origin.Add(direction.Scale(distanceFromOrigin))

		sphereDistance := sdf.Sphere(currentPosition, sphere)
		planeDistance := currentPosition.Y

		closest := math.Min(sphereDistance, planeDistance)

		// Hit Sphere
		if sphereDistance < planeDistance && sphereDistance < MIN_HIT_DISTANCE {
			n := sphereSurfaceNormal(currentPosition, sphere)

			lightDirection := light.Sub(currentPosition).Unit()
			diffuseIntensity := math.Max(0.0, n.Dot(lightDirection))

			return vec.V3(0, 0, 1).Scale(diffuseIntensity)
		}

		// Hit ground plane
		if planeDistance < sphereDistance && planeDistance < MIN_HIT_DISTANCE {
			return vec.V3(1, 0, 1)
		}

		// No hit
		if distanceFromOrigin > MAX_DISTANCE {
			return vec.ZeroV3
		}

		distanceFromOrigin += closest
	}

	return vec.ZeroV3
}

func main() {
	width := 256
	height := width
	aspectRatio := float64(width) / float64(height)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	cam := camera.New(vec.V3(0, 1, 0), -1, 45)

	sphere := shape.NewSphere(vec.V3(0, 1, -6), 1)
	light := vec.V3(-2, 5, -3)

	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			// Converting pixels positions from
			// "raster space": (0, 0) = top left, (width, height) = bottom right to
			// "camera space": (-1, 1) = top left, (1, -1) = bottom right

			// Normalize pixels with image dimensions to NDC (Normalized Device Coordinates) space.
			// Add offset of 0.5 to hit the pixel in the middle
			ndcX := (float64(x) + 0.5) / float64(width)
			ndcY := (float64(y) + 0.5) / float64(height)

			// NOTE: NDC coordinate range is [0, 1] but we want to remap
			// them to "screen space which" is in the range [-1, 1]
			screenX := (2.0 * ndcX) - 1
			screenY := 1 - (2.0 * ndcY)

			// Finally take the image aspect ratio and angle of the
			// cameras FOV (angle to the image plane) so we have coordinates
			// in camera space
			cameraX := screenX * cam.AngleToScreen * aspectRatio
			cameraY := screenY * cam.AngleToScreen

			rayDirection := vec.V3(cameraX, cameraY, cam.LookAt).Unit()

			hitPoint := rayMarch(cam.Position, rayDirection, light, sphere)

			r := uint8(255.99 * hitPoint.X)
			g := uint8(255.99 * hitPoint.Y)
			b := uint8(255.99 * hitPoint.Z)

			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
