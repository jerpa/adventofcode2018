package main

import (
	c "adventofcode2018/common"
	"strings"
)

func main() {
	data := c.ReadInputFile()[0]
	//data = "dabAcCaCBAcCcaDA"
	for {
		exists := false
		for i := 0; i < len(data)-1; i++ {

			if data[i] != data[i+1] && strings.ToLower(string(data[i])) == strings.ToLower(string(data[i+1])) {
				exists = true
				data = data[:i] + data[i+2:]
				i--
			}
		}
		if exists == false {
			c.Print(len(data))
			return
		}
	}

}
