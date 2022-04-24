package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	ox := filter(input, 0, true)
	co := filter(input, 0, false)
	fmt.Println(ox * co)
}

func filter(ss []string, i int, oxygen bool) int64 {
	if len(ss) == 1 {
		if j, err := strconv.ParseInt(ss[0], 2, 64); err != nil {
			panic(err)
		} else {
			return j
		}
	}
	var zeroSS, oneSS []string
	for _, val := range ss {
		if val[i] == '1' {
			oneSS = append(oneSS, val)
		} else {
			zeroSS = append(zeroSS, val)
		}
	}
	if oxygen {
		if len(oneSS) >= len(zeroSS) {
			return filter(oneSS, i+1, oxygen)
		} else if len(oneSS) < len(zeroSS) {
			return filter(zeroSS, i+1, oxygen)
		}
	} else {
		if len(oneSS) < len(zeroSS) {
			return filter(oneSS, i+1, oxygen)
		} else if len(oneSS) >= len(zeroSS) {
			return filter(zeroSS, i+1, oxygen)
		}
	}
	panic(1234)
}
