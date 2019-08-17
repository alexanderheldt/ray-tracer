package shape

import (
	"fmt"

	"github.com/heldtalex/ray-tracer/vec"
)

type Plane struct {
	Position vec.Vec3
	Normal   vec.Vec3
}

func NewPlane(p, n vec.Vec3) Plane {
	return Plane{p, n}
}

func (p Plane) String() string {
	return fmt.Sprintf("shape.Plane(%v, %v)", p.Position, p.Normal)
}

func (p Plane) DistanceToPoint(v vec.Vec3) float64 {
	return p.Position.Add(v).Dot(p.Normal)
}

func (p Plane) Hit(v vec.Vec3) Hit {
	return Hit{
		Distance: p.DistanceToPoint(v),
		Color:    vec.V3(1, 0, 1),
	}
}
