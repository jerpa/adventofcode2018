package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type entry struct {
	start time.Time
	text  string
}

func main() {
	data := []entry{}
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		e, _ := time.Parse("2006-01-02 15:04", s[1:17])
		data = append(data, entry{start: e, text: s[19:]})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].start.Before(data[j].start)
	})
	currentguard := ""
	var asleep int
	sleepsum := map[string]int{}
	sleeptime := map[string]map[int]int{}
	for _, v := range data {
		if v.text[:5] == "Guard" {
			i := strings.Index(v.text[7:], " ")
			currentguard = v.text[7 : i+7]
		} else if v.text == "falls asleep" {
			asleep = v.start.Minute()
		} else if v.text == "wakes up" {
			sleepsum[currentguard] += v.start.Minute() - asleep
			for t := asleep; t < v.start.Minute(); t++ {
				if sleeptime[currentguard] == nil {
					sleeptime[currentguard] = map[int]int{}
				}
				sleeptime[currentguard][t]++
			}
		}
	}
	max := 0

	for k, v := range sleepsum {
		if v > max {
			max = v
			currentguard = k
		}
	}

	maxtime := 0
	minute := -1
	for k, v := range sleeptime[currentguard] {
		if v > maxtime {
			maxtime = v
			minute = k
		}
	}

	fmt.Println(currentguard, minute)
}
