package knots

type Knots struct {
	visited      map[int]bool // int = x + y*100000
	headPosition int
	tail         []int
}

func New(length) *Knots {
	return &Knots{
		visited:      make(map[int]bool),
		headPosition: coordinates(50000, 50000),
		tail:         coordinates(50000, 50000),
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

func (k *Knots) Move(direction string, steps int) {
	xHead, yHead := parseCoordinates(k.headPosition)
	xTail, yTail := parseCoordinates(k.tailPosition)
	for i := 0; i < steps; i++ {
		switch direction {
		case up:
			yHead++
			if yHead-yTail > 1 {
				xTail = xHead
				yTail++
			}
		case right:
			xHead++
			if xHead-xTail > 1 {
				xTail++
				yTail = yHead
			}
		case down:
			yHead--
			if yTail-yHead > 1 {
				xTail = xHead
				yTail--
			}
		case left:
			xHead--
			if xTail-xHead > 1 {
				xTail--
				yTail = yHead
			}
		}
		k.visited[coordinates(xTail, yTail)] = true
	}
	k.headPosition = coordinates(xHead, yHead)
	k.tailPosition = coordinates(xTail, yTail)
}

func (k *Knots) CountVisited() int {
	return len(k.visited)
}
