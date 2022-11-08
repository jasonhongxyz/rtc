package features

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix4(t *testing.T) {
	vals := []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}
	m := New4DMatrix(vals)

	assert.Equal(t, 1.0, m[0][0])
	assert.Equal(t, 4.0, m[0][3])
	assert.Equal(t, 5.5, m[1][0])
	assert.Equal(t, 7.5, m[1][2])
	assert.Equal(t, 11.0, m[2][2])
	assert.Equal(t, 13.5, m[3][0])
	assert.Equal(t, 15.5, m[3][2])

}

func TestNewMatrix3(t *testing.T) {
	vals := []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}
	m := New3DMatrix(vals)

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, -2.0, m[1][1])
	assert.Equal(t, 1.0, m[2][2])
}

func TestNewMatrix2(t *testing.T) {
	vals := []float64{-3, 5, 1, -2}
	m := New2DMatrix(vals)

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, 5.0, m[0][1])
	assert.Equal(t, 1.0, m[1][0])
	assert.Equal(t, -2.0, m[1][1])
}

func TestEqualMatrix4(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}

	mA := New4DMatrix(valsA)
	mB := New4DMatrix(valsB)

	act := Equal4DMatrix(mA, mB)

	assert.True(t, act)
}

func TestNotEqualMatrix4(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	mA := New4DMatrix(valsA)
	mB := New4DMatrix(valsB)

	act := Equal4DMatrix(mA, mB)

	assert.False(t, act)
}

func TestMultiply4DMatrix(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	valsB := []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}

	mA := New4DMatrix(valsA)
	mB := New4DMatrix(valsB)

	act := Multiply4DMatrix(mA, mB)

	expVals := []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}
	exp := New4DMatrix(expVals)

	eq := Equal4DMatrix(act, exp)

	assert.True(t, eq)
}

func TestMultiply4DMatrixTuple(t *testing.T) {
	valsA := []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}
	mA := New4DMatrix(valsA)

	tuple := Tuple{1, 2, 3, 1}

	act := Multiply4DMatrixTuple(mA, tuple)
	exp := Tuple{18, 24, 33, 1}

	assert.Equal(t, exp, act)
}

func TestMultiplyIdentityMatrix(t *testing.T) {
	valsA := []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32}
	valsB := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	mA := New4DMatrix(valsA)
	mB := New4DMatrix(valsB)

	act := Multiply4DMatrix(mA, mB)

	eq := Equal4DMatrix(act, mA)

	assert.True(t, eq)
}

func TestTranspose4DMatrix(t *testing.T) {
	valsA := []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}
	transposeA := []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}

	mA := New4DMatrix(valsA)
	exp := New4DMatrix(transposeA)

	act := Transpose4DMatrix(mA)

	eq := Equal4DMatrix(exp, act)

	assert.True(t, eq)
}

func TestTransposeIdentity(t *testing.T) {
	valsA := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	identity_matrix := New4DMatrix(valsA)

	act := Transpose4DMatrix(identity_matrix)

	eq := Equal4DMatrix(act, identity_matrix)

	assert.True(t, eq)
}
