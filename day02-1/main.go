package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	two := 0
	three := 0
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		s2 := false
		s3 := false
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
	fmt.Println(three * two)
}
