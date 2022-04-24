package main

import "math"

type Bingo [5][5]int

func NewBingo(vals [5][5]int) Bingo {
	return vals
}

type BingoSolver struct {
	drawn               map[int]int
	worstBingo          Bingo
	worstSolutionLength int
	lastDrawnNumber     int
}

func NewBingoSolver(drawnSequence []int) *BingoSolver {
	m := make(map[int]int)
	for i, n := range drawnSequence {
		m[n] = i
	}
	return &BingoSolver{
		drawn:               m,
		worstBingo:          Bingo{},
		worstSolutionLength: 0,
	}
}

func (bs *BingoSolver) Solve(b Bingo) {
	bestSolutionLength := math.MaxInt
	var lastNumber int
	for i := 0; i < 5; i++ {
		var worstIndex, lastNo int
		for j := 0; j < 5; j++ {
			index, ok := bs.drawn[b[i][j]]
			if !ok {
				panic("!ok")
			}
			if index > worstIndex {
				worstIndex = index
				lastNo = b[i][j]
			}
		}
		if bestSolutionLength > worstIndex {
			bestSolutionLength = worstIndex
			lastNumber = lastNo
		}
	}
	for i := 0; i < 5; i++ {
		var worstIndex, lastNo int
		for j := 0; j < 5; j++ {
			index, ok := bs.drawn[b[j][i]]
			if !ok {
				panic("!ok")
			}
			if index > worstIndex {
				worstIndex = index
				lastNo = b[j][i]
			}
		}
		if bestSolutionLength > worstIndex {
			bestSolutionLength = worstIndex
			lastNumber = lastNo
		}
	}
	if bestSolutionLength > bs.worstSolutionLength {
		bs.worstSolutionLength = bestSolutionLength
		bs.worstBingo = b
		bs.lastDrawnNumber = lastNumber
	}
}

func (bs *BingoSolver) Best() int {
	var result int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := bs.worstBingo[i][j]
			index, ok := bs.drawn[val]
			if !ok || bs.worstSolutionLength < index {
				result += val
			}
		}
	}
	return result * bs.lastDrawnNumber
}
