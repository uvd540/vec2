package vec2

import "math"

type Vec2 struct {
	X float32
	Y float32
}

func New(x, y float32) Vec2 {
	return Vec2{x, y}
}

func Zero() Vec2 {
	return Vec2{0, 0}
}

func One() Vec2 {
	return Vec2{1, 1}
}

func (v *Vec2) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y)))
}

func (v *Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v *Vec2) Add(other *Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Sub(other *Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

func Add(v1, v2 Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

func Sub(v1, v2 Vec2) Vec2 {
	return Vec2{v1.X - v2.X, v1.Y - v2.Y}
}
