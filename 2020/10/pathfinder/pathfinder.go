package pathfinder

type Pathfinder struct {
	foundPaths map[int]int64
}

func New() *Pathfinder {
	return &Pathfinder{
		foundPaths: make(map[int]int64),
	}
}

func (p *Pathfinder) FindPaths(sortedints []int) int64 {
	if len(sortedints) == 1 {
		return 1
	}
	if val, ok := p.foundPaths[sortedints[0]]; ok {
		return val
	}
	var pathCount int64
	for j := 1; j < len(sortedints) && sortedints[j]-sortedints[0] <= 3; j++ {
		pathCount += p.FindPaths(sortedints[j:])
	}
	p.foundPaths[sortedints[0]] = pathCount
	return pathCount
}
