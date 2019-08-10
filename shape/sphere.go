package shape

import (
	"fmt"

	"github.com/heldtalex/ray-tracer/vec"
)

type Sphere struct {
	Position vec.Vec3
	Radius   float64
}

func NewSphere(p vec.Vec3, r float64) Sphere {
	return Sphere{p, r}
}

func (s Sphere) String() string {
	return fmt.Sprintf("shape.Sphere(%v, %v)", s.Position, s.Radius)
}
