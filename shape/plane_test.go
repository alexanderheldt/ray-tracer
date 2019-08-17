package shape

import (
	"testing"

	"github.com/heldtalex/ray-tracer/vec"
)

func TestPlaneDistanceToPoint(t *testing.T) {
	for _, tc := range []struct {
		v    vec.Vec3
		p    Plane
		want float64
	}{
		{vec.V3(1, 0, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(0, 1, 0)), 0},
		{vec.V3(0, 1, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(0, 1, 0)), 1},
		{vec.V3(0, 2, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(0, 1, 0)), 2},
		{vec.V3(0, 0, 1), NewPlane(vec.V3(0, 0, 0), vec.V3(0, 1, 0)), 0},
		{vec.V3(1, 0, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(1, 0, 0)), 1},
		{vec.V3(2, 0, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(1, 0, 0)), 2},
		{vec.V3(0, 1, 0), NewPlane(vec.V3(0, 0, 0), vec.V3(1, 0, 0)), 0},
		{vec.V3(0, 0, 1), NewPlane(vec.V3(0, 0, 0), vec.V3(1, 0, 0)), 0},
	} {
		if got := tc.p.DistanceToPoint(tc.v); got != tc.want {
			t.Fatalf("%v.DistanceToPoint(%v) = %v, want %v", tc.p, tc.v, got, tc.want)
		}
	}

}
