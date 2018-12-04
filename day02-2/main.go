package main

import (
	c "adventofcode2018/common"
)

func main() {
	lines := c.ReadInputFile()
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if c.GetDiff(lines[i], lines[j]) == 1 {
				c.Print(c.GetSimilarity(lines[i], lines[j]))
				return
			}
		}
	}
}
