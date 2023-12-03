package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"advent-of-code-2020/2022/16/graph"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	g := graph.New()
	for scanner.Scan() {
		var name string
		var value int
		var neighbourNamesStr string
		reformatT := strings.ReplaceAll(scanner.Text(), ", ", ",")
		reformatT = strings.ReplaceAll(reformatT, "valves", "valve")
		reformatT = strings.ReplaceAll(reformatT, "tunnels", "tunnel")
		reformatT = strings.ReplaceAll(reformatT, "leads", "lead")
		fmt.Sscanf(
			reformatT,
			"Valve %s has flow rate=%d; tunnel lead to valve %v",
			&name,
			&value,
			&neighbourNamesStr,
		)
		neighbourNames := strings.Split(neighbourNamesStr, ",")
		g.AddNode(name, value, neighbourNames)
	}

	fmt.Println(g.FindBest("AA", 30))
	for key, val := range graph.BestResultsCache {
		fmt.Println(key)
		for _, v := range val {
			fmt.Printf("%d ", v.Score)
		}
		fmt.Println()
	}
}
