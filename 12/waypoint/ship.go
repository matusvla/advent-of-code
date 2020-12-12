package waypoint

import (
	"math"
	"strconv"
)

type Ship struct {
	WaypointX, WaypointY int
	positionX, positionY int
}

func (s *Ship) north(i int) {
	s.WaypointY += i
}

func (s *Ship) south(i int) {
	s.WaypointY -= i
}

func (s *Ship) east(i int) {
	s.WaypointX += i
}

func (s *Ship) west(i int) {
	s.WaypointX -= i
}

func (s *Ship) left(i int) {
	s.right(360 - i)
}

func (s *Ship) right(i int) {
	//works for smaller example
	//a1 := float64(s.WaypointY)
	//b1 := float64(s.WaypointX)
	//alpha := math.Atan(a1/b1)
	//beta := - float64(i)* math.Pi / 180
	//c := math.Sqrt(a1*a1 + b1*b1)
	//a2 := math.Sin(alpha + beta) *c
	//b2 := math.Cos(alpha + beta) *c
	//s.WaypointY = int(a2)
	//s.WaypointX = int(b2)

	switch i {
	case 90:
		s.WaypointX, s.WaypointY = s.WaypointY, -s.WaypointX
	case 180:
		s.WaypointX, s.WaypointY = -s.WaypointX, -s.WaypointY
	case 270:
		s.WaypointX, s.WaypointY = -s.WaypointY, s.WaypointX
	case 0, 360:
	default:
		panic("unsupported angle " + strconv.Itoa(i))
	}
}

func (s *Ship) forward(i int) {
	s.positionX += i * s.WaypointX
	s.positionY += i * s.WaypointY
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
	return int(math.Abs(float64(s.positionX)) + math.Abs(float64(s.positionY)))
}
