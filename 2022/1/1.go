package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var best, second, third, currentSum int
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if currentSum > best {
				best, second, third = currentSum, best, second
			} else if currentSum > second {
				second, third = currentSum, second
			} else if currentSum > third {
				third = currentSum
			}
			currentSum = 0
			continue
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		currentSum += i
	}
	if currentSum > best {
		best, second, third = currentSum, best, second
	} else if currentSum > second {
		second, third = currentSum, second
	} else if currentSum > third {
		third = currentSum
	}
	fmt.Println(best + second + third)
}
