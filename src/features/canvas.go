package features

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Pixels [][]Tuple
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewCanvas(w, h int) Canvas {
	var canvas Canvas
	canvas.Width = w
	canvas.Height = h

	for x := 0; x < h; x++ {
		row := []Tuple{}
		for y := 0; y < w; y++ {
			row = append(row, Color(0.0, 0.0, 0.0))
		}
		canvas.Pixels = append(canvas.Pixels, row)
	}

	return canvas
}

func (c *Canvas) WritePixel(x, y int, color Tuple) {
	c.Pixels[y][x] = color
}

func (c *Canvas) ToPPM() string {
	ret := ""
	maxColor := 255

	// Line 1: Magic number
	ret += "P3\n"

	// Line 2: Width & Height
	w := len(c.Pixels[0])
	h := len(c.Pixels)
	ret += strconv.Itoa(w) + " " + strconv.Itoa(h) + "\n"

	// Line 3: Max color val
	ret += (strconv.Itoa(maxColor) + "\n")

	for row := range c.Pixels {
		ret += clampRowAndStringify(c.Pixels[row], maxColor)
	}

	return ret
}

func clampRowAndStringify(color []Tuple, maxColor int) string {
	str := ""

	row_arr := []string{}
	line_len := 0

	for x := range color {
		for y := 0; y < 3; y++ {
			c := color[x][y] * float64(maxColor)
			cr := math.Round(c)

			if cr < 0 {
				cr = 0
			} else if cr > 255 {
				cr = 255
			}
			row_arr = append(row_arr, strconv.Itoa(int(cr)))
		}
	}

	str = strings.Join(row_arr, " ")

	// Fold lines with newline > 70
	for x := range str {
		line_len++
		if line_len+4 >= 70 && string(str[x]) == " " {
			rune_str := []rune(str)
			rune_str[x] = rune('\n')
			str = string(rune_str)
			line_len = 0
		}
	}
	str += "\n"

	return str
}

func (c *Canvas) Out(file string) {
	f, err := os.Create(file)
	check(err)
	defer f.Close()

	n, err := f.WriteString(c.ToPPM())
	check(err)
	fmt.Printf("Created out.ppm file: %d\n", n)
	f.Sync()
}
