package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]int32
	for scanner.Scan() {
		var inputLine []int32
		for _, val := range scanner.Text() {
			inputLine = append(inputLine, val-'0')
		}
		input = append(input, inputLine)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 9 {
				fmt.Print(".")
			}
			fmt.Print(input[i][j])
		}
		fmt.Println()
	}

	var lowPoints []Point
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			pointVal := input[i][j]
			isValley := true
			if i > 0 {
				isValley = isValley && input[i-1][j] > pointVal
			}
			if i < len(input)-1 {
				isValley = isValley && input[i+1][j] > pointVal
			}
			if j > 0 {
				isValley = isValley && input[i][j-1] > pointVal
			}
			if j < len(input[i])-1 {
				isValley = isValley && input[i][j+1] > pointVal
			}
			if isValley {
				lowPoints = append(lowPoints, Point{i, j})
			}
		}
	}

	var results []int
	for _, lp := range lowPoints {
		m := make(map[Point]struct{})
		addSuitableNeighbours(input, lp.x, lp.y, m)
		results = append(results, len(m))
	}
	sort.Ints(results)
	l := len(results) - 1
	fmt.Println(results[l] * results[l-1] * results[l-2])
}

type Point struct {
	x, y int
}

func addSuitableNeighbours(floorplan [][]int32, i, j int, m map[Point]struct{}) {
	if floorplan[i][j] == 9 { // maybe better logic - hrbety
		return
	}
	m[Point{i, j}] = struct{}{}
	if i > 0 {
		if _, ok := m[Point{i - 1, j}]; !ok {
			addSuitableNeighbours(floorplan, i-1, j, m)
		}
	}
	if i < len(floorplan)-1 {
		if _, ok := m[Point{i + 1, j}]; !ok {
			addSuitableNeighbours(floorplan, i+1, j, m)
		}
	}
	if j > 0 {
		if _, ok := m[Point{i, j - 1}]; !ok {
			addSuitableNeighbours(floorplan, i, j-1, m)
		}
	}
	if j < len(floorplan[i])-1 {
		if _, ok := m[Point{i, j + 1}]; !ok {
			addSuitableNeighbours(floorplan, i, j+1, m)
		}
	}
}
