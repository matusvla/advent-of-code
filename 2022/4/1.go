package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result1, result2 int
	for scanner.Scan() {
		t1 := scanner.Text()
		var a1, a2, b1, b2 int
		if _, err := fmt.Sscanf(t1, "%d-%d,%d-%d", &a1, &a2, &b1, &b2); err != nil {
			panic(err)
		}
		if contains(a1, a2, b1, b2) {
			result1++
		}
		if hasIntersect(a1, a2, b1, b2) {
			result2++
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func contains(a1, a2, b1, b2 int) bool {
	return (a1 <= b1 && a2 >= b2) || (a1 >= b1 && a2 <= b2)
}

func hasIntersect(a1, a2, b1, b2 int) bool {
	return (b1 >= a1 && b1 <= a2) ||
		(b2 >= a1 && b2 <= a2) ||
		(a1 >= b1 && a1 <= b2) ||
		(a2 >= b1 && a2 <= b2)
}
