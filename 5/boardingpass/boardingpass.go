package boardingpass

func getValue(b []byte, bitOne byte) (result int) {
	pow2 := 1 << (len(b) - 1)
	for _, bb := range b {
		if bb == bitOne {
			result |= pow2
		}
		pow2 >>= 1
	}
	return
}

func Translate(code []byte) (row int, col int) {
	if len(code) != 10 {
		return
	}
	return getValue(code[:7], 'B'), getValue(code[7:10], 'R')
}
