package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	drawnNumbersStr := scanner.Text()
	drawnNumbersStrSeq := strings.Split(drawnNumbersStr, ",")
	var drawnNumbers []int
	for _, val := range drawnNumbersStrSeq {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		drawnNumbers = append(drawnNumbers, i)
	}
	bs := NewBingoSolver(drawnNumbers)

Outer:
	for {
		scanner.Scan() // skip empty line at the beginning
		var bingoVals [5][5]int
		for i := 0; i < 5; i++ {
			if !scanner.Scan() {
				break Outer
			}
			row := scanner.Text()
			rowStrVals := strings.Fields(row)
			for j, val := range rowStrVals {
				intVal, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				bingoVals[i][j] = intVal
			}
		}
		b := NewBingo(bingoVals)
		bs.Solve(b)
	}
	fmt.Println(bs.Best())
}
