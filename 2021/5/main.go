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

	plan := NewPlan()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line Line
		scanner.Text()
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &line.start.x, &line.start.y, &line.end.x, &line.end.y)
		if err != nil {
			panic(err)
		}
		plan.Add(line)
	}
	fmt.Println(plan.Count(2))
}
