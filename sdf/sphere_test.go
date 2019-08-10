package sdf

import (
	"testing"

	"github.com/heldtalex/ray-tracer/shape"
	"github.com/heldtalex/ray-tracer/vec"
)

func TestSphere(t *testing.T) {
	for _, tc := range []struct {
		p    vec.Vec3
		s    shape.Sphere
		want float64
	}{
		{vec.V3(1, 0, 0), shape.NewSphere(vec.V3(0, 0, 0), 1), 0},
		{vec.V3(0, 0, 0.5), shape.NewSphere(vec.V3(0, 0, 0), 1), -0.5},
		{vec.V3(0, 3, 0), shape.NewSphere(vec.V3(0, 0, 0), 1), 2},
	} {
		if got := Sphere(tc.p, tc.s); got != tc.want {
			t.Fatalf("sdf.Sphere(%v, %v) = %v, want %v", tc.p, tc.s, got, tc.want)
		}
	}

}
