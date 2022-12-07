package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"advent-of-code-2020/2022/7/tree"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	t := tree.New()
	// skip header
	for scanner.Scan() {
		if scanner.Text() == "$ ls" {
			for scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), "$") {
					break
				}
				var dirOrSize, name string
				if _, err := fmt.Sscanf(scanner.Text(), "%s %s", &dirOrSize, &name); err != nil {
					panic("ls scan")
				}
				size := -1
				if dirOrSize != "dir" {
					var err error
					size, err = strconv.Atoi(dirOrSize)
					if err != nil {
						panic("atoi")
					}
				}
				t.AddChildren([]tree.NodeData{
					{
						Name: name,
						Size: size,
					},
				})
			}
		}
		if strings.HasPrefix(scanner.Text(), "$ cd ") {
			dir := strings.TrimPrefix(scanner.Text(), "$ cd ")
			t.ChangeLoc(dir)
		}
	}
	t.PrettyPrint()
	sizes := t.AllSmallSizes(100000)
	fmt.Println(sizes)
	var result int
	for _, size := range sizes {
		result += size
	}
	fmt.Println(result)

	allSizes := t.AllSmallSizes(100_000_000)
	sort.Ints(allSizes)
	rootSize := allSizes[len(allSizes)-1]
	neededSpace := 70_000_000 - rootSize
	fmt.Println(allSizes, neededSpace)
	for _, s := range allSizes {
		if neededSpace+s > 30_000_000 {
			fmt.Println(s)
			break
		}
	}
}
