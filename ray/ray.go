package ray

import "github.com/heldtalex/ray-tracer/vec"

type Ray struct {
	Origin    vec.Vec3
	Direction vec.Vec3
}

func New(o, d vec.Vec3) Ray {
	return Ray{o, d}
}

func (r Ray) At(distance float64) vec.Vec3 {
	return r.Origin.Add(r.Direction.Scale(distance))
}
