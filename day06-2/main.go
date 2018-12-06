package main

import (
	c "adventofcode2018/common"
	"math"
	"strconv"
	"strings"
)

type P struct {
	X   int
	Y   int
	Inf bool
	ID  int
	Sum int
}

func main() {

	points := map[int]*P{}

	//var err error
	maxX, maxY := 0, 0
	id := 0
	for _, v := range c.ReadInputFile() {
		a := P{ID: id}
		id++
		xy := strings.Split(v, ", ")
		a.X, _ = strconv.Atoi(xy[0])
		a.Y, _ = strconv.Atoi(xy[1])
		if a.X > maxX {
			maxX = a.X
		}
		if a.Y > maxY {
			maxY = a.Y
		}
		points[a.ID] = &a
	}
	sum := 0
	for x := -50; x <= maxX+50; x++ {
		for y := -50; y <= maxY+50; y++ {
			curSum := 0
			ok := true
			for i := range points {
				d := int(math.Abs(float64(y-points[i].Y)) + math.Abs(float64(x-points[i].X)))
				curSum += d
				if curSum >= 10000 {
					ok = false
					break
				}
			}
			if ok {
				sum++
			}
		}
	}
	c.Print(sum)

}
