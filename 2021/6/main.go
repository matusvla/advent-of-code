package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "3,5,3,5,1,3,1,1,5,5,1,1,1,2,2,2,3,1,1,5,1,1,5,5,3,2,2,5,4,4,1,5,1,4,4,5,2,4,1,1,5,3,1,1,4,1,1,1,1,4,1,1,1,1,2,1,1,4,1,1,1,2,3,5,5,1,1,3,1,4,1,3,4,5,1,4,5,1,1,4,1,3,1,5,1,2,1,1,2,1,4,1,1,1,4,4,3,1,1,1,1,1,4,1,4,5,2,1,4,5,4,1,1,1,2,2,1,4,4,1,1,4,1,1,1,2,3,4,2,4,1,1,5,4,2,1,5,1,1,5,1,2,1,1,1,5,5,2,1,4,3,1,2,2,4,1,2,1,1,5,1,3,2,4,3,1,4,3,1,2,1,1,1,1,1,4,3,3,1,3,1,1,5,1,1,1,1,3,3,1,3,5,1,5,5,2,1,2,1,4,2,3,4,1,4,2,4,2,5,3,4,3,5,1,2,1,1,4,1,3,5,1,4,1,2,4,3,1,5,1,1,2,2,4,2,3,1,1,1,5,2,1,4,1,1,1,4,1,3,3,2,4,1,4,2,5,1,5,2,1,4,1,3,1,2,5,5,4,1,2,3,3,2,2,1,3,3,1,4,4,1,1,4,1,1,5,1,2,4,2,1,4,1,1,4,3,5,1,2,1"

func main() {
	vals := strings.Split(input, ",")
	var result int64
	resultsCache := make(map[int]int64)
	for _, val := range vals {
		daysToBreed, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		subresult, ok := resultsCache[daysToBreed]
		if !ok {
			subresult = countFishChildren(256, 0, daysToBreed, 6)
		}
		resultsCache[daysToBreed] = subresult
		result += subresult
		fmt.Println(result)
	}
	fmt.Println(result)
}

func countFishChildren(untilDay, currentDay, daysToBreed, defaultDaysToBreed int) int64 {
	if currentDay+daysToBreed >= untilDay {
		return 1
	}
	return countFishChildren(untilDay, currentDay+daysToBreed+1, defaultDaysToBreed, defaultDaysToBreed) + countFishChildren(untilDay, currentDay+daysToBreed+1, defaultDaysToBreed+2, defaultDaysToBreed)
}
