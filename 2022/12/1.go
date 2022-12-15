package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]byte
	var ex, ey, i int
	for scanner.Scan() {
		t := scanner.Text()
		if ind := strings.Index(t, "E"); ind != -1 {
			ex = i
			ey = ind
		}
		t = strings.ReplaceAll(t, "E", "z")
		t = strings.ReplaceAll(t, "S", "z")
		input = append(input, []byte(t))
		i++
	}
	fmt.Println(ex, ey)
	shortestPath[[2]int{0, 0}] = 0
	//nextStep(20, 0, input)
	nextStep2(0, 0, input)
	fmt.Println(shortestPath[[2]int{ex, ey}])

	shortestPath = make(map[[2]int]int)
	nextStep2(ex, ey, input)
	minA := 999999
	for coord, val := range shortestPath {
		x, y := coord[0], coord[1]
		if input[x][y] == 'a' && val < minA {
			minA = val
		}
	}
	fmt.Println(minA)
	//PrintShortestPath()
}

var shortestPath = map[[2]int]int{}

func nextStep(startx, starty int, all [][]byte) {
	var stack [][2]int
	stack = append(stack, [2]int{startx, starty})
	for len(stack) > 0 {
		coord := stack[0]
		x, y := coord[0], coord[1]
		shortestPathChar := shortestPath[coord]
		char := all[x][y]
		if x < len(all)-1 && all[x+1][y] <= char+1 {
			if val, ok := shortestPath[[2]int{x + 1, y}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x + 1, y}] = shortestPathChar + 1
				stack = append(stack, [2]int{x + 1, y})
			}
		}
		if y < len(all[0])-1 && all[x][y+1] <= char+1 {
			if val, ok := shortestPath[[2]int{x, y + 1}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x, y + 1}] = shortestPathChar + 1
				stack = append(stack, [2]int{x, y + 1})
			}
		}
		if x > 0 && all[x-1][y] <= char+1 {
			if val, ok := shortestPath[[2]int{x - 1, y}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x - 1, y}] = shortestPathChar + 1
				stack = append(stack, [2]int{x - 1, y})
			}
		}
		if y > 0 && all[x][y-1] <= char+1 {
			if val, ok := shortestPath[[2]int{x, y - 1}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x, y - 1}] = shortestPathChar + 1
				stack = append(stack, [2]int{x, y - 1})
			}
		}
		stack = stack[1:]
	}
}

func nextStep2(startx, starty int, all [][]byte) {
	var stack [][2]int
	stack = append(stack, [2]int{startx, starty})
	for len(stack) > 0 {
		coord := stack[0]
		x, y := coord[0], coord[1]
		shortestPathChar := shortestPath[coord]
		char := all[x][y]
		if x < len(all)-1 && all[x+1][y] >= char-1 {
			if val, ok := shortestPath[[2]int{x + 1, y}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x + 1, y}] = shortestPathChar + 1
				stack = append(stack, [2]int{x + 1, y})
			}
		}
		if y < len(all[0])-1 && all[x][y+1] >= char-1 {
			if val, ok := shortestPath[[2]int{x, y + 1}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x, y + 1}] = shortestPathChar + 1
				stack = append(stack, [2]int{x, y + 1})
			}
		}
		if x > 0 && all[x-1][y] >= char-1 {
			if val, ok := shortestPath[[2]int{x - 1, y}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x - 1, y}] = shortestPathChar + 1
				stack = append(stack, [2]int{x - 1, y})
			}
		}
		if y > 0 && all[x][y-1] >= char-1 {
			if val, ok := shortestPath[[2]int{x, y - 1}]; !ok || (ok && shortestPathChar+1 < val) {
				shortestPath[[2]int{x, y - 1}] = shortestPathChar + 1
				stack = append(stack, [2]int{x, y - 1})
			}
		}
		stack = stack[1:]
	}
}

func PrintShortestPath() {
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			fmt.Printf("%3d ", shortestPath[[2]int{i, j}])
		}
		fmt.Println()
	}
}
