package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"

	"github.com/heldtalex/ray-tracer/camera"
	"github.com/heldtalex/ray-tracer/shape"
	"github.com/heldtalex/ray-tracer/vec"
)

var (
	MAX_STEPS         = 1000
	MAX_DISTANCE      = 1000.0
	MIN_HIT_DISTANCE  = 0.001
	AASamplesPerPixel = 100
)

type Scene struct {
	Lights []vec.Vec3
	Shapes []shape.Shape
}

func main() {
	height := 256
	width := height * 2
	aspectRatio := float64(width) / float64(height)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	cam := camera.New(vec.V3(0, 1, 0), -1, 45)

	scene := Scene{
		Lights: []vec.Vec3{
			vec.V3(-2, 5, -3),
		},
		Shapes: []shape.Shape{
			shape.NewSphere(vec.V3(0, 1, -6), 1),
			shape.NewPlane(vec.ZeroV3, vec.V3(0, 1, 0)),
		},
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Create a color c that we will blend each pixel sample into to
			// create antialiasing
			c := vec.ZeroV3

			// Use a WaitGroup to run multiple gorutines for each pixel to parallelize
			// calculations
			wg := &sync.WaitGroup{}

			// Use a mutex to ensure that our color c is calculated correctly
			m := &sync.Mutex{}

			for s := 0; s < AASamplesPerPixel; s++ {
				// Converting pixels positions from
				// "raster space": (0, 0) = top left, (width, height) = bottom right to
				// "camera space": (-1, 1) = top left, (1, -1) = bottom right

				// Normalize pixels with image dimensions to NDC (Normalized Device Coordinates) space.
				// Add offset with a random value for the antialiasing
				ndcX := (float64(x) + rand.Float64()) / float64(width)
				ndcY := (float64(y) + rand.Float64()) / float64(height)

				// NOTE: NDC coordinate range is [0, 1] but we want to remap
				// them to "screen space" which is in the range [-1, 1]
				screenX := (2.0 * ndcX) - 1
				screenY := 1 - (2.0 * ndcY)

				//Finally, calculate the ray going through screen pixel x and y
				ray := cam.Ray(screenX, screenY, aspectRatio)

				wg.Add(1)
				go func(wg *sync.WaitGroup, m *sync.Mutex) {
					defer wg.Done()

					hitPoint := rayMarch(ray, scene)

					m.Lock()
					defer m.Unlock()

					c = c.Add(hitPoint)
				}(wg, m)
			}

			wg.Wait()

			// Normalize color
			c = c.Scale(1 / float64(AASamplesPerPixel))

			r := uint8(255.99 * c.X)
			g := uint8(255.99 * c.Y)
			b := uint8(255.99 * c.Z)

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

func rayMarch(ray camera.Ray, scene Scene) vec.Vec3 {
	distanceTraveled := 0.0

	for i := 0; i < MAX_STEPS; i++ {
		currentPosition := ray.At(distanceTraveled)
		closestDistance := closestDistanceFromPointToScene(currentPosition, scene)

		// Nothing is in the path of the ray
		if closestDistance == MAX_DISTANCE {
			return vec.ZeroV3
		}

		// We are close enough to something to call it a hit
		if closestDistance < MIN_HIT_DISTANCE {
			return calculateIntersect(currentPosition, scene).Color
		}

		// No hit, continue marching
		distanceTraveled += closestDistance
	}

	return vec.ZeroV3
}

func closestDistanceFromPointToScene(p vec.Vec3, scene Scene) float64 {
	closest := MAX_DISTANCE

	for _, s := range scene.Shapes {
		if distance := s.DistanceToPoint(p); distance < closest {
			closest = distance
		}
	}

	return closest
}

func calculateIntersect(p vec.Vec3, scene Scene) shape.Hit {
	closestHit := shape.Hit{
		Distance: MAX_DISTANCE,
		Color:    vec.ZeroV3,
	}

	var closestShape shape.Shape

	for _, s := range scene.Shapes {
		if hit := s.Hit(p); hit.Distance < closestHit.Distance {
			closestHit = hit
			closestShape = s
		}
	}

	if closestHit.Distance < MAX_DISTANCE {
		n := estimateSurfaceNormal(p, closestShape)

		totalLightIntensity := 0.0
		for _, l := range scene.Lights {
			lightDirection := l.Sub(p).Unit()

			// Move the origin of the ray slightly in the direction of the
			// light source. This is to avoid the distance calculation to
			// not trigger immediatly and think we're hitting the closest
			// object (`closestShape`)
			shadowOrigin := p.Add(lightDirection.Scale(0.01))

			// Create the shadow ray...
			shadowRay := camera.Ray{
				Origin:    shadowOrigin,
				Direction: lightDirection,
			}

			// ... and march it towards the light
			shadow := calculateShadow(shadowRay, scene)

			lightIntensity := n.Dot(lightDirection) * shadow
			totalLightIntensity += clamp(lightIntensity, 0, 1)
		}

		closestHit.Color = closestHit.Color.Scale(clamp(totalLightIntensity, 0, 1))
	}

	return closestHit
}

func estimateSurfaceNormal(p vec.Vec3, s shape.Shape) vec.Vec3 {
	epsilon := 0.1

	// Calculate the gradient at point p to estimate the surface normal
	nx := s.DistanceToPoint(p.Add(vec.V3(epsilon, 0, 0))) - s.DistanceToPoint(p.Sub(vec.V3(epsilon, 0, 0)))
	ny := s.DistanceToPoint(p.Add(vec.V3(0, epsilon, 0))) - s.DistanceToPoint(p.Sub(vec.V3(0, epsilon, 0)))
	nz := s.DistanceToPoint(p.Add(vec.V3(0, 0, epsilon))) - s.DistanceToPoint(p.Sub(vec.V3(0, 0, epsilon)))

	return vec.V3(nx, ny, nz).Unit()
}

func calculateShadow(ray camera.Ray, scene Scene) float64 {
	// shadow of 0 means total shadow, shadow of 1 means no shadow
	shadow := 1.0

	// sunSize controls how hard/soft the shadows will be
	// http://www.iquilezles.org/www/articles/rmshadows/rmshadows.htm
	sunSize := 8.0

	distanceTraveled := 0.0
	for i := 0; i < MAX_STEPS; i++ {
		currentPosition := ray.At(distanceTraveled)
		closestDistance := closestDistanceFromPointToScene(currentPosition, scene)

		// Something was in the way; the point is in shadow
		if closestDistance < MIN_HIT_DISTANCE {
			return 0.0
		}

		// We didn't hit anything, return how much shadow the
		// point is in
		if closestDistance == MAX_DISTANCE {
			return clamp(shadow, 0.0, 1.0)
		}

		shadow = math.Min(shadow, (closestDistance*sunSize)/distanceTraveled)
		distanceTraveled += closestDistance
	}

	return 1.0
}

// clamp returns x clamped to the range [min, max]
// If x is less than min, min is returned. If x is more than max, max is returned. Otherwise, x is
// returned.
func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}
