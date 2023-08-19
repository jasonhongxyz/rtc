package features

import (
	"fmt"
	"math"

	"github.com/jasonhongxyz/rtc/src/utils"
)

type Matrix interface {
	Multiply([][]float64, [][]float64) [][]float64
	MultiplyTuple([][]float64, Tuple) [][]float64
	Determinant([][]float64) float64
}

type Matrix4 struct{}
type Matrix3 struct{}
type Matrix2 struct{}

func NewMatrix(arr []float64) [][]float64 {
	dimension := int(math.Sqrt(float64(len(arr))))
	matrix := make([][]float64, dimension)

	i := 0
	for x := 0; x < dimension; x++ {
		matrix[x] = make([]float64, dimension)
		matrix[x] = arr[i : i+dimension]
		i += dimension
	}

	return matrix
}

func PrettyPrint(arr [][]float64) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func Equal(m1, m2 [][]float64) bool {
	for x := range m1 {
		for y := range m1[0] {
			if !utils.Equal(m1[x][y], m2[x][y]) {
				return false
			}
		}
	}
	return true
}

func (Matrix4) Multiply(m1, m2 [][]float64) [][]float64 {
	ret := make([][]float64, 4)

	for x := range m1 {
		for y := range m1[0] {
			ret[x] = append(ret[x], m1[x][0]*m2[0][y]+
				m1[x][1]*m2[1][y]+
				m1[x][2]*m2[2][y]+
				m1[x][3]*m2[3][y])
		}
	}

	return ret
}

func (Matrix4) MultiplyTuple(m [][]float64, t Tuple) Tuple {
	var ret Tuple

	z := 0
	for x := range m {
		ret[z] = m[x][0]*t[0] +
			m[x][1]*t[1] +
			m[x][2]*t[2] +
			m[x][3]*t[3]
		z++
	}

	return ret
}

func Transpose(m [][]float64) [][]float64 {
	ret := make([][]float64, len(m))
	for i := 0; i < len(m); i++ {
		ret[i] = make([]float64, len(m))
	}

	for x := range m {
		for y := range m[0] {
			ret[y][x] = m[x][y]
		}
	}

	return ret
}

func (Matrix2) Determinant(m [][]float64) float64 {
	determinant := (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	return determinant
}

func (Matrix3) Determinant(m [][]float64) float64 {
	det := 0.0

	for i := 0; i < 3; i++ {
		det = det + m[0][i]*Cofactor(m, 0, i)
	}

	return det
}

func Submatrix(m [][]float64, r, c int) [][]float64 {
	var vals []float64
	for row := range m {
		for col := range m {
			if row == r || col == c {
				continue
			} else {
				vals = append(vals, m[row][col])
			}
		}
	}

	ret := NewMatrix(vals)

	return ret
}

func Minor(m [][]float64, r, c int) float64 {
	subm := Submatrix(m, r, c)

	return Determinant(subm)
}

func Cofactor(m [][]float64, r, c int) float64 {
	if (r+c)%2 == 1 {
		return -1.0 * Minor(m, r, c)
	}

	return Minor(m, r, c)
}

func Determinant(m [][]float64) float64 {
	det := 0.0
	s := len(m)

	if s == 2 {
		det = (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	} else {
		for i := 0; i < s; i++ {
			det = det + m[0][i]*Cofactor(m, 0, i)
		}
	}

	return det
}

func IsInvertible(m [][]float64) bool {
	if Determinant(m) == 0 {
		return false
	}
	return true
}

func Inverse(m [][]float64) ([][]float64, error) {
	if !IsInvertible(m) {
		return nil, fmt.Errorf("Matrix not invertible!")
	}

	s := len(m)

	ret := make([][]float64, s)
	for i := 0; i < s; i++ {
		ret[i] = make([]float64, s)
	}

	for row := 0; row < s; row++ {
		for col := 0; col < s; col++ {
			c := Cofactor(m, row, col)
			ret[col][row] = c / Determinant(m)
		}
	}

	return ret, nil
}
