package camera

import "github.com/heldtalex/ray-tracer/vec"

type Ray struct {
	Origin    vec.Vec3
	Direction vec.Vec3
}

func (r Ray) At(distance float64) vec.Vec3 {
	return r.Origin.Add(r.Direction.Scale(distance))
}
