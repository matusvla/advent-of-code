package finder

func findAndMultiplyTwo(expReport []int, sumResult int) (int, bool) {
	m := map[int]struct{}{}
	for _, item := range expReport {
		if _, ok := m[item]; ok {
			return item * (sumResult - item), true
		}
		m[sumResult-item] = struct{}{}
	}
	return 0, false
}

func FindAndMultiply(expReport []int, sumResult int, numberCount int) (int, bool) {
	if numberCount < 2 {
		return 0, false
	} else if numberCount == 2 {
		return findAndMultiplyTwo(expReport, sumResult)
	}
	for i, val := range expReport[:len(expReport)-numberCount+1] {
		res, found := FindAndMultiply(expReport[i+1:], sumResult-val, numberCount-1)
		if found {
			return val * res, true
		}
	}
	return 0, false
}
