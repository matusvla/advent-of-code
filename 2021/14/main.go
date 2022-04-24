package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

//func main() {
//	file, err := os.Open("./input")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	scanner.Scan()
//	polymer := scanner.Text()
//	scanner.Scan()
//
//	rules := make(map[string]string)
//	// scan points
//	for scanner.Scan() {
//		input := strings.Split(scanner.Text(), " -> ")
//		rules[input[0]] = input[1]
//	}
//
//	for i := 0; i < 40; i++ {
//		var resultPolymer strings.Builder
//		resultPolymer.WriteString(string(polymer[0]))
//		for j := 0; j < len(polymer)-1; j++ {
//			resultPolymer.WriteString(rules[polymer[j:j+2]])
//			resultPolymer.WriteString(polymer[j+1:j+2])
//		}
//		polymer = resultPolymer.String()
//		fmt.Println(i, len(polymer))
//	}
//
//	charMap := make(map[rune]int64)
//	for _, r := range polymer {
//		charMap[r]++
//	}
//
//	min := int64(math.MaxInt64)
//	max := int64(0)
//	for _, i := range charMap {
//		if i < min {
//			min = i
//		}
//		if i > max {
//			max = i
//		}
//	}
//
//	fmt.Println(max-min)
//}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	rules := make(map[string]string)
	// scan points
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " -> ")
		rules[input[0]] = input[1]
	}

	result := make(map[string]int)
	for j := 0; j < len(polymer)-1; j++ {
		subres := processPair(polymer[j:j+2], rules, 40)
		for key, val := range subres {
			result[key] += val
		}
	}
	for _, r := range polymer {
		result[string(r)]++
	}

	fmt.Println(result)

	min := math.MaxInt64
	max := 0
	for _, i := range result {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	fmt.Println(max - min)
}

var cache = make(map[string]map[int]map[string]int)

func processPair(pair string, rules map[string]string, remainingDepth int) map[string]int {
	if cacheItem, ok := cache[pair]; ok {
		if depthEntry, ok := cacheItem[remainingDepth]; ok {
			return depthEntry
		}
	} else {
		cache[pair] = make(map[int]map[string]int)
	}

	if remainingDepth == 1 {
		result := map[string]int{
			rules[pair]: 1,
		}
		cache[pair][0] = result
		return result
	}

	r := rules[pair]
	subresult1 := processPair(string(pair[0])+r, rules, remainingDepth-1)
	cache[string(pair[0])+r][remainingDepth-1] = subresult1
	subresult2 := processPair(r+string(pair[1]), rules, remainingDepth-1)
	cache[r+string(pair[1])][remainingDepth-1] = subresult2

	result := make(map[string]int)
	for key, val := range subresult1 {
		result[key] += val
	}
	for key, val := range subresult2 {
		result[key] += val
	}
	result[r]++
	cache[pair][remainingDepth] = result
	return result
}
