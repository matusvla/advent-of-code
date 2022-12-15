package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const impactedLineY = 2000000

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	usedPositions := make(map[int]struct{})
	for scanner.Scan() {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			log.Fatal(err)
		}
		lb := sx - abs(sx-bx) + abs(sy-impactedLineY) - abs(sy-by)
		ub := sx + abs(sx-bx) - abs(sy-impactedLineY) + abs(sy-by)
		for i := lb; i < ub; i++ {
			usedPositions[i] = struct{}{}
		}
	}
	fmt.Println(len(usedPositions))
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var igs []IntervalGen
	for scanner.Scan() {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			log.Fatal(err)
		}
		ig := NewIntervalGen(sx, sy, bx, by)
		igs = append(igs, ig)
	}

	const sizeMax = 4_000_000
	for i := 0; i < sizeMax+1; i++ {
		intervals := make([]interval, 0, len(igs))
		for _, ig := range igs {
			inter := ig.GetInterval(i)
			if inter != nil {
				intervals = append(intervals, *inter)
			}
		}
		if res := isCover(intervals, 0, sizeMax); res != -1 {
			fmt.Printf("x: %d, y: %d, result: %d", res, i, res*4_000_000+i)
			break
		}
	}
}

type IntervalGen struct {
	centerX        int
	centerY        int
	halfDiagLength int
}

func NewIntervalGen(sx, sy, bx, by int) IntervalGen {
	return IntervalGen{
		centerX:        sx,
		centerY:        sy,
		halfDiagLength: abs(sx-bx) + abs(sy-by),
	}
}

type interval [2]int

func (ig IntervalGen) GetInterval(y int) *interval {
	start := ig.centerX - ig.halfDiagLength + abs(ig.centerY-y)
	end := ig.centerX + ig.halfDiagLength - abs(ig.centerY-y)
	if start > end {
		return nil
	}
	return &interval{start, end}
}

func isCover(intervals []interval, lb, ub int) int {
	// sort
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current[1] > ub {
			break
		}
		if current[1] < lb {
			current = intervals[i]
			continue
		}
		if current[1] > intervals[i][1] {
			continue
		}
		if current[1] < intervals[i][0] {
			return current[1] + 1
		}
		current = intervals[i]
	}
	return -1
}
