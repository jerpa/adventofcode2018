package main

import (
	c "adventofcode2018/common"
	"strings"
)

func main() {
	bestlen := 100000
	chars := "abcdefghijklmnopqrstuvwxyz"
	for x := 0; x < len(chars); x++ {
		data := c.ReadInputFile()[0]
		data = strings.Replace(data, string(chars[x]), "", -1)
		data = strings.Replace(data, strings.ToUpper(string(chars[x])), "", -1)
		c.Print(string(chars[x]))
		//data = "dabAcCaCBAcCcaDA"
		for {
			exists := false
			for i := len(data) - 2; i >= 0; i-- {

				if data[i] != data[i+1] && strings.ToLower(string(data[i])) == strings.ToLower(string(data[i+1])) {

					exists = true
					data = data[:i] + data[i+2:]
					i--
				}
			}
			if exists == false {
				if len(data) < bestlen {
					bestlen = len(data)
				}
				break

			}
		}
	}
	c.Print(bestlen)
}
