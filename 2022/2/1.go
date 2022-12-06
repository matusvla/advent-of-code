package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A X Rock
// B Y Paper
// C Z Scissors
func resolveScore(other, me byte) int {
	var result int
	switch me {
	case 'X':
		result += 1
		switch other {
		case 'A':
			return result + 3
		case 'B':
			return result
		}
		return result + 6

	case 'Y':
		result += 2
		switch other {
		case 'B':
			return result + 3
		case 'C':
			return result
		}
		return result + 6

	case 'Z':
		result += 3
		switch other {
		case 'C':
			return result + 3
		case 'A':
			return result
		}
		return result + 6
	}
	panic("invalid input")
}

// A Rock 1
// B Paper 2
// C Scissors 3
// X Lose
// Y Draw
// Z Win

var m = map[string]int{
	"AX": 0 + 3,
	"AY": 3 + 1,
	"AZ": 6 + 2,
	"BX": 0 + 1,
	"BY": 3 + 2,
	"BZ": 6 + 3,
	"CX": 0 + 2,
	"CY": 3 + 3,
	"CZ": 6 + 1,
}

func resolveScore2(other, me byte) int {
	return m[string([]byte{other, me})]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result1, result2 int
	for scanner.Scan() {
		t := scanner.Text()
		// parse t, it looks like "A X"
		other, me := t[0], t[2]
		result1 += resolveScore(other, me)
		result2 += resolveScore2(other, me)
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
