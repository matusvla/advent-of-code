package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2022/5/stack"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stackStrings := [9]string{
		"HRBDZFLS",
		"TBMRZ",
		"ZLCHNS",
		"SCFJ",
		"PGHWRZB",
		"VJZGDNMT",
		"GLNWFSPQ",
		"MZR",
		"MCLGVRT",
	}
	//stackStrings := [3]string{
	//	"ZN",
	//	"MCD",
	//	"P",
	//}
	// skip header
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
	}
	var stacks [9]*stack.IndexedStack
	for i, s := range stackStrings {
		stacks[i] = stack.New(s, false)
	}
	for scanner.Scan() {
		// scan digits from "move n from n to n"
		var count, from, to int
		if _, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to); err != nil {
			panic(err)
		}
		popped := stacks[from-1].Pop(count)
		stacks[to-1].Push(popped)

		fmt.Println(scanner.Text())
		for _, s := range stacks {
			fmt.Println(s.String())
		}
		fmt.Println("--------------------------------")
	}

	for _, s := range stacks {
		fmt.Print(string(s.Top()))
	}
}
