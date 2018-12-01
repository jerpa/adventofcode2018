package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum := 0
	seen := map[int]bool{}
	seen[0] = true
	for {
		file, err := os.Open("./input")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err.Error())
			}
			sum += val
			if _, exists := seen[sum]; exists {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
		file.Close()
	}

}
