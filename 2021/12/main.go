package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cave struct {
	ID         string
	visited    int
	neighbours []*Cave
}

type CaveSystem struct {
	caves map[string]*Cave // CaveID list
}

func NewCaveSystem() *CaveSystem {
	return &CaveSystem{caves: make(map[string]*Cave)}
}

func (s *CaveSystem) Add(start string, end string) {
	startCave, ok := s.caves[start]
	if !ok {
		startCave = &Cave{
			ID: start,
		}
	}
	endCave, ok := s.caves[end]
	if !ok {
		endCave = &Cave{
			ID: end,
		}
	}
	startCave.neighbours = append(startCave.neighbours, endCave)
	endCave.neighbours = append(endCave.neighbours, startCave)
	s.caves[start] = startCave
	s.caves[end] = endCave
}

func (s *CaveSystem) AllPaths(startID, doubledSmallCaveID string, path []string, pathMap map[string]struct{}) {
	newPath := append(path, startID)
	if startID == "end" {
		pathMap[fmt.Sprint(path)] = struct{}{}
		return
	}
	start := s.caves[startID]
	if start.visited > 0 && strings.ToLower(startID) == startID {
		if !(startID == doubledSmallCaveID && start.visited == 1) {
			return
		}
	}
	start.visited += 1
	for _, neighbor := range start.neighbours {
		s.AllPaths(neighbor.ID, doubledSmallCaveID, newPath, pathMap)
	}
	start.visited -= 1
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caveSystem := NewCaveSystem()
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "-")
		caveSystem.Add(input[0], input[1])
	}

	//fmt.Println("PART1")
	//fmt.Println(caveSystem.AllPathsCount("start", "", []string{}))

	var smallCaveIDs []string
	for _, cave := range caveSystem.caves {
		if cave.ID == strings.ToLower(cave.ID) && cave.ID != "start" && cave.ID != "end" {
			smallCaveIDs = append(smallCaveIDs, cave.ID)
		}
	}

	resultMap := make(map[string]struct{})
	for _, smallCaveID := range smallCaveIDs {
		caveSystem.AllPaths("start", smallCaveID, []string{}, resultMap)
	}

	fmt.Println("PART2")
	fmt.Println(len(resultMap))
}
