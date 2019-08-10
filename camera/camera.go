package camera

import (
	"math"

	"github.com/heldtalex/ray-tracer/vec"
)

type Camera struct {
	Position vec.Vec3
	LookAt   float64
	FOV      float64

	AngleToScreen float64
}

func New(p vec.Vec3, la, fov float64) Camera {
	angle := math.Tan(math.Pi * (0.5 * fov) / 180)

	return Camera{p, la, fov, angle}
}
