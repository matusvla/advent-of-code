package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"advent-of-code-2020/2020/9/processor"
)

func main() {
	f, err := os.Open("./9/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var rowsScanned int
	var loaded []int64
	var invalidNo int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if invalidNo == 0 {
			if rowsScanned >= 25 {
				p := processor.New(loaded[len(loaded)-25:])
				if !p.Validate(i) {
					invalidNo = i
				}
			}
		}
		loaded = append(loaded, i)
		rowsScanned++

	}
	fmt.Println(invalidNo)
	for i := 0; i < len(loaded)-1; i++ {
		sum := loaded[i]
		min := loaded[i]
		max := loaded[i]
		for j := i + 1; j < len(loaded); j++ {
			sum += loaded[j]
			if sum == invalidNo {
				fmt.Println(min + max)
				return
			}
			if sum > invalidNo {
				break
			}
			if loaded[j] > max {
				max = loaded[j]
			}
			if loaded[j] < min {
				min = loaded[j]
			}
		}
	}
}
