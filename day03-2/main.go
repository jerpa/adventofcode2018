package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang/geo/r2"
)

type d struct {
	id   string
	rect r2.Rect
}

func main() {
	data := []d{}

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		p1 := strings.Split(s[2], ",")
		p2 := strings.Split(s[3], "x")
		p1[1] = strings.TrimRight(p1[1], ":")
		px1, _ := strconv.ParseFloat(p1[0], 64)
		py1, _ := strconv.ParseFloat(p1[1], 64)
		px2, _ := strconv.ParseFloat(p2[0], 64)
		py2, _ := strconv.ParseFloat(p2[1], 64)
		data = append(data, d{rect: r2.RectFromPoints(r2.Point{X: px1, Y: py1}, r2.Point{X: px1 + px2, Y: py1 + py2}), id: s[0]})
	}
	for i := 0; i < len(data)-1; i++ {
		free := true
		for j := i + 1; j < len(data); j++ {
			r := data[i].rect
			t := data[j].rect
			if r.InteriorIntersects(t) {
				free = false
				break
			}

		}
		if free {
			fmt.Println(data[i].id)
			return
		}
	}

}
