package processor

type Processor map[int64]struct{}

func New(preamble []int64) Processor {
	p := map[int64]struct{}{}
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			p[preamble[i]+preamble[j]] = struct{}{}
		}
	}
	return p
}

func (p Processor) Validate(i int64) bool {
	_, ok := p[i]
	return ok
}
