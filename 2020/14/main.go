package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2020/14/masker"
)

func main() {
	f, err := os.Open("./14/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var m masker.Masker
	data := make(map[int64]int64)
	// Part1
	//for scanner.Scan() {
	//	var maskValue string
	//	n, _ := fmt.Sscanf(scanner.Text(), "mask = %s", &maskValue)
	//	if n > 0 {
	//		m.SetMask(maskValue)
	//		continue
	//	}
	//
	//	var index, value int64
	//	n, err := fmt.Sscanf(scanner.Text(), "mem[%d] = %d", &index, &value)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	data[index] = m.ApplyMask(value)
	//}

	//Part2
	for scanner.Scan() {
		var maskValue string
		n, _ := fmt.Sscanf(scanner.Text(), "mask = %s", &maskValue)
		if n > 0 {
			m.SetMask(maskValue)
			continue
		}

		var address, value int64
		n, err := fmt.Sscanf(scanner.Text(), "mem[%d] = %d", &address, &value)
		if err != nil {
			log.Fatal(err)
		}
		m.AddAddresses(data, address, value)
	}

	var result int64
	for _, val := range data {
		result += val
	}
	fmt.Println(result)
}
