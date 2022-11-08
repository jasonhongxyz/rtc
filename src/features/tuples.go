package features

import "math"

type Tuple [4]float64

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Color(r, g, b float64) Tuple {
	return Tuple{r, g, b, -1.0}
}

func (i Tuple) IsVector() bool {
	return i[3] == 0.0
}

func (i Tuple) IsPoint() bool {
	return i[3] == 1.0
}

func Add(a1, a2 Tuple) Tuple {
	var ret [4]float64
	for i := range a1 {
		ret[i] = a1[i] + a2[i]
	}
	return ret
}

func Subtract(a1, a2 Tuple) Tuple {
	var ret [4]float64
	for i := range a1 {
		ret[i] = a1[i] - a2[i]
	}
	return ret
}

func Negate(t Tuple) Tuple {
	var ret [4]float64
	for i := range t {
		ret[i] = -1 * t[i]
	}
	return ret
}

func Multiply(t Tuple, m float64) Tuple {
	var ret [4]float64
	for i := range t {
		ret[i] = m * t[i]
	}
	return ret
}

func Divide(t Tuple, m float64) Tuple {
	var ret [4]float64
	for i := range t {
		ret[i] = t[i] / m
	}
	return ret
}

func Magnitude(t Tuple) float64 {
	var ret float64

	u := 0.0
	for _, x := range t {
		u = u + math.Pow(x, 2.0)
	}

	ret = math.Sqrt(u)

	return ret
}

func Normalize(t Tuple) Tuple {
	var ret [4]float64

	mag := Magnitude(t)
	ret = Divide(t, mag)

	return ret
}

func DotProduct(a1, a2 Tuple) float64 {
	ret := 0.0

	for i := range a1 {
		ret = ret + (a1[i] * a2[i])
	}

	return ret
}

func CrossProduct(a, b Tuple) Tuple {
	x := (a[1] * b[2]) - (a[2] * b[1])
	y := (a[2] * b[0]) - (a[0] * b[2])
	z := (a[0] * b[1]) - (a[1] * b[0])

	return Vector(x, y, z)
}

func HadamardProduct(x, y Tuple) Tuple {
	r := x[0] * y[0]
	g := x[1] * y[1]
	b := x[2] * y[2]

	return Color(r, g, b)
}
