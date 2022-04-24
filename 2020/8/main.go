package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/2020/8/processor"
)

func main() {
	f, err := os.Open("./8/testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	p := processor.New()
	// Part 1
	//for scanner.Scan() {
	//	if !p.LoadAndProcess(scanner.Text()) {
	//		break
	//	}
	//}
	// Part 2
	for scanner.Scan() {
		p.Load(scanner.Text())
	}

	for i := range p.Instructions {
		switch p.Instructions[i].InstrType {
		case "nop":
			p.Instructions[i].InstrType = "jmp"
		case "jmp":
			p.Instructions[i].InstrType = "nop"
		}
		if p.Process() {
			break
		}
		switch p.Instructions[i].InstrType {
		case "nop":
			p.Instructions[i].InstrType = "jmp"
		case "jmp":
			p.Instructions[i].InstrType = "nop"
		}
		for i := range p.Instructions {
			p.Instructions[i].Used = false
		}
		p.Accumulator = 0
	}
	fmt.Println(p.Accumulator)
}
