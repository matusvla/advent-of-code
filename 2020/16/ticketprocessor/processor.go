package ticketprocessor

type TicketProcessor struct {
	validationFuncs map[string]func(int) bool
}

type Range [2]int

func New() *TicketProcessor {
	return &TicketProcessor{
		validationFuncs: make(map[string]func(int) bool),
	}
}

func (t *TicketProcessor) AddRule(fldName string, rs ...Range) {
	t.validationFuncs[fldName] = func(i int) bool {
		valid := false
		for _, r := range rs {
			if i >= r[0] && i <= r[1] {
				valid = true
				break
			}
		}
		return valid
	}
}

func (t *TicketProcessor) FilterInvalid(is []int) ([]int, bool) {
	var result []int
	validLine := true
	for _, val := range is {
		var valid bool
		for _, f := range t.validationFuncs {
			if f(val) {
				valid = true
				break
			}
		}
		if !valid {
			result = append(result, val)
			validLine = false
		}
	}
	return result, validLine
}

func (t *TicketProcessor) FindPossibleMapping(in [][]int) map[string][]int {
	possiblePlaces := make(map[string]map[int]bool)
	for key := range t.validationFuncs {
		possiblePlaces[key] = make(map[int]bool)
	}
	for _, ticket := range in {
		for i, fld := range ticket {
			for key, f := range t.validationFuncs {
				isValid := f(fld)
				if _, ok := possiblePlaces[key][i]; isValid && !ok {
					possiblePlaces[key][i] = true
				} else if !isValid {
					possiblePlaces[key][i] = false
				}
			}
		}
	}

	result := make(map[string][]int)
	for key, p := range possiblePlaces {
		for i, val := range p {
			if val {
				result[key] = append(result[key], i)
			}
		}
	}
	return result
}

func (t *TicketProcessor) RemapToUnique(mapping map[string][]int) map[string]int {
	result := make(map[string]int)
	used := make(map[int]string)
	for len(mapping) > len(used) {
		for key, val := range mapping {
			unused := -1
			for _, v := range val {
				if _, ok := used[v]; !ok {
					if unused != -1 {
						unused = -1
						break
					}
					unused = v
				}
			}
			if unused != -1 {
				used[unused] = key
				result[key] = unused
			}
		}
	}
	return result
}
