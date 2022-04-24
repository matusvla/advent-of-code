package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var horP, vertP int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		dir := input[0]
		howmuch, err := strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}
		switch dir {
		case "forward":
			horP += howmuch
		case "up":
			vertP -= howmuch
		case "down":
			vertP += howmuch
		default:
			panic("aaaaa")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(horP * vertP)
}
