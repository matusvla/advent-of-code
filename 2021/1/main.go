package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var res []int
	var count int
	for scanner.Scan() {
		count++
		for i, ch := range scanner.Text() {
			if ch == '1' {
				res[i]++
			}
		}
	}
	var gamma int
	for i, r := range res {
		if r > count/2 {
			gamma += 1 << i
		}
	}
	fmt.Println(gamma * ^gamma)
}
