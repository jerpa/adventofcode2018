package main

import (
	c "adventofcode2018/common"
)

func main() {
	sum := 0

	for _, v := range c.GetInts(c.ReadInputFile()) {
		sum += v
	}
	c.Print(sum)

}
