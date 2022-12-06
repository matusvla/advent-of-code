package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Part1
	var ch1, ch2, ch3, ch4 byte
	for i, ch := range input {
		ch1, ch2, ch3, ch4 = ch2, ch3, ch4, ch
		if i >= 3 && ch1 != ch2 && ch1 != ch3 && ch1 != ch4 && ch2 != ch3 && ch2 != ch4 && ch3 != ch4 {
			fmt.Println(i + 1)
			break
		}
	}

	// Part2
	var currentMsg string
	for i, b := range input {
		for j, c := range currentMsg {
			if b == byte(c) {
				currentMsg = currentMsg[j+1:]
			}
		}
		currentMsg += string(b)
		if len(currentMsg) == 14 {
			fmt.Println(i + 1)
			break
		}
	}
}
