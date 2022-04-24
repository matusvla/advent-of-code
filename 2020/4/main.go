package main

import (
	"fmt"
	"log"

	"advent-of-code-2020/2020/4/passport"
	"advent-of-code-2020/2020/fileprocessing"
)

func main() {
	fp := fileprocessing.New(
		func(s string) int {
			if passport.Process(s) && passport.ProcessAndValidate(s) {
				return 1
			}
			return 0
		},
		fileprocessing.ScanBlocks,
	)

	res, err := fp.Process("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", res)
}
