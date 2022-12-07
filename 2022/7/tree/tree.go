package tree

import "fmt"

type NodeData struct {
	Name string
	Size int
}

type Node struct {
	NodeData
	parent   *Node
	children []*Node
}

type IndexedTree struct {
	root       *Node
	currentLoc *Node
}

func New() *IndexedTree {
	root := Node{
		NodeData: NodeData{
			Name: "/",
			Size: -1,
		},
	}
	return &IndexedTree{
		root:       &root,
		currentLoc: &root,
	}
}

func (it *IndexedTree) AddChildren(children []NodeData) {
	for _, child := range children {
		ch := Node{
			NodeData: child,
			parent:   it.currentLoc,
			children: nil,
		}
		it.currentLoc.children = append(it.currentLoc.children, &ch)
	}
}

func (it *IndexedTree) ChangeLoc(name string) {
	switch name {
	case "..":
		if it.currentLoc.parent == nil {
			panic("nil parent")
		}
		it.currentLoc = it.currentLoc.parent
	case "/":
		it.currentLoc = it.root
	default:
		for _, ch := range it.currentLoc.children {
			if ch.Name == name {
				it.currentLoc = ch
				return
			}
		}
		panic("no child with that Name found")
	}
}

// method PrettyPrint prints the tree in a human-readable format
func (it *IndexedTree) PrettyPrint() {
	var pp func(int, *Node)
	pp = func(indent int, node *Node) {
		for i := 0; i < indent; i++ {
			fmt.Print("---")
		}
		fmt.Println(node.Name)
		for _, child := range node.children {
			pp(indent+1, child)
		}
	}
	pp(0, it.root)
}

func (it *IndexedTree) AllSmallSizes(limit int) []int {
	_, found := it.allSmallSizes(it.root, limit)
	return found
}

func (it *IndexedTree) allSmallSizes(n *Node, limit int) (int, []int) {
	var size int
	var foundSizes []int
	for _, ch := range n.children {
		if ch.Size == -1 {
			childSize, childFoundSizes := it.allSmallSizes(ch, limit)
			size += childSize
			foundSizes = append(foundSizes, childFoundSizes...)
		} else {
			size += ch.Size
		}
	}
	if size <= limit {
		foundSizes = append(foundSizes, size)
	}
	fmt.Println("out:", n.Name, size, foundSizes)
	return size, foundSizes
}
