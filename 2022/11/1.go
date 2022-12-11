package main

import (
	"fmt"
	"math/big"

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
		fmt.Println(i)
	}
	monkeys.Print()
}

var monkeys = monkey.NewMonkeys(
	[]*monkey.Monkey{
		monkey.NewMonkey(
			[]int{92, 73, 86, 83, 65, 51, 55, 93},
			func(i *big.Int) *big.Int { return i.Mul(i, big.NewInt(5)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(11)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{99, 67, 62, 61, 59, 98},
			func(i *big.Int) *big.Int { return i.Mul(i, i) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(2)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{81, 89, 56, 61, 99},
			func(i *big.Int) *big.Int { return i.Mul(i, big.NewInt(7)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(5)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{97, 74, 68},
			func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(1)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(17)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{78, 73},
			func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(3)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(19)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{50},
			func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(5)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(7)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{95, 88, 53, 75},
			func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(8)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(3)).String() == "0" },
		),
		monkey.NewMonkey(
			[]int{50, 77, 98, 85, 94, 56, 89},
			func(i *big.Int) *big.Int { return i.Add(i, big.NewInt(2)) },
			func(i *big.Int) bool { return new(big.Int).Mod(i, big.NewInt(13)).String() == "0" },
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

//
//var testMonkeys = monkey.NewMonkeys(
//	[]*monkey.Monkey{
//		monkey.NewMonkey(
//			[]int{79, 98},
//			func(i int) int { return i * 19 },
//			func(i int) bool { return i%23 == 0 },
//		),
//		monkey.NewMonkey(
//			[]int{54, 65, 75, 74},
//			func(i int) int { return i + 6 },
//			func(i int) bool { return i%19 == 0 },
//		),
//		monkey.NewMonkey(
//			[]int{79, 60, 97},
//			func(i int) int { return i * i },
//			func(i int) bool { return i%13 == 0 },
//		),
//		monkey.NewMonkey(
//			[]int{74},
//			func(i int) int { return i + 3 },
//			func(i int) bool { return i%17 == 0 },
//		),
//	},
//	[][2]int{
//		{2, 3},
//		{2, 0},
//		{1, 3},
//		{0, 1},
//	},
//)
