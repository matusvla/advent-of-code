package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var result int
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
				result += int(pointVal) + 1
			}
		}
	}

	fmt.Println(result)
}
