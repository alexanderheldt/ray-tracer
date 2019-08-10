package sdf

import (
	"github.com/heldtalex/ray-tracer/shape"
	"github.com/heldtalex/ray-tracer/vec"
)

// Sphere returns the distane from a point p to the surface
// of the sphere s
func Sphere(p vec.Vec3, s shape.Sphere) float64 {
	return s.Position.Dist(p) - s.Radius
}
