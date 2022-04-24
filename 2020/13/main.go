package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"advent-of-code-2020/2020/13/crt"
)

func main() {
	f, err := ioutil.ReadFile("./13/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)
	data := strings.Split(input, "\n")
	time, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	var min, bestLine int
	min = 9000000000
	busLines := strings.Split(data[1], ",")
	for _, line := range busLines {
		if line == "x" {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if time%id == 0 || id-time%id < min {
			min = id - time%id
			bestLine = id
		}
	}
	fmt.Println(min * bestLine)

	// Prepare variables for Chinese remainder theorem
	//var x int64
	var a, n []int
	for i, line := range busLines {
		id, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		a = append(a, id-(i-i/id*id))
		n = append(n, id)
	}

	val, err := crt.ChineseRemainderTheorem(a, n)
	fmt.Println(val, err)
}
