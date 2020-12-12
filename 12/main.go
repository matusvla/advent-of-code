package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/12/ship"
	"advent-of-code-2020/12/waypoint"
)

func main() {
	f, err := os.Open("./12/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	s := ship.Ship{}
	w := waypoint.Ship{
		WaypointX: 10,
		WaypointY: 1,
	}
	for scanner.Scan() {
		s.Move(scanner.Text())
		w.Move(scanner.Text())
	}
	fmt.Println(s.ManhattanDistance())
	fmt.Println(w.ManhattanDistance())
}
