package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2020/18/calculator"
)

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var res int
	for scanner.Scan() {
		t := scanner.Text()
		res += calculator.EvaluateExpression2(t)
	}
	fmt.Println(res)
}
