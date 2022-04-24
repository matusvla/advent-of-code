package main

import (
	"advent-of-code-2020/2020/2/parser"

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
	var result int
	for scanner.Scan() {
		unparsedLine := scanner.Text()
		valid, err := parser.ValidateLine(unparsedLine, parser.NewPolicy)
		if err != nil {
			log.Fatal(err)
		}
		if valid {
			result++
		}
	}
	fmt.Printf("Result old: %v\n", result)
}
