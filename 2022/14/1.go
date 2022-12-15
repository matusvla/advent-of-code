package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// todo change these
const (
	caveSize   = 5000
	sandStartX = 2500
	sandStartY = 0
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cm caveMap
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		prevPointX, prevPointY := -1, -1
		for _, s := range strings.Split(line, " -> ") {
			var x, y int
			if _, err := fmt.Sscanf(s, "%d,%d", &x, &y); err != nil {
				panic(err)
			}
			x += 2000
			if prevPointX != -1 {
				cm.drawLine(prevPointX, prevPointY, x, y)
			}
			if y > maxY {
				maxY = y
			}
			prevPointX, prevPointY = x, y
		}
	}
	// print cave map

	var result int
	for {
		if success := cm.dropSand(sandStartX, sandStartY); !success {
			break
		}
		result++
	}
	fmt.Println("PART1", result)

	// PART2
	// add floor
	for i := 0; i < caveSize; i++ {
		cm[i][maxY+2] = '#'
	}
	for {
		if success := cm.dropSand(sandStartX, sandStartY); !success {
			break
		}
		result++
	}
	fmt.Println("PART2", result)
}

type caveMap [caveSize][caveSize]byte

func (cm *caveMap) Print(maxY int) {
	for i := 0; i < maxY; i++ {
		for j := 0; j < caveSize; j++ {
			if cm[j][i] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(cm[j][i]))
			}
		}
		fmt.Println()
	}
}

func (cm *caveMap) drawLine(x int, y int, x2 int, y2 int) {
	if x == x2 {
		yMin, yMax := y, y2
		if y > y2 {
			yMin, yMax = y2, y
		}
		for i := yMin; i <= yMax; i++ {
			cm[x][i] = '#'
		}
	} else if y == y2 {
		xMin, xMax := x, x2
		if x > x2 {
			xMin, xMax = x2, x
		}
		for i := xMin; i <= xMax; i++ {
			cm[i][y] = '#'
		}
	}
}

func (cm *caveMap) dropSand(x int, y int) bool {
	if cm[x][y] != 0 {
		return false
	}
	for {
		if y == caveSize {
			return false
		}
		if cm[x][y] != 0 {
			if cm[x-1][y] == 0 {
				x = x - 1
			} else if cm[x+1][y] == 0 {
				x = x + 1
			} else {
				cm[x][y-1] = 'o'
				return true
			}
		}
		y++
	}
}
