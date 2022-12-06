package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//var result1 int
	//for scanner.Scan() {
	//	t := scanner.Text()
	//	result1 += findCommon(t)
	//}
	//fmt.Println(result1)

	var result2 int
	for scanner.Scan() {
		t1 := scanner.Text()
		scanner.Scan()
		t2 := scanner.Text()
		scanner.Scan()
		t3 := scanner.Text()
		result2 += findBadge(t1, t2, t3)
	}

	fmt.Println(result2)
}

func priority[T int32 | byte](b T) int {
	if 'a' <= b && b <= 'z' {
		return int(b-'a') + 1
	}
	return int(b-'A') + 27
}

func findCommon(s string) int {
	var usedBytes int64
	for i := 0; i < len(s)/2; i++ {
		charI := priority(s[i])
		usedBytes |= 1 << charI
	}
	for i := len(s) / 2; i < len(s); i++ {
		charI := priority(s[i])
		if usedBytes&int64(1<<charI) != 0 {
			return charI
		}
	}
	panic("nothing in common")
}

func findBadge(s1, s2, s3 string) int {
	var usedBytes1 int64
	for _, b := range s1 {
		charI := priority(b)
		usedBytes1 |= 1 << charI
	}
	var usedBytes2 int64
	for _, b := range s2 {
		charI := priority(b)
		usedBytes2 |= 1 << charI
	}
	var usedBytes3 int64
	for _, b := range s3 {
		charI := priority(b)
		usedBytes3 |= 1 << charI
	}

	commonBytes := usedBytes1 & usedBytes2 & usedBytes3

	ba := fmt.Sprintf("%064b\n", commonBytes)
	// backwards for loop
	for i := len(ba) - 1; i >= 0; i-- {
		if ba[i] == '1' {
			return len(ba) - i - 2
		}
	}
	panic("nothing in common")
}
