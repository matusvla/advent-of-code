package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var wordsToDisplay = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		row := scanner.Text()
		rowSplit := strings.Split(row, "|")
		input := strings.Fields(rowSplit[0])
		output := strings.Fields(rowSplit[1])

		countM := make(map[string]int)
		var one, seven, four string
		for _, word := range input {
			for _, letter := range word {
				countM[string(letter)]++
			}
			switch len(word) {
			case 2:
				one = word
			case 3:
				seven = word
			case 4:
				four = word
			}
		}

		var keyMap = map[int]string{
			6: "b",
			4: "e",
			9: "f",
		}
		valM := make(map[string]string)

		var sevenCounts []string
		for codedKey, val := range countM {
			if val == 7 {
				sevenCounts = append(sevenCounts, codedKey)
			}
			if key, ok := keyMap[val]; ok {
				valM[codedKey] = key
			}
		}

		// one
		if _, ok := valM[string(one[0])]; !ok {
			valM[string(one[0])] = "c"
		} else {
			valM[string(one[1])] = "c"
		}

		// seven
		valM[strings.ReplaceAll(strings.ReplaceAll(seven, string(one[0]), ""), string(one[1]), "")] = "a"

		// four
		unusedFour := strings.ReplaceAll(strings.ReplaceAll(four, string(one[0]), ""), string(one[1]), "")
		if _, ok := valM[string(unusedFour[0])]; !ok {
			valM[string(unusedFour[0])] = "d"
		} else {
			valM[string(unusedFour[1])] = "d"
		}

		if _, ok := valM[sevenCounts[0]]; !ok {
			valM[sevenCounts[0]] = "g"
		} else {
			valM[sevenCounts[1]] = "g"
		}

		var realVals []string
		for _, outputVal := range output {
			var realVal string
			for _, letter := range outputVal {
				newLetter, ok := valM[string(letter)]
				if !ok {
					panic("letter not found")
				}
				realVal += newLetter
			}
			sorted := []rune(realVal)
			sort.Slice(sorted, func(i int, j int) bool { return sorted[i] < sorted[j] })
			realVals = append(realVals, string(sorted))
		}
		var fullN int
		for i, word := range realVals {
			n, ok := wordsToDisplay[word]
			if !ok {
				panic("not found in dict")
			}
			for j := 0; j < len(realVals)-1-i; j++ {
				n *= 10
			}
			fullN += n
		}
		fmt.Println(fullN)
		result += fullN
	}
	fmt.Println(result)
}
