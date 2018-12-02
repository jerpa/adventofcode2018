package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := []string{}
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			diff := 0
			s1 := strings.Split(lines[i], "")
			s2 := strings.Split(lines[j], "")

			for p := 0; p < len(s1) && p < len(s2); p++ {
				if s1[p] != s2[p] {
					diff++
				}
			}
			if diff == 1 {
				for p := 0; p < len(s1) && p < len(s2); p++ {
					if s1[p] == s2[p] {
						fmt.Print(s1[p])
					}
				}
				fmt.Println("")
				return
			}
		}
	}
}
