package main

import (
	"advent-of-code-2020/6/answers"
	"advent-of-code-2020/fileprocessing"
	"fmt"
	"log"
)

func main() {
	fp := fileprocessing.New(
		//answers.Union,
		answers.Intersection,
		fileprocessing.ScanBlocks,
	)

	res, err := fp.Process("./6/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", res)
}
