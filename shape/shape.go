package shape

import "github.com/heldtalex/ray-tracer/vec"

type Shape interface {
	// DistanceToPoint returns the distance from the
	// surface of the Shape to vector v
	DistanceToPoint(v vec.Vec3) float64

	// Hit returns a Hit of the Shape at vector v
	Hit(v vec.Vec3) Hit
}

type Hit struct {
	// Distance from a vector to the hit point
	Distance float64

	// Color of the hit shape
	Color vec.Vec3
}
