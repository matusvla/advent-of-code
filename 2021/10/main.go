package main

import (
	"bufio"
	"container/list"
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
	var result int
	var notCorruptedRows []string
	for scanner.Scan() {
		t := scanner.Text()
		errCode := rowErrCode(t)
		result += errCode
		if errCode == 0 {
			notCorruptedRows = append(notCorruptedRows, t)
		}
	}
	fmt.Printf("Part1 result: %d\n", result)

	var rowCompletionCodes []int
	for _, row := range notCorruptedRows {
		rcc := rowCompletionCode(row)
		rowCompletionCodes = append(rowCompletionCodes, rcc)
	}
	sort.Ints(rowCompletionCodes)
	fmt.Printf("Part2 result: %d\n", rowCompletionCodes[(len(rowCompletionCodes)-1)/2])
}

// PART 2

func rowCompletionCode(row string) int {
	stack := list.New()
	for _, b := range row {
		switch b {
		case '(', '[', '{', '<':
			stack.PushBack(b)
		case ')', ']', '}', '>':
			stackTop := stack.Back()
			stack.Remove(stackTop)
		}
	}
	var result int
	for {
		stackTop := stack.Back()
		if stackTop == nil {
			break
		}
		stack.Remove(stackTop)
		switch stackTop.Value.(rune) {
		case '(':
			result *= 5
			result += 1
		case '[':
			result *= 5
			result += 2
		case '{':
			result *= 5
			result += 3
		case '<':
			result *= 5
			result += 4
		default:
			panic("unexpected char")
		}
	}
	return result
}

// PART 1

const (
	roundBracketErr  = 3
	squareBracketErr = 57
	braceErr         = 1197
	pointyBracketErr = 25137
)

func rowErrCode(row string) int {
	stack := list.New()
	for _, b := range row {
		switch b {
		case '(', '[', '{', '<':
			stack.PushBack(b)
		case ')':
			stackTop := stack.Back()
			stack.Remove(stackTop)
			if stackTop.Value.(rune) != '(' {
				return roundBracketErr
			}
		case ']':
			stackTop := stack.Back()
			stack.Remove(stackTop)
			if stackTop.Value.(rune) != '[' {
				return squareBracketErr
			}
		case '}':
			stackTop := stack.Back()
			stack.Remove(stackTop)
			if stackTop.Value.(rune) != '{' {
				return braceErr
			}
		case '>':
			stackTop := stack.Back()
			stack.Remove(stackTop)
			if stackTop.Value.(rune) != '<' {
				return pointyBracketErr
			}
		}
	}
	return 0
}
