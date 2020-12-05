package main

import (
	"advent-of-code-2020/5/boardingpass"
	"advent-of-code-2020/5/plane"
	"bufio"
	"fmt"
	"log"
	"os"
)

const lastSeat = 127*8 + 8

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	highest := 0
	lowest := lastSeat
	p := plane.NewPlane(0, lastSeat)
	for scanner.Scan() {
		bp := scanner.Text()
		r, c := boardingpass.Translate([]byte(bp))
		id := r*8 + c
		if id > highest {
			highest = id
		}
		if id < lowest {
			lowest = id
		}
		p.MarkOccupied(id)
	}
	for _, emptySeat := range p.EmptySeats() {
		if emptySeat > lowest && emptySeat < highest {
			fmt.Printf("Result 2: %v\n", emptySeat)
			break
		}
	}
	fmt.Printf("Result 1: %v\n", highest)
}
