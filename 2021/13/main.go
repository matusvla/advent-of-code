package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pp := NewPointPaper()
	// scan points
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		tParts := strings.Split(t, ",")
		x, err := strconv.Atoi(tParts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(tParts[1])
		if err != nil {
			panic(err)
		}
		pp[Point{
			x: x,
			y: y,
		}] = struct{}{}
	}

	//inversePP := NewPointPaper()
	//
	//for p := range pp {
	//	inversePP[Point{1310-p.x, p.y}] = struct{}{}
	//}
	//
	//result := 1500
	//for p := range pp {
	//	if _, ok :=  inversePP[Point{p.x, p.y}]; ok {
	//		result--
	//	}
	//}
	//fmt.Println(result / 2)

	// scan folds
	mapFunc := func(point Point) Point {
		return point
	}
	for scanner.Scan() {
		t := scanner.Text()
		tParts := strings.Split(t, "=")
		foldLineIndex, err := strconv.Atoi(tParts[1])
		switch tParts[0] {
		case "fold along x":
			mf := mapFunc
			mapFunc = func(point Point) Point {
				p := mf(point)
				x := p.x
				if x > foldLineIndex-1 {
					x = 2*foldLineIndex - x
				}
				return Point{x, p.y}
			}

		case "fold along y":
			mf := mapFunc
			mapFunc = func(point Point) Point {
				p := mf(point)
				y := p.y
				if y > foldLineIndex-1 {
					y = 2*foldLineIndex - y
				}
				return Point{p.x, y}
			}
		default:
			panic("no fold line")
		}
		if err != nil {
			panic(err)
		}
	}

	newPP := NewPointPaper()
	for p := range pp {
		newP := mapFunc(p)
		newPP[newP] = struct{}{}
	}
	//pp.Print()
	newPP.Print()
	fmt.Println(newPP)
}

type Point struct {
	x, y int
}

type PointPaper map[Point]struct{}

func (p PointPaper) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 80; j++ {
			if _, ok := p[Point{j, i}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func NewPointPaper() PointPaper {
	return make(map[Point]struct{})
}
