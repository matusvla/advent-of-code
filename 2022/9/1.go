package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2022/9/rope"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	k := rope.New(9)
	for scanner.Scan() {
		// scan line looking like this R 4
		var dir string
		var dist int
		if _, err := fmt.Sscanf(scanner.Text(), "%s %d", &dir, &dist); err != nil {
			panic(err)
		}
		k.Move(dir, dist)
	}
	fmt.Println(k.CountVisited())
}
