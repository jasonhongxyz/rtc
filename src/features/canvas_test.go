package features

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	cv := NewCanvas(10, 20)

	assert.Equal(t, 10, cv.Width)
	assert.Equal(t, 20, cv.Height)

	for x := range cv.Pixels {
		for y := range cv.Pixels[0] {
			assert.Equal(t, Tuple{0.0, 0.0, 0.0, -1.0}, cv.Pixels[x][y])
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Color(1, 0, 0)

	c.WritePixel(2, 3, red)

	assert.Equal(t, red, c.Pixels[3][2])
}

func TestCanvasToPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)

	strc := c.ToPPM()
	exp := "P3\n5 3\n255\n"

	assert.True(t, strings.HasPrefix(strc, exp))

}

func TestCanvasToPPM(t *testing.T) {
	c := NewCanvas(5, 3)

	c1 := Color(1.5, 0, 0)
	c2 := Color(0, 0.5, 0)
	c3 := Color(-0.5, 0, 1)

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	strc := c.ToPPM()

	exp := "P3\n5 3\n255\n255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"

	assert.Equal(t, exp, strc)

}

func TestCanvasToPPMRowLengthLimit(t *testing.T) {
	c := NewCanvas(10, 2)
	c1 := Color(1, 0.8, 0.6)

	for x := 0; x < 10; x++ {
		for y := 0; y < 2; y++ {
			c.WritePixel(x, y, c1)
		}
	}

	strc := c.ToPPM()

	exp := "P3\n10 2\n255\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n"

	assert.Equal(t, exp, strc)
}

func TestCanvasToPPMEndNewline(t *testing.T) {
	c := NewCanvas(5, 3)
	strc := c.ToPPM()
	exp := "\n"

	assert.True(t, strings.HasSuffix(strc, exp))
}
