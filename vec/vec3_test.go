package vec

import "testing"

func TestV3(t *testing.T) {
	x, y, z := 1.1, 2.2, 3.3

	want := Vec3{X: 1.1, Y: 2.2, Z: 3.3}

	if got := V3(x, y, z); !got.Eq(want) {
		t.Fatalf("V3(%v, %v, %v) = %v, want %v", x, y, z, got, want)
	}
}

func TestVec3Scale(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		s    float64
		want Vec3
	}{
		{V3(1, 2, 3), 0, V3(0, 0, 0)},
		{V3(1, 2, 3), 0.5, V3(0.5, 1, 1.5)},
		{V3(1, 2, 3), 2, V3(2, 4, 6)},
	} {
		if got := tc.u.Scale(tc.s); !got.Eq(tc.want) {
			t.Fatalf("%v.Scale(%v) = %v, want %v", tc.u, tc.s, got, tc.want)
		}
	}
}

func TestVec3Dot(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		v    Vec3
		want float64
	}{
		{V3(1, 1, 1), V3(0, 0, 0), 0},
		{V3(1, 1, 1), V3(1, 1, 1), 3},
		{V3(1, 2, 3), V3(1, 1, 1), 6},
	} {
		if got := tc.u.Dot(tc.v); got != tc.want {
			t.Fatalf("%v.Dot(%v) = %v, want %v", tc.u, tc.v, got, tc.want)
		}
	}
}

func TestVec3Len(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		want float64
	}{
		{V3(1, 0, 0), 1},
		{V3(2, 0, 0), 2},
		{V3(1, 1, 0), 1.4142135623730951},
		{V3(1, 1, 1), 1.7320508075688772},
		{V3(1, 2, 3), 3.7416573867739413},
	} {
		if got := tc.u.Len(); got != tc.want {
			t.Fatalf("%v.Len() = %v, want %v", tc.u, got, tc.want)
		}
	}
}

func TestVec3Add(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		v    Vec3
		want Vec3
	}{
		{V3(0, 0, 0), V3(0, 0, 0), V3(0, 0, 0)},
		{V3(1, 0, 0), V3(1, 0, 0), V3(2, 0, 0)},
		{V3(1, 1, 1), V3(1, 1, 1), V3(2, 2, 2)},
		{V3(1, 2, 3), V3(1, 2, 3), V3(2, 4, 6)},
	} {
		if got := tc.u.Add(tc.v); got != tc.want {
			t.Fatalf("%v.Add(%v) = %v, want %v", tc.u, tc.v, got, tc.want)
		}
	}
}

func TestVec3Sub(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		v    Vec3
		want Vec3
	}{
		{V3(0, 0, 0), V3(0, 0, 0), V3(0, 0, 0)},
		{V3(1, 0, 0), V3(1, 0, 0), V3(0, 0, 0)},
		{V3(1, 1, 1), V3(1, 1, 1), V3(0, 0, 0)},
		{V3(2, 4, 6), V3(1, 2, 3), V3(1, 2, 3)},
	} {
		if got := tc.u.Sub(tc.v); got != tc.want {
			t.Fatalf("%v.Sub(%v) = %v, want %v", tc.u, tc.v, got, tc.want)
		}
	}
}

func TestVec3Dist(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		v    Vec3
		want float64
	}{
		{V3(0, 0, 0), V3(0, 0, 0), 0},
		{V3(2, 3, 5), V3(2, 0, 9), 5},
		{V3(1, 0, 5), V3(0, 2, 4), 2.449489742783178},
	} {
		if got := tc.u.Dist(tc.v); got != tc.want {
			t.Fatalf("%v.Dist(%v) = %v, want %v", tc.u, tc.v, got, tc.want)
		}
	}
}

func TestVec3Unit(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		want Vec3
	}{
		{V3(1, 0, 0), V3(1, 0, 0)},
		{V3(1, 1, 0), V3(0.7071067811865475, 0.7071067811865475, 0)},
		{V3(1, 1, 1), V3(0.5773502691896258, 0.5773502691896258, 0.5773502691896258)},
		{V3(1, 2, 3), V3(0.2672612419124244, 0.5345224838248488, 0.8017837257372732)},
	} {
		if got := tc.u.Unit(); !got.Eq(tc.want) {
			t.Fatalf("%v.Normalize() = %v, want %v", tc.u, got, tc.want)
		}
	}
}

func TestVec3Eq(t *testing.T) {
	for _, tc := range []struct {
		u    Vec3
		v    Vec3
		want bool
	}{
		{V3(1, 2, 3), V3(1, 2, 3), true},
		{V3(1, 2, 3), V3(4, 5, 6), false},
	} {
		if got := tc.u.Eq(tc.v); got != tc.want {
			t.Fatalf("u.Eq(%v)\ngot %v, want %v", tc.v, got, tc.want)
		}
	}
}
