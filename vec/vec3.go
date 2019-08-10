package vec

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// ZeroVec3 is the vector at origo
var ZeroV3 = Vec3{0.0, 0.0, 0.0}

// New creates a new vector
func V3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (u Vec3) String() string {
	return fmt.Sprintf("vec.Vec3(%v, %v, %v)", u.X, u.Y, u.Z)
}

// Scale returns vector u scaled by s
func (u Vec3) Scale(s float64) Vec3 {
	return Vec3{
		u.X * s,
		u.Y * s,
		u.Z * s,
	}
}

// Dot returns dot product of vectors u and v
func (u Vec3) Dot(v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// Len returns length of vector u
func (u Vec3) Len() float64 {
	return math.Sqrt(u.Dot(u))
}

// Add returns the sum of vectors u and v
func (u Vec3) Add(v Vec3) Vec3 {
	return Vec3{
		u.X + v.X,
		u.Y + v.Y,
		u.Z + v.Z,
	}
}

// Sub returns the difference betweeen vectors u and v
func (u Vec3) Sub(v Vec3) Vec3 {
	return Vec3{
		u.X - v.X,
		u.Y - v.Y,
		u.Z - v.Z,
	}
}

// Dist returns the distance between vectors u and v
func (u Vec3) Dist(v Vec3) float64 {
	return u.Sub(v).Len()
}

// Unit returns the vector u as a unit vector
func (u Vec3) Unit() Vec3 {
	return u.Scale(1 / u.Len())
}

// Eq checks equality of vectors u and v
func (u Vec3) Eq(v Vec3) bool {
	return u.X == v.X && u.Y == v.Y && u.Z == v.Z
}
