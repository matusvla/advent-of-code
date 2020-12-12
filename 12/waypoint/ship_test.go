package waypoint

import (
	"fmt"
	"testing"
)

func TestShip(t *testing.T) {
	s := Ship{
		WaypointX: 10,
		WaypointY: 1,
	}
	inputs := []string{"W10", "F1"}
	for _, input := range inputs {
		s.Move(input)
	}
	fmt.Println(s.ManhattanDistance())
}
