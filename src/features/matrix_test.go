package features

import (
	"log"
	"testing"

	"github.com/jasonhongxyz/rtc/src/utils"
	"github.com/stretchr/testify/assert"
)

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

	matrix3 := Matrix3{}

	act := matrix3.Minor(A, 1, 0)

	assert.Equal(t, act, exp)
}

func TestNewMatrix(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := NewMatrix(valsA)

	valsB := []float64{1}
	B := NewMatrix(valsB)

	valsC := []float64{1, 2, 3, 4}
	C := NewMatrix(valsC)

	valsD := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	D := NewMatrix(valsD)

	log.Println(A, B, C, D)
}
