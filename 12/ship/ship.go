package ship

import (
	"math"
	"strconv"
)

type Ship struct {
	movementX, movementY int
	direction            int
}

func (s *Ship) north(i int) {
	s.movementY += i
}

func (s *Ship) south(i int) {
	s.movementY -= i
}

func (s *Ship) east(i int) {
	s.movementX += i
}

func (s *Ship) west(i int) {
	s.movementX -= i
}

func (s *Ship) left(i int) {
	s.direction += i
	s.direction %= 360
}

func (s *Ship) right(i int) {
	s.direction -= i
	s.direction %= 360
	if s.direction < 0 {
		s.direction += 360
	}
}

func (s *Ship) forward(i int) {
	sin, cos := math.Sincos(float64(s.direction) * math.Pi / 180)
	s.movementX += int(float64(i) * cos)
	s.movementY += int(float64(i) * sin)
}

func (s *Ship) Move(str string) {
	val, _ := strconv.Atoi(str[1:])
	switch str[0] {
	case 'N':
		s.north(val)
	case 'S':
		s.south(val)
	case 'E':
		s.east(val)
	case 'W':
		s.west(val)
	case 'R':
		s.right(val)
	case 'L':
		s.left(val)
	case 'F':
		s.forward(val)
	}
}

func (s *Ship) ManhattanDistance() int {
	return int(math.Abs(float64(s.movementX)) + math.Abs(float64(s.movementY)))
}
