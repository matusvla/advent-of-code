package rope

import "fmt"

type Rope struct {
	visited      map[int]bool // int = x + y*100000
	headPosition int
	tail         []int
}

const startPos = 20000

func New(length int) *Rope {
	var tail []int
	for i := 0; i < length; i++ {
		tail = append(tail, coordinates(startPos, startPos))
	}
	return &Rope{
		visited:      make(map[int]bool),
		headPosition: coordinates(startPos, startPos),
		tail:         tail,
	}
}

func coordinates(x, y int) int {
	return x + y*100000
}

func parseCoordinates(n int) (x, y int) {
	return n % 100000, n / 100000
}

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

func (k *Rope) Move(direction string, steps int) {
	for i := 0; i < steps; i++ {
		xHead, yHead := parseCoordinates(k.headPosition)
		switch direction {
		case up:
			yHead++
		case right:
			xHead++
		case down:
			yHead--
		case left:
			xHead--
		}
		k.headPosition = coordinates(xHead, yHead)
		movedPrevPos := k.headPosition
		for j := 0; j < len(k.tail); j++ {
			tailPosition := k.tail[j]
			tailPosition = followPreviousKnot(movedPrevPos, tailPosition)
			k.tail[j] = tailPosition
			if j == len(k.tail)-1 {
				k.visited[tailPosition] = true
			}
			movedPrevPos = tailPosition
			//k.Print()
			//fmt.Printf("Moved %s %d/%d\n", direction, i+1, steps)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func followPreviousKnot(headLocation, tailLocation int) int {
	xHead, yHead := parseCoordinates(headLocation)
	xTail, yTail := parseCoordinates(tailLocation)
	if abs(xHead-xTail) > 1 && abs(yHead-yTail) > 1 {
		if xHead > xTail {
			xTail++
		} else {
			xTail--
		}
		if yHead > yTail {
			yTail++
		} else {
			yTail--
		}
	} else if abs(xHead-xTail) > 1 {
		if xHead > xTail {
			xTail++
			yTail = yHead
		} else {
			xTail--
			yTail = yHead
		}
	} else if abs(yHead-yTail) > 1 {
		if yHead > yTail {
			yTail++
			xTail = xHead
		} else {
			yTail--
			xTail = xHead
		}
	}
	return coordinates(xTail, yTail)
}

func (k *Rope) CountVisited() int {
	return len(k.visited)
}

func (k *Rope) Print() {
	xHead, yHead := parseCoordinates(k.headPosition)
	for i := startPos*2 - 1; i >= 0; i-- {
	Outer:
		for j := 0; j < startPos*2; j++ {
			if i == yHead && j == xHead {
				fmt.Print("H")
				continue
			}
			for tp, tailPosition := range k.tail {
				xTail, yTail := parseCoordinates(tailPosition)
				if i == yTail && j == xTail {
					fmt.Print(tp + 1)
					continue Outer
				}
			}
			if k.visited[coordinates(j, i)] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}
