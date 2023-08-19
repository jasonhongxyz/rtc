package features

import (
	"testing"

	"github.com/jasonhongxyz/rtc/src/utils"
	"github.com/stretchr/testify/assert"
)

// func TestNewMatrix(t *testing.T) {
// 	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	A := NewMatrix(valsA)
//
// 	valsB := []float64{1}
// 	B := NewMatrix(valsB)
//
// 	valsC := []float64{1, 2, 3, 4}
// 	C := NewMatrix(valsC)
//
// 	valsD := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
// 	D := NewMatrix(valsD)
//
// 	log.Println(A, B, C, D)
// }

func TestNewMatrix4D(t *testing.T) {
	vals := []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}
	m := NewMatrix(vals)

	assert.Equal(t, 1.0, m[0][0])
	assert.Equal(t, 4.0, m[0][3])
	assert.Equal(t, 5.5, m[1][0])
	assert.Equal(t, 7.5, m[1][2])
	assert.Equal(t, 11.0, m[2][2])
	assert.Equal(t, 13.5, m[3][0])
	assert.Equal(t, 15.5, m[3][2])
}

func TestNewMatrix3D(t *testing.T) {
	vals := []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}
	m := NewMatrix(vals)

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, -2.0, m[1][1])
	assert.Equal(t, 1.0, m[2][2])
}

func TestNewMatrix2D(t *testing.T) {
	vals := []float64{-3, 5, 1, -2}
	m := NewMatrix(vals)

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, 5.0, m[0][1])
	assert.Equal(t, 1.0, m[1][0])
	assert.Equal(t, -2.0, m[1][1])
}

func TestEqual4D(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}

	mA := NewMatrix(valsA)
	mB := NewMatrix(valsB)

	act := Equal(mA, mB)

	assert.True(t, act)
}

func TestNotEqualMatrix4D(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	mA := NewMatrix(valsA)
	mB := NewMatrix(valsB)

	act := Equal(mA, mB)

	assert.False(t, act)
}

func TestMultiply4D(t *testing.T) {
	matrix4 := Matrix4{}
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}

	mA := NewMatrix(valsA)
	mB := NewMatrix(valsB)

	act := matrix4.Multiply(mA, mB)

	expVals := []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}
	exp := NewMatrix(expVals)

	eq := Equal(act, exp)

	assert.True(t, eq)
}

func TestMultiplyTuple(t *testing.T) {
	matrix4 := Matrix4{}
	valsA := []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}
	mA := NewMatrix(valsA)

	tuple := Tuple{1, 2, 3, 1}

	act := matrix4.MultiplyTuple(mA, tuple)
	exp := Tuple{18, 24, 33, 1}

	assert.Equal(t, exp, act)
}

func TestMultiplyIdentityMatrix(t *testing.T) {
	matrix4 := Matrix4{}
	valsA := []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32}
	valsB := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	mA := NewMatrix(valsA)
	mB := NewMatrix(valsB)

	act := matrix4.Multiply(mA, mB)

	eq := Equal(act, mA)

	assert.True(t, eq)
}

func TestTranspose4D(t *testing.T) {
	valsA := []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}
	transposeA := []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}

	mA := NewMatrix(valsA)
	exp := NewMatrix(transposeA)

	act := Transpose(mA)

	eq := Equal(exp, act)

	assert.True(t, eq)
}

func TestTransposeIdentity(t *testing.T) {
	valsA := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	identity_matrix := NewMatrix(valsA)

	act := Transpose(identity_matrix)

	eq := Equal(act, identity_matrix)

	assert.True(t, eq)
}

func TestDeterminant2D(t *testing.T) {
	matrix2 := Matrix2{}
	valsA := []float64{1, 5, -3, 2}
	A := NewMatrix(valsA)

	exp := 17.0
	act := matrix2.Determinant(A)

	eq := utils.Equal(exp, act)

	assert.True(t, eq)
}

func TestSubmatrix3D(t *testing.T) {
	valsA := []float64{1, 5, 0, -3, 2, 7, 0, 6, -3}
	A := NewMatrix(valsA)

	valsExp := []float64{-3, 2, 0, 6}
	exp := NewMatrix(valsExp)

	act := Submatrix(A, 0, 2)

	assert.Equal(t, act, exp)
}

func TestSubmatrix4D(t *testing.T) {
	valsA := []float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1}
	A := NewMatrix(valsA)

	valsExp := []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1}
	exp := NewMatrix(valsExp)

	act := Submatrix(A, 2, 1)

	assert.Equal(t, act, exp)
}

func TestMinor3D(t *testing.T) {
	valsA := []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}
	A := NewMatrix(valsA)

	exp := 25.0

	act := Minor(A, 1, 0)

	assert.Equal(t, act, exp)
}

func TestCofactor3D(t *testing.T) {
	valsA := []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}
	A := NewMatrix(valsA)

	exp1 := -12.0
	exp2 := -25.0

	act1 := Cofactor(A, 0, 0)
	act2 := Cofactor(A, 1, 0)

	assert.Equal(t, act1, exp1)
	assert.Equal(t, act2, exp2)

}

func TestDeterminant3D(t *testing.T) {
	valsA := []float64{1, 2, 6, -5, 8, -4, 2, 6, 4}
	A := NewMatrix(valsA)

	exp := -196.0

	act := Determinant(A)

	assert.Equal(t, exp, act)
}

func TestDeterminant4D(t *testing.T) {
	valsA := []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9}
	A := NewMatrix(valsA)

	exp := -4071.0

	act := Determinant(A)

	assert.Equal(t, exp, act)
}

func TestInvertible(t *testing.T) {
	valsA := []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6}
	valsB := []float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0}
	A := NewMatrix(valsA)
	B := NewMatrix(valsB)

	assert.True(t, IsInvertible(A))
	assert.False(t, IsInvertible(B))
}

func TestInverse(t *testing.T) {
	valsA := []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4}
	A := NewMatrix(valsA)

	valsB := []float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4}
	B := NewMatrix(valsB)

	valsC := []float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2}
	C := NewMatrix(valsC)

	valsExpA := []float64{0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639}
	expA := NewMatrix(valsExpA)

	valsExpB := []float64{-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308}
	expB := NewMatrix(valsExpB)

	valsExpC := []float64{-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333}
	expC := NewMatrix(valsExpC)

	actA, _ := Inverse(A)
	actB, _ := Inverse(B)
	actC, _ := Inverse(C)

	assert.True(t, Equal(actA, expA))
	assert.True(t, Equal(actB, expB))
	assert.True(t, Equal(actC, expC))
}

func TestProductByInverse(t *testing.T) {
	valsA := []float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1}
	valsB := []float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5}
	A := NewMatrix(valsA)
	B := NewMatrix(valsB)

	matrix4 := Matrix4{}
	C := matrix4.Multiply(A, B)

	invB, _ := Inverse(B)

	assert.True(t, Equal(A, matrix4.Multiply(C, invB)))
}
