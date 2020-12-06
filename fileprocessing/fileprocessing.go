package fileprocessing

import (
	"bufio"
	"bytes"
	"os"
)

type FileProcessor struct {
	processFunc func(dataBlock string) int
	splitFunc   bufio.SplitFunc
}

func New(pf func(dataBlock string) int, sf bufio.SplitFunc) FileProcessor {
	return FileProcessor{
		processFunc: pf,
		splitFunc:   sf,
	}
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// modified bufio.ScanLines()
func ScanBlocks(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func (fp FileProcessor) Process(filePath string) (i int, rerr error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(fp.splitFunc)

	var result int
	for scanner.Scan() {
		result += fp.processFunc(scanner.Text())
	}

	return result, nil
}
