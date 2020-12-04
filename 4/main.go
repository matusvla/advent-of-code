package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"advent-of-code-2020/4/passport"
)

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// modified bufio.ScanLines()
func scanPassports(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(scanPassports)
	var result int
	for scanner.Scan() {
		p := scanner.Text()
		if passport.Process(p) && passport.ProcessAndValidate(p) {
			result++
		}
	}
	fmt.Printf("Result: %v\n", result)
}
