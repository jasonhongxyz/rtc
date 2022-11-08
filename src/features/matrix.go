package features

import (
	"fmt"
	"github.com/jasonhongxyz/rtc/src/utils"
	"golang.org/x/exp/constraints"
)

type Matrix4 [4][4]float64
type Matrix3 [3][3]float64
type Matrix2 [2][2]float64

type Number interface {
	constraints.Integer | constraints.Float
}

func New4DMatrix[T Number](vals []T) [4][4]T {
	var ret [4][4]T

	z := 0
	for x := range ret {
		for y := range ret[0] {
			ret[x][y] = vals[z]
			z++
		}
	}

	return ret
}

func New3DMatrix[T Number](vals []T) [3][3]T {
	var ret [3][3]T

	z := 0
	for x := range ret {
		for y := range ret[0] {
			ret[x][y] = vals[z]
			z++
		}
	}

	return ret
}

func New2DMatrix[T Number](vals []T) [2][2]T {
	var ret [2][2]T

	z := 0
	for x := range ret {
		for y := range ret[0] {
			ret[x][y] = vals[z]
			z++
		}
	}

	return ret
}

func Equal4DMatrix(m1, m2 Matrix4) bool {
	for x := range m1 {
		for y := range m1[0] {
			if !utils.Equal(m1[x][y], m2[x][y]) {
				return false
			}
		}
	}

	return true
}


func Multiply4DMatrix[T Number](m1, m2 [4][4]T) [4][4]T {
	var ret [4][4]T

	for x := range m1 {
		for y := range m1[0] {
			ret[x][y] = m1[x][0] * m2[0][y] +
									m1[x][1] * m2[1][y] +
									m1[x][2] * m2[2][y] +
									m1[x][3] * m2[3][y]
		}
	}

	return ret
}

func Multiply4DMatrixTuple(m [4][4]float64, t Tuple) Tuple {
	var ret Tuple

	z := 0
	for x := range(m) {
		ret[z] = m[x][0] * t[0] +
						m[x][1] * t[1] +
						m[x][2] * t[2] +
						m[x][3] * t[3]
		z ++
	}

	return ret
}

func Transpose4DMatrix[T Number](m [4][4]T) [4][4]T {
	var ret [4][4]T

	for x := range(m) {
		for y := range(m[0]) {
			ret[y][x] = m[x][y]
		}
	}

	fmt.Print(ret)

	return ret
}
