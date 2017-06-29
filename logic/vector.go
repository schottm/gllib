package logic

import "math"

//Vector2f//

type Vector2f struct {

	X, Y float32
}

func (vec *Vector2f) Length() float32 {

	return float32(math.Sqrt(float64(vec.X * vec.X + vec.Y * vec.Y)))
}

func (vec *Vector2f) Normalize() {

	l := vec.Length()
	vec.X /= l
	vec.Y /= l
}

func (vec *Vector2f) Dot(in *Vector2f) float32 {

	return vec.X * in.X + vec.Y * in.Y
}

func (vec *Vector2f) Vector3f(z float32) *Vector3f {

	return &Vector3f{vec.X, vec.Y, z}
}

func (vec *Vector2f) Vector4f(z, w float32) *Vector4f {

	return &Vector4f{vec.X, vec.Y, z, w}
}
//Vector3f

type Vector3f struct {

	X, Y, Z float32
}

func (vec *Vector3f) Length() float32 {

	return float32(math.Sqrt(float64(vec.X * vec.X + vec.Y * vec.Y + vec.Z * vec.Z)))
}

func (vec *Vector3f) Normalize() {

	l := vec.Length()
	vec.X /= l
	vec.Y /= l
	vec.Z /= l
}

func (vec *Vector3f) Dot(in *Vector3f) float32 {

	return vec.X * in.X + vec.Y * in.Y + vec.Z * in.Z
}

func (vec *Vector3f) Cross(in *Vector3f) *Vector3f {

	x := vec.Y * in.Z - vec.Z * in.Y
	y := vec.Z * in.X - vec.X * in.Z
	z := vec.X * in.Y - vec.Y * in.X

	return &Vector3f{x, y, z}
}

func (vec *Vector3f) Vector4f(w float32) *Vector4f {

	return &Vector4f{vec.X, vec.Y, vec.Z, w}
}

//Vector4f

type Vector4f struct {

	X, Y, Z, W float32
}

func (vec *Vector4f) Length() float32 {

	return float32(math.Sqrt(float64(vec.X * vec.X + vec.Y * vec.Y + vec.Z * vec.Z + vec.W * vec.W)))
}

func (vec *Vector4f) Normalize() {

	l := vec.Length()
	vec.X /= l
	vec.Y /= l
	vec.Z /= l
	vec.W /= l
}

func (vec *Vector4f) Dot(in *Vector4f) float32 {

	return vec.X * in.X + vec.Y * in.Y + vec.Z * in.Z + vec.W * in.W
}
