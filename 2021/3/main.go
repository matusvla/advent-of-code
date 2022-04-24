package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var res = make([]uint, 12)
	var count uint
	for scanner.Scan() {
		count++
		text := scanner.Text()
		for i, ch := range text {
			if ch == '1' {
				res[i]++
			}
		}
	}
	var gamma, epsilon int64
	for i, r := range res {
		if r > count/2 {
			gamma += 1 << (len(res) - 1 - i)
		} else {
			epsilon += 1 << (len(res) - 1 - i)
		}
	}
	fmt.Printf(strconv.FormatInt(gamma, 2))
	fmt.Println()
	fmt.Printf(strconv.FormatInt(epsilon, 2))
	fmt.Println()
	fmt.Println(gamma * epsilon)
}
