package graph

type Graph struct {
	nodes map[string]*Node
}

func New() *Graph {
	return &Graph{
		nodes: map[string]*Node{},
	}
}

type Node struct {
	name       string
	value      int
	neighbours map[string]int
}

func (g *Graph) AddNode(name string, value int, neighbourNames []string) {
	neighbours := map[string]int{}
	for _, neighbourName := range neighbourNames {
		neighbours[neighbourName] = 1
	}
	g.nodes[name] = &Node{
		name:       name,
		value:      value,
		neighbours: neighbours,
	}
}

// THIS IS BAD APPROACH
//
//var BestResultsCache = make(map[string][30]Cached)
//
//type Cached struct {
//	Score int
//	Used  map[string]struct{}
//}
//
//func (g *Graph) FindBest(nodename string, timeRemaining int) int {
//	for i := 1; i <= timeRemaining; i++ {
//		fmt.Print(i, " ")
//		for _, n := range g.nodes {
//			g.findBest(n.name, i)
//		}
//	}
//	return BestResultsCache[nodename][timeRemaining-1].Score
//}
//
//func (g *Graph) findBest(nodename string, timeRemaining int) {
//	if timeRemaining < 1 {
//		return // no time left to do anything useful
//	}
//	if timeRemaining == 1 {
//		nodeCache, ok := BestResultsCache[nodename]
//		if !ok {
//			BestResultsCache[nodename] = [30]Cached{}
//			nodeCache = BestResultsCache[nodename]
//		}
//		nodeCache[timeRemaining-1] = Cached{
//			0,
//			map[string]struct{}{nodename: {}},
//		}
//		BestResultsCache[nodename] = nodeCache
//		return
//	}
//	// timeRemaining > 1
//	maxContinueResult := 0
//	maxContinueUsed := make(map[string]struct{})
//	if nodeCache, ok := BestResultsCache[nodename]; ok {
//		if _, ok := nodeCache[timeRemaining-1-1].Used[nodename]; !ok {
//			maxContinueResult = nodeCache[timeRemaining-1-1].Score + g.nodes[nodename].value
//			maxContinueUsed[nodename] = struct{}{}
//		}
//	}
//	for _, neighbourName := range g.nodes[nodename].neighbourNames {
//		res := BestResultsCache[neighbourName][timeRemaining-1-1]
//		score := res.Score
//
//		if score > maxContinueResult {
//			maxContinueResult = score
//			for used := range res.Used {
//				maxContinueUsed[used] = struct{}{}
//			}
//		}
//		if timeRemaining-1-2 > 0 {
//			res = BestResultsCache[neighbourName][timeRemaining-1-2]
//			score := res.Score + g.nodes[nodename].value
//			if score > maxContinueResult {
//				for used := range res.Used {
//					maxContinueUsed[used] = struct{}{}
//
//				}
//				maxContinueUsed[nodename] = struct{}{}
//			}
//		}
//	}
//
//	generated := BestResultsCache[nodename][timeRemaining-1-1].Score
//	for n := range BestResultsCache[nodename][timeRemaining-1-1].Used {
//		generated += g.nodes[n].value
//		maxContinueUsed[n] = struct{}{}
//	}
//
//	nodeCache := BestResultsCache[nodename]
//	nodeCache[timeRemaining-1] = Cached{
//		maxContinueResult + generated,
//		maxContinueUsed,
//	}
//	BestResultsCache[nodename] = nodeCache
//}
