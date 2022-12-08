package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2022/8/forest"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	f := forest.New()
	for scanner.Scan() {
		f.AddTreeLine(scanner.Text())
	}
	fmt.Println("PART 1:", f.CountVisibleTrees())
	fmt.Println("PART 2", f.FindBestScenicScore())
}
