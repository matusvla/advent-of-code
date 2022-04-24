package seatfiller_wrong

import (
	"math"
)

func FillSeats(inb [][]byte) ([][]byte, int) {
	var outb [][]byte
	for i := 0; i < len(inb); i++ {
		var b []byte
		for j := 0; j < len(inb[0]); j++ {
			b = append(b, inb[i][j])
		}
		outb = append(outb, b)
	}
	var changed int
	for i := 0; i < len(inb); i++ {
		for j := 0; j < len(inb[0]); j++ {
			if inb[i][j] == '.' {
				continue
			}
			filled := 0
			for fi := i - 1; fi <= i+1; fi++ {
				if fi < 0 || fi >= len(inb) {
					continue
				}
				for fj := j - 1; fj <= j+1; fj++ {
					if fj < 0 || fj >= len(inb[0]) || (fi == i && fj == j) {
						continue
					}
					if inb[fi][fj] == '#' {
						filled++
					}
				}
			}
			if filled == 0 && inb[i][j] == 'L' {
				outb[i][j] = '#'
				changed++
			}
			if filled >= 4 && inb[i][j] == '#' {
				outb[i][j] = 'L'
				changed++
			}
		}
	}
	return outb, changed
}

func FillSeats2(inb [][]byte) ([][]byte, int) {
	neighbours := make([][]int, len(inb)+2)
	neighbours[0] = make([]int, len(inb[0])+2)
	neighbours[1] = make([]int, len(inb[0])+2)
	for i := 0; i < len(inb); i++ {
		neighbours[i+2] = make([]int, len(inb[0])+2)
		for j := 0; j < len(inb[0]); j++ {
			if inb[i][j] == '#' {
				neighbours[i][j]++
				neighbours[i][j+1]++
				neighbours[i][j+2]++
				neighbours[i+1][j]++
				neighbours[i+1][j+2]++
				neighbours[i+2][j]++
				neighbours[i+2][j+1]++
				neighbours[i+2][j+2]++
			}
		}
	}

	var changes int
	for i := 1; i < len(neighbours)-1; i++ {
		for j := 1; j < len(neighbours[i])-1; j++ {
			if neighbours[i][j] >= 4 && inb[i-1][j-1] == '#' {
				inb[i-1][j-1] = 'L'
				changes++
			}
			if neighbours[i][j] == 0 && inb[i-1][j-1] == 'L' {
				inb[i-1][j-1] = '#'
				changes++
			}
		}
	}
	return inb, changes
}

func FillSeats3(inb [][]byte) ([][]byte, int) {
	filledNeighbours := make([][]int, len(inb))
	for i := range inb {
		filledNeighbours[i] = make([]int, len(inb[0]))
	}

	for i := 0; i < len(inb); i++ {
		for j := 0; j < len(inb[0]); j++ {
			if inb[i][j] == '.' {
				filledNeighbours[i][j] = -1
				continue
			}
			for fi := int(math.Max(0, float64(i-1))); fi <= int(math.Min(float64(i+1), float64(len(inb)-1))); fi++ {
				for fj := int(math.Max(0, float64(j-1))); fj <= int(math.Min(float64(j+1), float64(len(inb[0])-1))); fj++ {
					if fi == i && fj == j {
						continue
					}
					if inb[fi][fj] == '#' {
						filledNeighbours[i][j]++
					}
				}
			}
		}
	}

	var changes int
	for i := 0; i < len(filledNeighbours); i++ {
		for j := 0; j < len(filledNeighbours[0]); j++ {
			if filledNeighbours[i][j] >= 4 && inb[i][j] == '#' {
				inb[i][j] = 'L'
				changes++
			}
			if filledNeighbours[i][j] == 0 && inb[i][j] == 'L' {
				inb[i][j] = '#'
				changes++
			}
		}
	}
	return inb, changes
}
