package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"advent-of-code-2020/10/pathfinder"
)

// from https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
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
	f, err := os.Open("./10/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ints, err := readInts(f)
	if err != nil {
		log.Fatal(err)
	}
	sort.Ints(ints)
	ints = append([]int{0}, ints...)
	ints = append(ints, ints[len(ints)-1]+3)

	//var ones, threes int
	//for i := 0; i < len(ints)-1; i++ {
	//	switch ints[i+1] - ints[i] {
	//	case 1:
	//		ones++
	//	case 3:
	//		threes++
	//	}
	//}
	pf := pathfinder.New()
	fmt.Println(pf.FindPaths(ints))
}
