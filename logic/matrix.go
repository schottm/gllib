package logic

//Matrix2f//

type Matrix2f struct {
	M00, M01 float32
	M10, M11 float32
}

func NewScaleMatrix2f(scaleX, scaleY float32) *Matrix2f {

	return &Matrix2f{scaleX, 0.0, 0.0, scaleY}
}

func NewIdentityMatrix2f() *Matrix2f {

	return &Matrix2f{1.0, 0.0, 0.0, 1.0}
}

func (mat *Matrix2f) Identity() {

	mat.M00 = 1.0; mat.M01 = 0.0
	mat.M10 = 0.0; mat.M11 = 1.0
}

func (mat *Matrix2f) Invert() *Matrix2f {

	det := mat.M00 * mat.M11 - mat.M10 * mat.M01
	if det == 0 {
		return nil
	}
	det = 1.0 / det

	return &Matrix2f{det * mat.M11, det * -mat.M01, det * -mat.M10, det * mat.M00}
}

func (mat *Matrix2f) Mul(in *Matrix2f) *Matrix2f {

	var res = Matrix2f{}

	res.M00 = mat.M00* in.M00 + mat.M01* in.M10
	res.M01 = mat.M00* in.M01 + mat.M01* in.M11
	res.M10 = mat.M10* in.M00 + mat.M11* in.M10
	res.M11 = mat.M10* in.M01 + mat.M11* in.M11

	return &res
}

func (mat *Matrix2f) MulV(in *Vector2f) *Vector2f {

	var res = Vector2f{}

	res.X = mat.M00 * in.X + mat.M01 * in.Y
	res.Y = mat.M10 * in.X + mat.M11 * in.Y

	return &res
}

//Matrix3f//

type Matrix3f struct {
	M00, M01, M02 float32
	M10, M11, M12 float32
	M20, M21, M22 float32
}

func NewScaleMatrix3f(scaleX, scaleY, scaleZ float32) *Matrix3f {

	return &Matrix3f{
		scaleX, 0, 0,
		0, scaleY, 0,
		0, 0, scaleZ}
}

func NewIdentityMatrix3f() *Matrix3f {

	return &Matrix3f{1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0}
}

func (mat *Matrix3f) Identity() {

	mat.M00 = 1.0; mat.M01 = 0.0; mat.M02 = 0.0
	mat.M10 = 0.0; mat.M11 = 1.0; mat.M12 = 0.0
	mat.M20 = 0.0; mat.M21 = 0.0; mat.M22 = 1.0
}

func (mat *Matrix3f) Invert() *Matrix3f {

	var res = Matrix3f{}

	det := mat.M00 * (mat.M11 * mat.M22 - mat.M21 * mat.M12) -
		mat.M01 * (mat.M10 * mat.M22 - mat.M12 * mat.M20) +
		mat.M02 * (mat.M10 * mat.M21 - mat.M11 * mat.M20)
	if det == 0 {
		return nil
	}
	det = 1.0 / det

	res.M00 = (mat.M11 * mat.M22 - mat.M21 * mat.M12) * det
	res.M01 = (mat.M02 * mat.M21 - mat.M01 * mat.M22) * det
	res.M02 = (mat.M01 * mat.M12 - mat.M02 * mat.M11) * det
	res.M10 = (mat.M12 * mat.M20 - mat.M10 * mat.M22) * det
	res.M11 = (mat.M00 * mat.M22 - mat.M02 * mat.M20) * det
	res.M12 = (mat.M10 * mat.M02 - mat.M00 * mat.M12) * det
	res.M20 = (mat.M10 * mat.M21 - mat.M20 * mat.M11) * det
	res.M21 = (mat.M20 * mat.M01 - mat.M00 * mat.M21) * det
	res.M22 = (mat.M00 * mat.M11 - mat.M10 * mat.M01) * det

	return &res
}

func (mat *Matrix3f) Mul(in *Matrix3f) *Matrix3f {

	var res = Matrix3f{}

	res.M00 = mat.M00 * in.M00 + mat.M10 * in.M01 + mat.M20 * in.M02
	res.M01 = mat.M01 * in.M00 + mat.M11 * in.M01 + mat.M21 * in.M02
	res.M02 = mat.M02 * in.M00 + mat.M12 * in.M01 + mat.M22 * in.M02
	res.M10 = mat.M00 * in.M10 + mat.M10 * in.M11 + mat.M20 * in.M12
	res.M11 = mat.M01 * in.M10 + mat.M11 * in.M11 + mat.M21 * in.M12
	res.M12 = mat.M02 * in.M10 + mat.M12 * in.M11 + mat.M22 * in.M12
	res.M20 = mat.M00 * in.M20 + mat.M10 * in.M21 + mat.M20 * in.M22
	res.M21 = mat.M01 * in.M20 + mat.M11 * in.M21 + mat.M21 * in.M22
	res.M22 = mat.M02 * in.M20 + mat.M12 * in.M21 + mat.M22 * in.M22

	return &res
}

func (mat *Matrix3f) MulV(in *Vector3f) {

	var res = Vector3f{}

	res.X = mat.M00 * in.X + mat.M10 * in.Y + mat.M20 * in.Z
	res.Y = mat.M01 * in.X + mat.M11 * in.Y + mat.M21 * in.Z
	res.Z = mat.M02 * in.X + mat.M12 * in.Y + mat.M22 * in.Z
}

func (mat *Matrix3f) Translate(in *Vector2f) {

	mat.M20 += mat.M00 * in.X + mat.M10 * in.Y
	mat.M21 += mat.M01 * in.X + mat.M11 * in.Y
}

//Matrix4f//

type Matrix4f struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

func NewScaleMatrix4f(scaleX, scaleY, scaleZ, scaleW float32) *Matrix4f {

	return &Matrix4f{
		scaleX, 0, 0, 0,
		0, scaleY, 0, 0,
		0, 0, scaleZ, 0,
		0, 0, 0, scaleW}
}

func NewIdentityMatrix4f() *Matrix4f {

	return &Matrix4f{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0}
}

func (mat *Matrix4f) Identity() {

	mat.M00 = 1.0; mat.M01 = 0.0; mat.M02 = 0.0; mat.M03 = 0.0
	mat.M10 = 0.0; mat.M11 = 1.0; mat.M12 = 0.0; mat.M13 = 0.0
	mat.M20 = 0.0; mat.M21 = 0.0; mat.M22 = 1.0; mat.M23 = 0.0
	mat.M30 = 0.0; mat.M31 = 0.0; mat.M32 = 0.0; mat.M33 = 1.0
}

func (mat *Matrix4f) Invert() *Matrix4f {

	var res = Matrix4f{}

	s0 := mat.M00 * mat.M11 - mat.M10 * mat.M01
	s1 := mat.M00 * mat.M12 - mat.M10 * mat.M02
	s2 := mat.M00 * mat.M13 - mat.M10 * mat.M03
	s3 := mat.M01 * mat.M12 - mat.M11 * mat.M02
	s4 := mat.M01 * mat.M13 - mat.M11 * mat.M03
	s5 := mat.M02 * mat.M13 - mat.M12 * mat.M03
	c5 := mat.M22 * mat.M33 - mat.M32 * mat.M23
	c4 := mat.M21 * mat.M33 - mat.M31 * mat.M23
	c3 := mat.M21 * mat.M32 - mat.M31 * mat.M22
	c2 := mat.M20 * mat.M33 - mat.M30 * mat.M23
	c1 := mat.M20 * mat.M32 - mat.M30 * mat.M22
	c0 := mat.M20 * mat.M31 - mat.M30 * mat.M21

	det := s0 * c5 - s1 * c4 + s2 * c3 + s3 * c2 - s4 * c1 + s5 * c0
	if det == 0 {
		return nil
	}
	det = 1.0 / det

	res.M00 = (mat.M11 * c5 - mat.M12 * c4 + mat.M13 * c3) * det
	res.M01 = (-mat.M01 * c5 + mat.M02 * c4 - mat.M03 * c3) * det
	res.M02 = (mat.M31 * s5 - mat.M32 * s4 + mat.M33 * s3) * det
	res.M03 = (-mat.M21 * s5 + mat.M22 * s4 - mat.M23 * s3) * det
	res.M10 = (-mat.M10 * c5 + mat.M12 * c2 - mat.M13 * c1) * det
	res.M11 = (mat.M00 * c5 - mat.M02 * c2 + mat.M03 * c1) * det
	res.M12 = (-mat.M30 * s5 + mat.M32 * s2 - mat.M33 * s1) * det
	res.M13 = (mat.M20 * s5 - mat.M22 * s2 + mat.M23 * s1) * det
	res.M20 = (mat.M10 * c4 - mat.M11 * c2 + mat.M13 * c0) * det
	res.M21 = (-mat.M00 * c4 + mat.M01 * c2 - mat.M03 * c0) * det
	res.M22 = (mat.M30 * s4 - mat.M31 * s2 + mat.M33 * s0) * det
	res.M23 = (-mat.M20 * s4 + mat.M21 * s2 - mat.M23 * s0) * det
	res.M30 = (-mat.M10 * c3 + mat.M11 * c1 - mat.M12 * c0) * det
	res.M31 = (mat.M00 * c3 - mat.M01 * c1 + mat.M02 * c0) * det
	res.M32 = (-mat.M30 * s3 + mat.M31 * s1 - mat.M32 * s0) * det
	res.M33 = (mat.M20 * s3 - mat.M21 * s1 + mat.M22 * s0) * det

	return &res
}

func (mat *Matrix4f) Mul(in *Matrix4f) *Matrix4f{

	var res = Matrix4f{}

	res.M00 = mat.M00 * in.M00 + mat.M10 * in.M01 + mat.M20 * in.M02 + mat.M30 * in.M03
	res.M01 = mat.M01 * in.M00 + mat.M11 * in.M01 + mat.M21 * in.M02 + mat.M31 * in.M03
	res.M02 = mat.M02 * in.M00 + mat.M12 * in.M01 + mat.M22 * in.M02 + mat.M32 * in.M03
	res.M03 = mat.M03 * in.M00 + mat.M13 * in.M01 + mat.M23 * in.M02 + mat.M33 * in.M03
	res.M10 = mat.M00 * in.M10 + mat.M10 * in.M11 + mat.M20 * in.M12 + mat.M30 * in.M13
	res.M11 = mat.M01 * in.M10 + mat.M11 * in.M11 + mat.M21 * in.M12 + mat.M31 * in.M13
	res.M12 = mat.M02 * in.M10 + mat.M12 * in.M11 + mat.M22 * in.M12 + mat.M32 * in.M13
	res.M13 = mat.M03 * in.M10 + mat.M13 * in.M11 + mat.M23 * in.M12 + mat.M33 * in.M13
	res.M20 = mat.M00 * in.M20 + mat.M10 * in.M21 + mat.M20 * in.M22 + mat.M30 * in.M23
	res.M21 = mat.M01 * in.M20 + mat.M11 * in.M21 + mat.M21 * in.M22 + mat.M31 * in.M23
	res.M22 = mat.M02 * in.M20 + mat.M12 * in.M21 + mat.M22 * in.M22 + mat.M32 * in.M23
	res.M23 = mat.M03 * in.M20 + mat.M13 * in.M21 + mat.M23 * in.M22 + mat.M33 * in.M23
	res.M30 = mat.M00 * in.M30 + mat.M10 * in.M31 + mat.M20 * in.M32 + mat.M30 * in.M33
	res.M31 = mat.M01 * in.M30 + mat.M11 * in.M31 + mat.M21 * in.M32 + mat.M31 * in.M33
	res.M32 = mat.M02 * in.M30 + mat.M12 * in.M31 + mat.M22 * in.M32 + mat.M32 * in.M33
	res.M33 = mat.M03 * in.M30 + mat.M13 * in.M31 + mat.M23 * in.M32 + mat.M33 * in.M33

	return &res
}

func (mat *Matrix4f) MulV(in *Vector4f) *Vector4f {

	var res = Vector4f{}

	res.X = mat.M00 * in.X + mat.M01 * in.Y + mat.M02 * in.Z + mat.M03 * in.W
	res.Y = mat.M10 * in.X + mat.M11 * in.Y + mat.M12 * in.Z + mat.M13 * in.W
	res.Z = mat.M20 * in.X + mat.M21 * in.Y + mat.M22 * in.Z + mat.M23 * in.W
	res.W = mat.M30 * in.X + mat.M31 * in.Y + mat.M32 * in.Z + mat.M33 * in.W

	return &res
}

func (mat *Matrix4f) Translate(in *Vector3f) {

	mat.M30 += mat.M00 * in.X + mat.M10 * in.Y + mat.M20 * in.Z
	mat.M31 += mat.M01 * in.X + mat.M11 * in.Y + mat.M21 * in.Z
	mat.M32 += mat.M02 * in.X + mat.M12 * in.Y + mat.M22 * in.Z
}


