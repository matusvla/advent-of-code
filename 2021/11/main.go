package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		i   int
		jfs AllJellyFish
	)
	for scanner.Scan() {
		jfs.rowLen = len(scanner.Text())
		for j, jfVal := range scanner.Text() {
			jf := Jellyfish{
				value:   int(jfVal - '0'),
				posX:    j,
				posY:    i,
				flashed: false,
			}
			jfs.PushBack(&jf)
		}
		i++
	}

	// Part 1
	//var result int
	//for i := 0; i < 100; i++ {
	//	result += jfs.NextRound()
	//	//jfs.Print()
	//}
	//fmt.Println(result)

	// Part 2
	var firstAllFlash int
	for {
		firstAllFlash++
		flashedCount := jfs.NextRound()
		if flashedCount == len(jfs.jf) {
			break
		}
		//jfs.Print()
	}
	fmt.Println(firstAllFlash)
}
