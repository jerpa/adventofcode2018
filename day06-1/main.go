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
	field := map[int]map[int]int{}
	for x := 0; x <= maxX; x++ {
		field[x] = map[int]int{}
		for y := 0; y <= maxY; y++ {
			curID := -1
			curDist := -1
			eq := false
			for i := range points {
				d := int(math.Abs(float64(y-points[i].Y)) + math.Abs(float64(x-points[i].X)))
				if d == curDist {
					eq = true
				} else if curDist == -1 || d < curDist {
					eq = false
					curDist = d
					curID = points[i].ID
				}
			}
			if eq {
				field[x][y] = -1
			} else {
				field[x][y] = curID
				points[curID].Sum = points[curID].Sum + 1
			}
		}
	}
	for x := 0; x <= maxX; x++ {
		if field[x][0] > -1 {
			points[field[x][0]].Inf = true
		}
		if field[x][maxY] > -1 {
			points[field[x][maxY]].Inf = true
		}
	}
	for y := 0; y <= maxY; y++ {
		if field[0][y] > -1 {
			points[field[0][y]].Inf = true
		}
		if field[maxX][y] > -1 {
			points[field[maxX][y]].Inf = true
		}
	}
	maxVal := 0
	for i := range points {
		if points[i].Inf == false && maxVal < points[i].Sum {
			maxVal = points[i].Sum
		}
	}
	c.Print(maxVal)

}
