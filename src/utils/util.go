package utils

import (
	"math"
	"time"
)

const EPSILON = 0.00001

func GetFilenameDateTime() string {
	const layout = "15_04_05.01-02-2006"

	t := time.Now()

	return "output/output-" + t.Format(layout) + ".ppm"
}

func Equal(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	}
	return false
}
