package main

import (
	c "adventofcode2018/common"
	"strings"
)

func main() {
	two, three := 0, 0

	for _, s := range c.ReadInputFile() {
		s2, s3 := false, false
		for _, v := range strings.Split(s, "") {
			if strings.Count(s, v) == 3 {
				s3 = true
				continue
			}
			if strings.Count(s, v) == 2 {
				s2 = true
			}
		}
		if s2 {
			two++
		}
		if s3 {
			three++
		}
	}
	c.Print(two, three, two*three)
}
