package ship

import (
	"fmt"
	"testing"
)

func TestShip(t *testing.T) {
	s := Ship{}
	inputs := []string{"F10", "N3", "F7", "R90", "F11"}
	for _, input := range inputs {
		s.Move(input)
	}
	fmt.Println(s.ManhattanDistance())
}
