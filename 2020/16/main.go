package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent-of-code-2020/2020/16/ticketprocessor"
)

func main() {
	f, err := os.Open("./16/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	tp := ticketprocessor.New()
	for scanner.Scan() {
		r := make([]ticketprocessor.Range, 2)
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) < 2 {
			break
		}
		_, err := fmt.Sscanf(parts[1], "%d-%d or %d-%d", &r[0][0], &r[0][1], &r[1][0], &r[1][1])
		if err != nil {
			panic(err)
		}
		tp.AddRule(parts[0], r...)
	}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "your ticket:") {
			break
		}
	}

	scanner.Scan()
	t := scanner.Text()
	var myticket []int
	for _, v := range strings.Split(t, ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		myticket = append(myticket, i)
	}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "nearby tickets:") {
			break
		}
	}
	var result1 int
	var validTickets [][]int
	for scanner.Scan() {
		t := scanner.Text()
		var input []int
		for _, v := range strings.Split(t, ",") {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			input = append(input, i)
		}

		invalidVals, valid := tp.FilterInvalid(input)
		for _, v := range invalidVals {
			result1 += v
		}
		if valid {
			validTickets = append(validTickets, input)
		}
	}

	var invalid [20][]int
	for _, t := range validTickets {
		for j, v := range t {
			if !((46 <= v && v <= 893) || (913 <= v && v <= 964)) {
				invalid[j] = append(invalid[j], v*1000+j)
			}
		}
	}
	fmt.Println(invalid)

	res := tp.FindPossibleMapping(validTickets)
	uniqres := tp.RemapToUnique(res)

	fmt.Println(result1)

	fmt.Println(res)
	fmt.Println(uniqres)

	finres := 1
	for key, index := range uniqres {
		if strings.Contains(key, "departure") {
			finres *= myticket[index]
		}
	}
	fmt.Println(finres)
}
