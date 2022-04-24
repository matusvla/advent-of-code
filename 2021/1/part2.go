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
	var count int
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	j, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		if i+j+k < j+k+l {
			count++
		}
		i, j, k = j, k, l
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
