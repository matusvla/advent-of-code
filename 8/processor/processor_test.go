package processor

import (
	"fmt"
	"testing"
)

func TestProcessor_Process(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	p := New()
	for _, val := range input {
		p.Load(val)
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
