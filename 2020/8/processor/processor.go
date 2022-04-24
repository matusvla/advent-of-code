package processor

import (
	"strconv"
	"strings"
)

type Processor struct {
	Accumulator  int
	ptr          int
	Instructions []*instruction
}

type instruction struct {
	InstrType string
	val       int
	Used      bool
}

func New() *Processor {
	return &Processor{}
}

func newInstruction(s string) *instruction {
	flds := strings.Fields(s)
	if len(flds) < 2 {
		return nil
	}
	i, err := strconv.Atoi(strings.TrimPrefix(flds[1], "+"))
	if err != nil {
		return nil
	}
	return &instruction{
		InstrType: flds[0],
		val:       i,
	}
}

func (p *Processor) Process() bool {
	p.ptr = 0
	for p.ptr < len(p.Instructions) {
		if p.ptr < 0 {
			return false
		}
		toProcess := p.Instructions[p.ptr]
		if toProcess.Used {
			return false
		}
		p.Instructions[p.ptr].Used = true
		switch toProcess.InstrType {
		case "acc":
			p.Accumulator += toProcess.val
			p.ptr++
		case "jmp":
			p.ptr += toProcess.val
		case "nop":
			p.ptr++
		}
	}
	if p.ptr == len(p.Instructions) {
		return true
	}
	return false
}

func (p *Processor) Load(s string) {
	//load
	instr := newInstruction(s)
	p.Instructions = append(p.Instructions, instr)
}

func (p *Processor) LoadAndProcess(s string) bool {
	//load
	instr := newInstruction(s)
	p.Instructions = append(p.Instructions, instr)

	//Process at pointer
	return p.Process()
}
