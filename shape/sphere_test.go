package shape

import (
	"testing"

	"github.com/heldtalex/ray-tracer/vec"
)

func TestSphereDistanceToPoint(t *testing.T) {
	for _, tc := range []struct {
		v    vec.Vec3
		s    Sphere
		want float64
	}{
		{vec.V3(1, 0, 0), NewSphere(vec.V3(0, 0, 0), 1), 0},
		{vec.V3(0, 0, 0.5), NewSphere(vec.V3(0, 0, 0), 1), -0.5},
		{vec.V3(0, 3, 0), NewSphere(vec.V3(0, 0, 0), 1), 2},
	} {
		if got := tc.s.DistanceToPoint(tc.v); got != tc.want {
			t.Fatalf("%v.DistanceToPoint(%v) = %v, want %v", tc.s, tc.v, got, tc.want)
		}
	}

}
