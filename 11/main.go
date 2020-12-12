package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/11/seatfiller"
)

func main() {
	f, err := os.Open("./11/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var loaded [][]byte
	for scanner.Scan() {
		loaded = append(loaded, scanner.Bytes())
	}

	changes := 1
	for changes > 0 {
		loaded, changes = seatfiller.FillSeats2(loaded)
		//fmt.Println(changes)
	}
	var result int
	for _, val := range loaded {
		result += bytes.Count(val, []byte{'#'})
		fmt.Println(string(val))
	}
	fmt.Println(result)

}
