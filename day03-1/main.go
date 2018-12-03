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

func main() {
	data := map[r2.Point]int{}

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
		for x := px1; x < px1+px2; x++ {
			for y := py1; y < py1+py2; y++ {
				p := r2.Point{X: x, Y: y}
				if _, exists := data[p]; !exists {
					data[p] = 1
				} else {
					data[p] = 2
				}
			}
		}
	}
	sum := 0
	for _, v := range data {
		if v == 2 {
			sum++
		}
	}
	fmt.Println(sum)
}
