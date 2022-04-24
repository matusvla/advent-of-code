package seatfiller

import "fmt"

type point struct {
	used       bool
	neighbours int
}

type floorplan struct {
	seats      []point
	colLength  int
	nextToFill []int
}

func (f *floorplan) String() string {
	var result string
	for i, s := range f.seats {
		if i%f.colLength == 0 {
			result += fmt.Sprintln()
		}
		if s.used {
			result += fmt.Sprintf("X")
		} else {
			result += fmt.Sprintf("O")
		}
	}
	return result
}

func (f *floorplan) addNeighbour(x, y int) {
	index := x + f.colLength*y
	if x < 0 || x > f.colLength-1 || y < 0 || f.colLength*y > len(f.seats)-1 {
		return
	}
	f.seats[index].neighbours++
}

func (f *floorplan) setUsed(x, y int) {
	index := x + f.colLength*y
	if x < 0 || x > f.colLength-1 || y < 0 || f.colLength*y > len(f.seats)-1 {
		return
	}
	f.seats[index].used = true
}

func newFloorplan(b [][]byte) *floorplan {
	xSize, ySize := len(b), len(b[0])
	f := floorplan{
		seats:     make([]point, xSize*ySize),
		colLength: ySize,
	}

	for i, row := range b {
		for j, val := range row {
			if val == '.' {
				f.setUsed(j, i)
			}
			if val == 'L' {
				f.addNeighbour(j-1, i-1)
				f.addNeighbour(j-1, i)
				f.addNeighbour(j-1, i+1)
				f.addNeighbour(j, i-1)
				f.addNeighbour(j, i+1)
				f.addNeighbour(j+1, i-1)
				f.addNeighbour(j+1, i)
				f.addNeighbour(j+1, i+1)
			}
		}
	}

	for i, val := range f.seats {
		if val.neighbours < 4 {
			f.nextToFill = append(f.nextToFill, i)
		}
	}

	return &f
}

func (f *floorplan) getCoordinates(x, y int) (int, bool) {
	if x < 0 || x > f.colLength-1 || y < 0 || f.colLength*y > len(f.seats)-1 {
		return 0, false
	}
	return x + f.colLength*y, true
}

func getNeighbourCoordinates(x, y int) [8][2]int {
	return [8][2]int{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}
}

func (f *floorplan) findNeighboursToFill(coordinates int) {
	fmt.Println(f.String())
	i, j := coordinates/f.colLength, coordinates%f.colLength

	nc := getNeighbourCoordinates(j, i)
	for _, c := range nc {
		f.setUsed(c[0], c[1])
	}

	for _, c := range nc {
		nnc := getNeighbourCoordinates(c[0], c[1])
		for _, cc := range nnc {
			coord, ok := f.getCoordinates(cc[0], cc[1])
			if !ok {
				continue
			}
			if f.seats[coord].used {
				continue
			}
			f.seats[coord].used = true
			f.nextToFill = append(f.nextToFill, coord)
		}
	}
}

func (f *floorplan) fillOneIteration() int {
	var filled int
	var newToFill []int
	prevLen := len(f.nextToFill)
	for _, coordinates := range f.nextToFill {
		filled++
		f.findNeighboursToFill(coordinates)
	}
	f.nextToFill = append(f.nextToFill, newToFill...)[prevLen:]
	return filled
}

func ProcessFloor(b [][]byte) int {
	f := newFloorplan(b)
	var totFilled int
	fmt.Println(f.String())
	for len(f.nextToFill) > 0 {
		totFilled += f.fillOneIteration()
		fmt.Println(f.String())
	}
	return totFilled
}
