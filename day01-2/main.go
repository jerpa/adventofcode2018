package main

import (
	c "adventofcode2018/common"
)

func main() {
	sum := 0
	seen := map[int]bool{}
	seen[0] = true
	for {
		for _, v := range c.GetInts(c.ReadInputFile()) {
			sum += v
			if _, exists := seen[sum]; exists {
				c.Print(sum)
				return
			}
			seen[sum] = true
		}
	}
}
