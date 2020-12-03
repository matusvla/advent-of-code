package main

import (
	"advent-of-code-2020/3/tobogan"

	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lineNo int
	var results [5]int
	steps := [5]float32{1, 3, 5, 7, 0.5}
	for scanner.Scan() {
		line := scanner.Text()
		for i, step := range steps {
			valid := tobogan.CheckTree(line, lineNo, step)
			if valid {
				results[i]++
			}
		}
		lineNo++
	}
	fmt.Printf("Results: %v\n", results)
	finalResult := 1
	for _, res := range results {
		finalResult *= res
	}
	fmt.Printf("Multiplied: %v\n", finalResult)
}
