package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2022/10/display"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i, val := 1, 1
	var resultset []int
	var registerHistory []int
	for scanner.Scan() {
		t := scanner.Text()
		if t == "noop" {
			i++
			if (i+20)%40 == 0 {
				resultset = append(resultset, val*i)
			}
			registerHistory = append(registerHistory, val+40*((i-1)/40))
		} else {
			var dist int
			if _, err := fmt.Sscanf(t, "addx %d", &dist); err != nil {
				panic(err)
			}
			i++
			registerHistory = append(registerHistory, val+40*((i-1)/40))
			if (i+20)%40 == 0 {
				resultset = append(resultset, val*i)
			}
			val += dist
			i++
			if (i+20)%40 == 0 {
				resultset = append(resultset, val*i)
			}
			registerHistory = append(registerHistory, val+40*((i-1)/40))
		}
	}
	fmt.Println(registerHistory)
	var sum int
	for _, r := range resultset {
		sum += r
	}
	fmt.Println("PART 1:", sum)

	d := display.New(40, 6)
	for _, r := range registerHistory {
		d.Iterate(r)
	}
	fmt.Println("PART 2:")
	d.Print()
}
