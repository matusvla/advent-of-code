package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scan points
	var input AllPoints
	var i int
	for scanner.Scan() {
		var p []*Point
		for j, r := range scanner.Text() {
			val := r - '0'
			p = append(p, &Point{
				coordinates: Coordinates{i, j},
				val:         int(val),
				bestPathLen: math.MaxInt,
			})
		}
		input = append(input, p)
		i++
	}

	input = loadFullMap(input)

	// connect points
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			var neighbours []*Point
			if i > 0 {
				neighbours = append(neighbours, input[i-1][j])
			}
			if j > 0 {
				neighbours = append(neighbours, input[i][j-1])
			}
			if i < len(input)-1 {
				neighbours = append(neighbours, input[i+1][j])
			}
			if j < len(input[i])-1 {
				neighbours = append(neighbours, input[i][j+1])
			}
			input[i][j].neighbours = neighbours
		}
	}

	pl, _ := input.FindShortestPath()
	fmt.Println(pl)
	//input.PrintPath(path)
}

func loadFullMap(input AllPoints) AllPoints {
	result := make([][]*Point, 5*len(input))
	for k := 0; k < 5*len(input); k++ {
		result[k] = make([]*Point, 5*len(input[k%len(input)]))
		offset := k / len(input)
		for l := 0; l < 5*len(input[k%len(input)]); l++ {
			orig := input[k%len(input)][l%len(input[k%len(input)])]
			val := ((orig.val+l/len(input[k%len(input)])+offset)-1)%9 + 1
			result[k][l] = &Point{
				coordinates: Coordinates{
					x: k,
					y: l,
				},
				val:         val,
				bestPathLen: orig.bestPathLen,
			}
		}
	}
	return result
}

type AllPoints [][]*Point

func (ap AllPoints) Print() {
	for i := 0; i < len(ap); i++ {
		for j := 0; j < len(ap[i]); j++ {
			fmt.Printf("%d(%d)", ap[i][j].val, len(ap[i][j].neighbours))
		}
		fmt.Println()
	}
}

func (ap AllPoints) PrintPath(path []Coordinates) {
	pm := make(map[Coordinates]struct{})
	for _, p := range path {
		pm[Coordinates{p.x, p.y}] = struct{}{}
	}
	var result int
	for i := 0; i < len(ap); i++ {
		for j := 0; j < len(ap[i]); j++ {
			item := ap[i][j]
			_, ok := pm[Coordinates{item.coordinates.x, item.coordinates.y}]
			if ok {
				fmt.Printf("\u001b[31m")
				result += item.val
			}
			fmt.Printf("%d", item.val)
			if ok {
				fmt.Printf("\u001b[0m")
			}
		}
		fmt.Println()
	}
}

func (ap AllPoints) FindShortestPath() (int, []Coordinates) {
	stack := list.New()
	endPoint := ap[len(ap)-1][len(ap[0])-1]
	stack.PushBack(endPoint)
	endPoint.bestPathLen = endPoint.val
	endPoint.bestPath = []Coordinates{}
	for stack.Len() > 0 {
		elem := stack.Front()
		stack.Remove(elem)
		p := elem.Value.(*Point)
		pathLen := p.bestPathLen
		for _, n := range ap[p.coordinates.x][p.coordinates.y].neighbours {
			if n.bestPathLen <= pathLen+n.val {
				continue
			}
			var path []Coordinates
			for _, p := range p.bestPath {
				path = append(path, p)
			}
			path = append(path, n.coordinates)
			n.bestPath = path
			n.bestPathLen = pathLen + n.val
			stack.PushBack(n)
		}
	}
	return ap[0][0].bestPathLen - ap[0][0].val, ap[0][0].bestPath
}

type Coordinates struct {
	x, y int
}

type Point struct {
	coordinates Coordinates
	val         int
	bestPathLen int
	bestPath    []Coordinates
	neighbours  []*Point
}
