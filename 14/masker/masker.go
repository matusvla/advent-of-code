package masker

import (
	"fmt"
	"strings"
)

type Masker struct {
	mask, maskedFields int64
	relativeAddresses  []int64
}

func (m *Masker) SetMask(s string) {
	m.mask = 0
	m.maskedFields = 0
	binaryPos := int64(1 << (len(s) - 1))
	for _, val := range s {
		switch val {
		case '0':
			m.maskedFields += binaryPos
		case '1':
			m.mask += binaryPos
			m.maskedFields += binaryPos
		}
		binaryPos >>= 1
	}
	m.relativeAddresses = countRelativeAddresses(fmt.Sprintf(fmt.Sprintf("%%0%vb", len(s)), m.maskedFields))
}

func countRelativeAddresses(addressBinary string) []int64 {
	leftmost0 := strings.Index(addressBinary, "0")
	if leftmost0 == -1 {
		return []int64{0}
	}
	res := countRelativeAddresses(addressBinary[leftmost0+1:])
	var new []int64
	currVal := int64(1) << (len(addressBinary) - leftmost0 - 1)
	for _, val := range res {
		val += currVal
		new = append(new, val)
	}
	new = append(new, res...)
	return new
}

func (m *Masker) ApplyMask(value int64) int64 {
	return m.maskedFields&m.mask + ^m.maskedFields&value
}

func (m *Masker) AddAddresses(memMap map[int64]int64, address, value int64) {
	address = (address & m.maskedFields) | m.mask
	for _, relativeAddr := range m.relativeAddresses {
		memMap[address+relativeAddr] = value
	}
}
