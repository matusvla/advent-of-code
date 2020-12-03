package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"advent-of-code-2020/1/finder"
)

// from https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ints, err := readInts(f)
	if err != nil {
		log.Fatal(err)
	}
	result1, _ := finder.FindAndMultiply(ints, 2020, 2)
	fmt.Printf("Result 1: %v\n", result1)
	result2, _ := finder.FindAndMultiply(ints, 2020, 3)
	fmt.Printf("Result 2: %v\n", result2)
}
