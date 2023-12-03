package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		t := scanner.Text()
		total += lineValue2(t)
	}
	fmt.Println(total)
}

var numberMap = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func lineValue(s string) int {
	firstN := -1
	lastN := -1
	for _, b := range s {
		if b < '0' || b > '9' {
			continue
		}
		if firstN < 0 {
			firstN = int(b - '0')
		}
		lastN = int(b - '0')
	}
	return 10*firstN + lastN
}

func lineValue2(s string) int {
	firstN := 0
	lastN := 0
	firstI := len(s) + 1
	lastI := -1
	for key, val := range numberMap {
		first := strings.Index(s, key)
		last := strings.LastIndex(s, key)
		if first != -1 && first < firstI {
			firstN = val
			firstI = first
		}
		if last > lastI {
			lastN = val
			lastI = last
		}
	}
	return 10*firstN + lastN
}
