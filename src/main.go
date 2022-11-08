package main

import (
	"github.com/jasonhongxyz/rtc/src/features"
	"github.com/jasonhongxyz/rtc/src/utils"
)

func main() {
	p := features.Proj{Position: features.Point(0, 1, 0), Velocity: features.Normalize(features.Vector(1, 1, 0))}
	e := features.Env{Gravity: features.Vector(0, -0.1, 0), Wind: features.Vector(-0.01, 0, 0)}

	for x := 0; x < 500 && p.Position[1] >= 0; x++ {
		p = features.Tick(e, p)
		// fmt.Printf("Projectile: %v\n", p.Position)
		// fmt.Printf("Bounce: %v\n", x)
	}

	start := features.Point(0, 1, 0)
	velocity := features.Multiply(features.Normalize(features.Vector(1, 1.8, 0)), 11.25)

	p = features.Proj{Position: start, Velocity: velocity}

	gravity := features.Vector(0, -0.1, 0)
	wind := features.Vector(-0.01, 0, 0)

	e = features.Env{Gravity: gravity, Wind: wind}

	cv := features.NewCanvas(900, 550)

	for x := 0; x < 500 && p.Position[1] >= 0; x++ {
		c := features.Color(1, 0, 0)
		p = features.Tick(e, p)

		if p.Position[0] > 0 && p.Position[1] > 0 {
			cv.WritePixel(int(p.Position[0]), 550-int(p.Position[1]), c)
		}
	}
	cv.Out(utils.GetFilenameDateTime())
}
