package main

import (
	"fmt"
	"math"
	"sort"
)

// target area: x=195..238, y=-93..-67
//const (
//	xStart = 195
//	xEnd   = 238
//	yStart = 93
//	yEnd   = -67
//)

const (
	xStart = 5
	xEnd   = 6
	yStart = 1
	yEnd   = -1
)

//const (
//	xStart = 20
//	xEnd   = 30
//	yStart = -5
//	yEnd   = -10
//)

//func main() {
//	var maxYVelocity, maxXVelocity int
//	var allPos int
//	for i := -yEnd; i > -yStart; i-- {
//		var positive bool
//		yVelocity := i
//		if yVelocity < 0 {
//			positive = true
//			yVelocity = -yVelocity
//		}
//		for x := xStart; x < xEnd+1; x++ {
//			n, ok := isXPossible(x, 2*yVelocity-1)
//			if ok {
//				fmt.Println(i,x)
//				allPos++
//				if n > maxYVelocity {
//					if positive {
//						yVelocity = n
//					} else {
//						yVelocity = n - 1
//					}
//				}
//				if yVelocity > maxYVelocity {
//					maxYVelocity = yVelocity
//				}
//			}
//		}
//	}
//
//	fmt.Println(maxXVelocity)
//	fmt.Println(maxYVelocity)
//
//	fmt.Println(((maxYVelocity + 1) * (maxYVelocity)) / 2)
//	fmt.Println(allPos)
//
//	//printMap(xStart, xEnd, yStart,yEnd, maxXVelocity, maxYVelocity-1)
//}

type pair struct {
	x, y int
}

func main() {
	allPos := make(map[pair]struct{})
	maxYVelocity := int(math.Max(math.Abs(yStart), math.Abs(yEnd-1)))
	for i := maxYVelocity; i > -maxYVelocity; i-- {
		yVelocity := i
		moves := 0
		var posY int
		for {
			moves++
			posY += yVelocity
			if yEnd <= posY && yStart >= posY {
				fmt.Println(i)
				for x := xStart; x < xEnd+1; x++ {
					possibleX := listPossible(x, moves, i > 0)
					for _, xx := range possibleX {
						allPos[pair{xx, i}] = struct{}{}
					}
				}
			}
			if posY < yEnd {
				break
			}
			yVelocity--
		}
	}

	var allPosSort []pair
	for ap := range allPos {
		allPosSort = append(allPosSort, ap)
	}
	sort.SliceStable(allPosSort, func(i, j int) bool {
		return allPosSort[i].y > allPosSort[j].y || (allPosSort[i].y == allPosSort[j].y && allPosSort[i].x > allPosSort[j].x)
	})
	for _, ap := range allPosSort {
		fmt.Println(ap.x, ap.y)
	}

	fmt.Println(len(allPos))
}

func listPossible(x, moves int, positive bool) []int {
	var result []int
	for n := 1; n <= x; n++ { // can be shorter
		for m := 1; m < n+1; m++ {
			if 2*x == -(m*m)+m+(n*n)+n {
				if n-m+1 == moves || (positive && m == 1) {
					result = append(result, n)
				}
			}
		}
	}
	return result
}

func isXPossible(x, moves int) (int, bool) {
	for n := 1; n < x; n++ { // can be shorter
		for m := 1; m < n+1; m++ {
			if 2*x == -(m*m)+m+(n*n)+n {
				if n-m+1 == moves || m == 1 {
					return n, true
				}
			}
		}
	}
	return 0, false
}
