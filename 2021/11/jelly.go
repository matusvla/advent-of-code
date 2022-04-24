package main

import "fmt"

type Jellyfish struct {
	value      int
	posX, posY int
	flashed    bool
	neighbours []*Jellyfish
}

func (jf *Jellyfish) Flash() int {
	if jf.flashed || jf.value < 10 {
		return 0
	}
	jf.flashed = true
	result := 1
	for _, neighbour := range jf.neighbours {
		neighbour.value++
		if neighbour.value >= 10 {
			result += neighbour.Flash()
		}
	}
	return result
}

type AllJellyFish struct {
	jf     []*Jellyfish
	rowLen int
}

func (jfs *AllJellyFish) getIndex(i, j int) int {
	return jfs.rowLen*j + i
}

func (jfs *AllJellyFish) PushBack(jf *Jellyfish) {
	i, j := jf.posX, jf.posY
	jfs.jf = append(jfs.jf, jf)
	if i > 0 {
		neighbour := jfs.jf[jfs.getIndex(i-1, j)]
		jf.neighbours = append(jf.neighbours, neighbour)
		neighbour.neighbours = append(neighbour.neighbours, jf)
	}
	if j > 0 {
		neighbour := jfs.jf[jfs.getIndex(i, j-1)]
		jf.neighbours = append(jf.neighbours, neighbour)
		neighbour.neighbours = append(neighbour.neighbours, jf)
	}
	if i > 0 && j > 0 {
		neighbour := jfs.jf[jfs.getIndex(i-1, j-1)]
		jf.neighbours = append(jf.neighbours, neighbour)
		neighbour.neighbours = append(neighbour.neighbours, jf)
	}
	if i < jfs.rowLen-1 && j > 0 {
		neighbour := jfs.jf[jfs.getIndex(i+1, j-1)]
		jf.neighbours = append(jf.neighbours, neighbour)
		neighbour.neighbours = append(neighbour.neighbours, jf)
	}
}

func (jfs *AllJellyFish) NextRound() int {
	for _, jf := range jfs.jf {
		jf.flashed = false
		jf.value++
	}
	var result int
	for _, jf := range jfs.jf {
		result += jf.Flash()
	}
	for _, jf := range jfs.jf {
		if jf.flashed {
			jf.value = 0
		}
	}
	return result
}

func (jfs *AllJellyFish) Print() {
	j := 0
	for {
		for i := 0; i < jfs.rowLen; i++ {
			index := jfs.getIndex(i, j)
			if index > len(jfs.jf)-1 {
				fmt.Println()
				fmt.Println()
				return
			}
			fmt.Print(jfs.jf[index].value)
		}
		fmt.Println()
		j++
	}
}
