package calc

import (
	"unsafe"
)

//Matrix2f//

type Matrix2f struct {
	M00, M01 float32
	M10, M11 float32
}

func NewMatrix2f() *Matrix2f {

	return &Matrix2f{}
}

func NewScaleMatrix2f(in *Vector2f) *Matrix2f {

	return &Matrix2f{in.X, 0.0, 0.0, in.Y}
}

func (mat *Matrix2f) Identity() {

	mat.M00 = 1.0; mat.M01 = 0.0
	mat.M10 = 0.0; mat.M11 = 1.0
}

func (mat *Matrix2f) GetData() *[4]float32 {

	return (*[4]float32)(unsafe.Pointer(&mat.M00))
}

func (mat *Matrix2f) Mul(in *Matrix2f) *Matrix2f {

	m00 := mat.M00* in.M00 + mat.M01* in.M10
	m01 := mat.M00* in.M01 + mat.M01* in.M11
	m10 := mat.M10* in.M00 + mat.M11* in.M10
	m11 := mat.M10* in.M01 + mat.M11* in.M11

	return &Matrix2f{m00, m01, m10, m11}
}

func (mat *Matrix2f) MulV(in *Vector2f) *Vector2f {

	x := mat.M00 * in.X + mat.M01 * in.Y
	y := mat.M10 * in.X + mat.M11 * in.Y

	return &Vector2f{x, y}
}

//Matrix4f//

type Matrix4f struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

func NewMatrix4f() *Matrix4f {

	return &Matrix4f{}
}

func NewScaleMatrix4f(scaleX, scaleY, scaleZ, scaleW float32) *Matrix4f {

	return &Matrix4f{
		scaleX, 0, 0, 0,
		0, scaleY, 0, 0,
		0, 0, scaleZ, 0,
		0, 0, 0, scaleW}
}

func NewTransformMatrix4f(in *Vector3f) *Matrix4f {

	return &Matrix4f{
		1, 0, 0, in.X,
		0, 1, 0, in.Y,
		0, 0, 1, in.Z,
		0, 0, 0, 1}
}

func (mat *Matrix4f) Identity() {

	mat.M00 = 1.0; mat.M01 = 0.0; mat.M02 = 0.0; mat.M03 = 0.0
	mat.M10 = 0.0; mat.M11 = 1.0; mat.M12 = 0.0; mat.M13 = 0.0
	mat.M20 = 0.0; mat.M21 = 0.0; mat.M22 = 1.0; mat.M23 = 0.0
	mat.M30 = 0.0; mat.M31 = 0.0; mat.M32 = 0.0; mat.M33 = 1.0
}

func (mat *Matrix4f) GetData() *[16]float32 {

	return (*[16]float32)(unsafe.Pointer(&mat.M00))
}

func (mat *Matrix4f) Mul(in *Matrix4f) *Matrix4f{

	m00 := mat.M00* in.M00 + mat.M01* in.M10 + mat.M02* in.M20 + mat.M03* in.M30
	m01 := mat.M00* in.M01 + mat.M01* in.M11 + mat.M02* in.M21 + mat.M03* in.M31
	m02 := mat.M00* in.M02 + mat.M01* in.M12 + mat.M02* in.M22 + mat.M03* in.M32
	m03 := mat.M00* in.M03 + mat.M01* in.M13 + mat.M02* in.M23 + mat.M03* in.M33
	m10 := mat.M10* in.M00 + mat.M11* in.M10 + mat.M12* in.M20 + mat.M13* in.M30
	m11 := mat.M10* in.M01 + mat.M11* in.M11 + mat.M12* in.M21 + mat.M13* in.M31
	m12 := mat.M10* in.M02 + mat.M11* in.M12 + mat.M12* in.M22 + mat.M13* in.M32
	m13 := mat.M10* in.M03 + mat.M11* in.M13 + mat.M12* in.M23 + mat.M13* in.M33
	m20 := mat.M20* in.M00 + mat.M21* in.M10 + mat.M22* in.M20 + mat.M23* in.M30
	m21 := mat.M20* in.M01 + mat.M21* in.M11 + mat.M22* in.M21 + mat.M23* in.M31
	m22 := mat.M20* in.M02 + mat.M21* in.M12 + mat.M22* in.M22 + mat.M23* in.M32
	m23 := mat.M20* in.M03 + mat.M21* in.M13 + mat.M22* in.M23 + mat.M23* in.M33
	m30 := mat.M30* in.M00 + mat.M31* in.M10 + mat.M32* in.M20 + mat.M33* in.M30
	m31 := mat.M30* in.M01 + mat.M31* in.M11 + mat.M32* in.M21 + mat.M33* in.M31
	m32 := mat.M30* in.M02 + mat.M31* in.M12 + mat.M32* in.M22 + mat.M33* in.M32
	m33 := mat.M30* in.M03 + mat.M31* in.M13 + mat.M32* in.M23 + mat.M33* in.M33

	return &Matrix4f{
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	}
}

func (mat *Matrix4f) MulV(in *Vector4f) *Vector4f {

	x := mat.M00 * in.X + mat.M01 * in.Y + mat.M02 * in.Z + mat.M03 * in.W
	y := mat.M10 * in.X + mat.M11 * in.Y + mat.M12 * in.Z + mat.M13 * in.W
	z := mat.M20 * in.X + mat.M21 * in.Y + mat.M22 * in.Z + mat.M23 * in.W
	w := mat.M30 * in.X + mat.M31 * in.Y + mat.M32 * in.Z + mat.M33 * in.W

	return &Vector4f{x, y, z, w}
}


