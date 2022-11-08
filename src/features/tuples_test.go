package features

import (
	"math"
	"math/rand"
	"testing"

	"github.com/jasonhongxyz/rtc/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsPoint(t *testing.T) {
	t.Parallel()

	mockPoint := Point(4.3, -4.2, 3.1)
	assert.True(t, mockPoint.IsPoint(), "mockPoint is a point")
	assert.False(t, mockPoint.IsVector(), "mockPoint is not a vector")
}

func TestIsVector(t *testing.T) {
	t.Parallel()

	mockVector := Vector(4.3, -4.2, 3.1)
	assert.True(t, mockVector.IsVector(), "mockVector is a vector")
	assert.False(t, mockVector.IsPoint(), "mockVector is not a point")
}

func TestFloatEqual(t *testing.T) {
	t.Parallel()

	x := rand.Float64()
	y := x

	assert.True(t, utils.Equal(x, y))
}

func TestAddVectorToPoint(t *testing.T) {
	t.Parallel()

	t1 := Point(3, -2, 5)
	t2 := Vector(-2, 3, 1)
	ans := Add(t1, t2)
	exp := Point(1, 1, 6)

	assert.Equal(t, exp, ans)
}

func TestAddVectorToVector(t *testing.T) {
	t.Parallel()

	t1 := Vector(3, -2, 5)
	t2 := Vector(-2, 3, 1)
	ans := Add(t1, t2)
	exp := Vector(1, 1, 6)

	assert.Equal(t, exp, ans)
}

func TestAddPointToPoint(t *testing.T) {
	t.Parallel()

	t1 := Point(3, -2, 5)
	t2 := Point(-2, 3, 1)
	ans := Add(t1, t2)
	exp := Tuple{1, 1, 6, 2}

	assert.Equal(t, exp, ans)
}

func TestSubtractPointFromPoint(t *testing.T) {
	t.Parallel()

	t1 := Point(3, 2, 1)
	t2 := Point(5, 6, 7)
	ans := Subtract(t1, t2)
	exp := Vector(-2, -4, -6)

	assert.Equal(t, exp, ans)
}

func TestSubtractVectorFromPoint(t *testing.T) {
	t.Parallel()

	t1 := Point(3, 2, 1)
	t2 := Vector(5, 6, 7)
	ans := Subtract(t1, t2)
	exp := Point(-2, -4, -6)

	assert.Equal(t, exp, ans)
}

func TestSubtractVectorFromVector(t *testing.T) {
	t.Parallel()

	t1 := Vector(3, 2, 1)
	t2 := Vector(5, 6, 7)
	ans := Subtract(t1, t2)
	exp := Vector(-2, -4, -6)

	assert.Equal(t, exp, ans)
}

func TestSubtractPointFromVector(t *testing.T) {
	t.Parallel()

	t1 := Vector(3, 2, 1)
	t2 := Point(5, 6, 7)
	ans := Subtract(t1, t2)
	exp := Tuple{-2, -4, -6, -1}

	assert.Equal(t, exp, ans)
}

func TestNegate(t *testing.T) {
	t.Parallel()

	t1 := Tuple{1, -2, 3, -4}
	t2 := Vector(3, 2, 1)
	t3 := Point(1, 2, 3)

	expA1 := Tuple{-1, 2, -3, 4}
	expA2 := Tuple{-3, -2, -1, 0}
	expA3 := Tuple{-1, -2, -3, -1}

	assert.Equal(t, expA1, Negate(t1))
	assert.Equal(t, expA2, Negate(t2))
	assert.Equal(t, expA3, Negate(t3))
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	t1 := Tuple{1, -2, 3, -4}
	scalar := 3.5
	fraction := 0.5
	expA1 := Tuple{3.5, -7, 10.5, -14}
	expA2 := Tuple{0.5, -1, 1.5, -2}

	assert.Equal(t, expA1, Multiply(t1, scalar))
	assert.Equal(t, expA2, Multiply(t1, fraction))
}

func TestDivide(t *testing.T) {
	t.Parallel()

	mockTupleA := Tuple{1, -2, 3, -4}
	div := 2.0
	expA := Tuple{0.5, -1, 1.5, -2}

	assert.Equal(t, expA, Divide(mockTupleA, div))
}

func TestMagnitude(t *testing.T) {
	t.Parallel()

	t1 := Vector(0, 1, 0)
	t2 := Vector(0, 0, 1)
	t3 := Vector(1, 2, 3)
	t4 := Vector(-1, -2, -3)

	expA1 := 1.0
	expA2 := 1.0
	expA3 := math.Sqrt(14)
	expA4 := math.Sqrt(14)

	assert.Equal(t, expA1, Magnitude(t1))
	assert.Equal(t, expA2, Magnitude(t2))
	assert.Equal(t, expA3, Magnitude(t3))
	assert.Equal(t, expA4, Magnitude(t4))
}

func TestNormalize(t *testing.T) {
	t.Parallel()

	t1 := Vector(4, 0, 0)
	t2 := Vector(1, 2, 3)

	expA1 := Vector(1, 0, 0)
	expA2 := Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))

	assert.Equal(t, expA1, Normalize(t1))
	assert.Equal(t, expA2, Normalize(t2))
}

func TestDot(t *testing.T) {
	t.Parallel()

	t1 := Vector(1, 2, 3)
	t2 := Vector(2, 3, 4)
	exp := 20.0

	assert.Equal(t, exp, DotProduct(t1, t2))
}

func TestCross(t *testing.T) {
	t.Parallel()

	t1 := Vector(1, 2, 3)
	t2 := Vector(2, 3, 4)
	exp1 := Vector(-1, 2, -1)
	exp2 := Vector(1, -2, 1)

	assert.Equal(t, exp1, CrossProduct(t1, t2))
	assert.Equal(t, exp2, CrossProduct(t2, t1))
}

func TestColor(t *testing.T) {
	t.Parallel()

	c1 := Color(-0.5, 0.4, 1.7)
	exp1 := -0.5
	exp2 := 0.4
	exp3 := 1.7

	assert.Equal(t, exp1, c1[0])
	assert.Equal(t, exp2, c1[1])
	assert.Equal(t, exp3, c1[2])
}

func TestAddColor(t *testing.T) {
	t.Parallel()

	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	c3 := Add(c1, c2)

	exp1 := 1.6
	exp2 := 0.7
	exp3 := 1.0

	assert.True(t, utils.Equal(exp1, c3[0]))
	assert.True(t, utils.Equal(exp2, c3[1]))
	assert.True(t, utils.Equal(exp3, c3[2]))
}

func TestSubtractColor(t *testing.T) {
	t.Parallel()

	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	c3 := Subtract(c1, c2)

	exp1 := 0.2
	exp2 := 0.5
	exp3 := 0.5

	assert.True(t, utils.Equal(exp1, c3[0]))
	assert.True(t, utils.Equal(exp2, c3[1]))
	assert.True(t, utils.Equal(exp3, c3[2]))
}

func TestMultiplyColor(t *testing.T) {
	t.Parallel()

	c1 := Color(0.9, 0.6, 0.75)
	c2 := Multiply(c1, 2.0)

	exp1 := 1.8
	exp2 := 1.2
	exp3 := 1.5

	assert.True(t, utils.Equal(exp1, c2[0]))
	assert.True(t, utils.Equal(exp2, c2[1]))
	assert.True(t, utils.Equal(exp3, c2[2]))
}

func TestHadamardProduct(t *testing.T) {
	t.Parallel()

	c1 := Color(1.0, 2.0, 3.0)
	c2 := Color(2.0, 3.0, 4.0)

	act := HadamardProduct(c1, c2)
	exp1 := 2.0
	exp2 := 6.0
	exp3 := 12.0

	assert.True(t, utils.Equal(exp1, act[0]))
	assert.True(t, utils.Equal(exp2, act[1]))
	assert.True(t, utils.Equal(exp3, act[2]))
}
