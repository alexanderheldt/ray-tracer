package camera

import (
	"math"

	"github.com/heldtalex/ray-tracer/vec"
)

type Camera struct {
	Position vec.Vec3
	LookAt   float64
	FOV      float64
}

func New(p vec.Vec3, la, fov float64) Camera {
	return Camera{p, la, fov}
}

// Ray returns a ray going through screen pixel x and y
func (c Camera) Ray(x, y, aspectRatio float64) Ray {
	angleToScreen := math.Tan(math.Pi * (0.5 * c.FOV) / 180)

	cameraX := x * angleToScreen * aspectRatio
	cameraY := y * angleToScreen

	direction := vec.V3(cameraX, cameraY, c.LookAt).Unit()

	return Ray{
		Origin:    c.Position,
		Direction: direction,
	}
}
