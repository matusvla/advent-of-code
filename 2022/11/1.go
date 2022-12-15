package main

import (
	"fmt"

	"advent-of-code-2020/2022/11/monkey"
)

// part1
//func main() {
//
//	for i := 0; i < 20; i++ {
//		monkeys.ThrowAll()
//	}
//	monkeys.Print()
//}

// part2
func main() {
	for i := 0; i < 10000; i++ {
		monkeys.ThrowAll()
	}
	fmt.Println(monkeys.MonkeyBusiness())
}

var monkeys = monkey.NewMonkeys(
	[]*monkey.Monkey{
		monkey.NewMonkey(
			[]int{92, 73, 86, 83, 65, 51, 55, 93},
			func(i int) int { return i * 5 },
			11,
		),
		monkey.NewMonkey(
			[]int{99, 67, 62, 61, 59, 98},
			func(i int) int { return i * i },
			2,
		),
		monkey.NewMonkey(
			[]int{81, 89, 56, 61, 99},
			func(i int) int { return i * 7 },
			5,
		),
		monkey.NewMonkey(
			[]int{97, 74, 68},
			func(i int) int { return i + 1 },
			17,
		),
		monkey.NewMonkey(
			[]int{78, 73},
			func(i int) int { return i + 3 },
			19,
		),
		monkey.NewMonkey(
			[]int{50},
			func(i int) int { return i + 5 },
			7,
		),
		monkey.NewMonkey(
			[]int{95, 88, 53, 75},
			func(i int) int { return i + 8 },
			3,
		),
		monkey.NewMonkey(
			[]int{50, 77, 98, 85, 94, 56, 89},
			func(i int) int { return i + 2 },
			13,
		),
	},
	[][2]int{
		{3, 4},
		{6, 7},
		{1, 5},
		{2, 5},
		{2, 3},
		{1, 6},
		{0, 7},
		{4, 0},
	},
)

var testMonkeys = monkey.NewMonkeys(
	[]*monkey.Monkey{
		monkey.NewMonkey(
			[]int{79, 98},
			func(i int) int { return i * 19 },
			23,
		),
		monkey.NewMonkey(
			[]int{54, 65, 75, 74},
			func(i int) int { return i + 6 },
			19,
		),
		monkey.NewMonkey(
			[]int{79, 60, 97},
			func(i int) int { return i * i },
			13,
		),
		monkey.NewMonkey(
			[]int{74},
			func(i int) int { return i + 3 },
			17,
		),
	},
	[][2]int{
		{2, 3},
		{2, 0},
		{1, 3},
		{0, 1},
	},
)
