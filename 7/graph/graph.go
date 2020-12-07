package graph

import (
	"regexp"
	"strconv"
	"strings"
)

type Directed struct {
	predecessors, followers map[string][]nodePath
}

type nodePath struct {
	nodeID  string
	edgeVal int
}

func New() *Directed {
	return &Directed{
		predecessors: make(map[string][]nodePath),
		followers:    make(map[string][]nodePath),
	}
}

//1 dark olive bag
var itemRe = regexp.MustCompile(`(?:(\d+|no) )?([a-z ]*) bags?.?`)

func extractNodeInfo(s string) (int, string) {
	res := itemRe.FindStringSubmatch(s)
	if len(res) < 3 {
		return 0, ""
	}
	i, _ := strconv.Atoi(res[1])
	return i, res[2]
}

func (d *Directed) AddNode(nodeDescription string) {
	bagDesc := strings.Split(nodeDescription, " contain ")
	container := bagDesc[0]
	containees := bagDesc[1]
	_, startNode := extractNodeInfo(container)

	for _, v := range strings.Split(containees, ",") {
		edgeVal, nodeID := extractNodeInfo(v)
		d.followers[nodeID] = append(d.followers[nodeID], nodePath{
			nodeID:  startNode,
			edgeVal: edgeVal,
		})
		d.predecessors[startNode] = append(d.predecessors[startNode], nodePath{
			nodeID:  nodeID,
			edgeVal: edgeVal,
		})
	}
}

func (d *Directed) FindContainers(nodeID string, used []string, depth int) []string {
	used = append(used, nodeID)
	var result []string
	if depth > 0 {
		result = append(result, nodeID)
	}

	for _, neighbour := range d.followers[nodeID] {
		var alreadyUsed bool
		for _, u := range used {
			if neighbour.nodeID == u {
				alreadyUsed = true
				break
			}
		}
		if !alreadyUsed {
			deeperUsed := d.FindContainers(neighbour.nodeID, used, depth+1)
			used = append(used, deeperUsed...)
			result = append(result, deeperUsed...)
		}
	}
	return result
}

func (d *Directed) CountBags(nodeID string) int {
	var result = 1

	for _, neighbour := range d.predecessors[nodeID] {
		result += neighbour.edgeVal * d.CountBags(neighbour.nodeID)
	}
	return result
}
