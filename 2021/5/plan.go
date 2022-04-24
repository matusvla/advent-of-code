package main

import "fmt"

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (ln Line) Print() {
	fmt.Printf("%d,%d -> %d,%d\n", ln.start.x, ln.start.y, ln.end.x, ln.end.y)
}

func (ln Line) Points(onlyStraight bool) []Point {
	if onlyStraight && ln.start.x != ln.end.x && ln.start.y != ln.end.y {
		return nil
	}
	var result []Point
	// todo this is not general, works only for straight lines
	xStart, xStop := ln.start.x, ln.end.x
	var xRotated bool
	if xStart > xStop {
		xRotated = true
		xStart, xStop = xStop, xStart
	}
	yStart, yStop := ln.start.y, ln.end.y
	var yRotated bool
	if yStart > yStop {
		yStart, yStop = yStop, yStart
		yRotated = true
	}
	if ln.start.x != ln.end.x && ln.start.y != ln.end.y { // diagonal lines
		for i := 0; i < xStop-xStart+1; i++ {
			x := xStart + i
			y := yStart + i
			if xRotated != yRotated {
				y = yStop - i
			}
			result = append(result, Point{
				x: x,
				y: y,
			})
		}
		return result
	}
	for i := xStart; i < xStop+1; i++ {
		for j := yStart; j < yStop+1; j++ {
			result = append(result, Point{
				x: i,
				y: j,
			})
		}
	}
	return result
}

type Plan struct {
	data [1000][1000]int
} // todo dynamic allocation and growth

func (p *Plan) Print() {
	var transposedResult [10][10]int
	for i := 0; i < len(p.data); i++ {
		for j := 0; j < len(p.data[i]); j++ {
			transposedResult[j][i] = p.data[i][j]
		}
	}
	for i := 0; i < len(transposedResult); i++ {
		for j := 0; j < len(transposedResult[i]); j++ {
			if transposedResult[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(transposedResult[i][j])
			}
		}
		fmt.Println()
	}
}

func NewPlan() *Plan {
	return &Plan{}
}

func (p *Plan) Add(line Line) {
	lp := line.Points(false)
	//line.Print()
	for _, point := range lp {
		p.data[point.x][point.y] += 1
	}
	//p.Print()
}

func (p *Plan) Count(threshold int) int {
	var result int
	for i := 0; i < len(p.data); i++ {
		for j := 0; j < len(p.data[i]); j++ {
			if p.data[i][j] >= threshold {
				result++
			}
		}
	}
	return result
}
