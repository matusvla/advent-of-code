package main

import "fmt"

func main() {
	lastseen := make(map[int]int)
	input := []int{6, 4, 12, 1, 20, 0, 16}
	for i, val := range input[:len(input)-1] {
		lastseen[val] = i
	}
	prev := input[len(input)-1]
	for i := len(input); i < 30000000; i++ {
		index, ok := lastseen[prev]
		lastseen[prev] = i - 1
		if ok {
			prev = i - 1 - index
		} else {
			prev = 0
		}
	}
	fmt.Println(prev)
}
