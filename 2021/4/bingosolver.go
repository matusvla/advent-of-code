package main

//
//import "math"
//
//type Bingo [5][5]int
//
//func NewBingo(vals [5][5]int) Bingo {
//	return vals
//}
//
//type BingoSolver struct {
//	drawn              map[int]int
//	bestBingo          Bingo
//	bestSolutionLength int
//	lastDrawnNumber    int
//}
//
//func NewBingoSolver(drawnSequence []int) *BingoSolver {
//	m := make(map[int]int)
//	for i, n := range drawnSequence {
//		m[n] = i
//	}
//	return &BingoSolver{
//		drawn:              m,
//		bestBingo:          Bingo{},
//		bestSolutionLength: math.MaxInt,
//	}
//}
//
//func (bs *BingoSolver) Solve(b Bingo) {
//Outer:
//	for i := 0; i < 5; i++ {
//		var worstIndex, worstNumber int
//		for j := 0; j < 5; j++ {
//			index, ok := bs.drawn[b[i][j]]
//			if !ok {
//				continue Outer
//			}
//			if index > bs.bestSolutionLength {
//				continue Outer
//			}
//			if index > worstIndex {
//				worstIndex = index
//				worstNumber = b[i][j]
//			}
//		}
//		bs.bestSolutionLength = worstIndex
//		bs.bestBingo = b
//		bs.lastDrawnNumber = worstNumber
//	}
//Outer2:
//	for i := 0; i < 5; i++ {
//		var worstIndex, worstNumber int
//		for j := 0; j < 5; j++ {
//			index, ok := bs.drawn[b[j][i]]
//			if !ok {
//				continue Outer2
//			}
//			if index > bs.bestSolutionLength {
//				continue Outer2
//			}
//			if index > worstIndex {
//				worstIndex = index
//				worstNumber = b[j][i]
//			}
//		}
//		bs.bestSolutionLength = worstIndex
//		bs.bestBingo = b
//		bs.lastDrawnNumber = worstNumber
//	}
//}
//
//func (bs *BingoSolver) Best() int {
//	var result int
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 5; j++ {
//			val := bs.bestBingo[i][j]
//			index, ok := bs.drawn[val]
//			if !ok || bs.bestSolutionLength < index {
//				result += val
//			}
//		}
//	}
//	return result * bs.lastDrawnNumber
//}
