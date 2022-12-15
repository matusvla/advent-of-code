package display

import "fmt"

type Display struct {
	pixels         []bool
	width          int
	height         int
	spritePosition int
	iteration      int
}

func New(width, height int) *Display {
	pixels := make([]bool, width*height)
	return &Display{
		pixels:         pixels,
		width:          width,
		height:         height,
		spritePosition: 1,
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (d *Display) Iterate(newSpritePosition int) {
	if abs(d.spritePosition-d.iteration) < 2 {
		d.pixels[d.iteration] = true
	}
	d.iteration++
	d.spritePosition = newSpritePosition
	d.Print()
	fmt.Println()
}

func (d *Display) Print() {
	for i := 0; i < d.width*d.height; i++ {
		if d.pixels[i] {
			fmt.Print("#")
		} else if i == d.spritePosition {
			fmt.Print("S")
		} else {
			fmt.Print(".")
		}
		if i%d.width == d.width-1 {
			fmt.Println()
		}
	}
}
