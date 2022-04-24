package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2020/7/graph"
)

func main() {
	f, err := os.Open("./7/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	d := graph.New()
	for scanner.Scan() {
		d.AddNode(scanner.Text())
	}
	fmt.Println(len(d.FindContainers("shiny gold", nil, 0)))
	fmt.Println(d.CountBags("shiny gold") - 1)
}
