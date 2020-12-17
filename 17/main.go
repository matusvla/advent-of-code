package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"advent-of-code-2020/17/convolution"
)

func main() {
	f, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	//input := make([][][]int, 1)
	input := make([][][][]int, 1)
	input[0] = make([][][]int, 1)
	for scanner.Scan() {
		t := scanner.Text()
		t = strings.ReplaceAll(t, ".", "0")
		t = strings.ReplaceAll(t, "#", "1")
		var row []int
		fmt.Println(t)
		for _, ch := range t {
			row = append(row, int(ch-'0'))
		}
		input[0][0] = append(input[0][0], row)
	}

	//c := convolution.New(
	//	[3][3][3]int{
	//		{
	//			{1, 1, 1},
	//			{1, 1, 1},
	//			{1, 1, 1},
	//		},
	//		{
	//			{1, 1, 1},
	//			{1, 0, 1},
	//			{1, 1, 1},
	//		},
	//		{
	//			{1, 1, 1},
	//			{1, 1, 1},
	//			{1, 1, 1},
	//		},
	//	},
	//	map0,
	//	map1,
	//)
	c := convolution.New4D(
		[3][3][3][3]int{
			{
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			{
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			{
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
		},
		map0,
		map1,
	)

	for i := 0; i < 6; i++ {
		input = c.ConvolutionExtendBoundsWithMod(input)
	}
	var result int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			for k := 0; k < len(input[i][j]); k++ {
				for l := 0; l < len(input[i][j][k]); l++ {
					if input[i][j][k][l] > 0 {
						result++
					}
				}
			}
		}
	}
	fmt.Println(result)
}

func map0(i int) int {
	if i == 3 {
		return 1
	}
	return 0
}

func map1(i int) int {
	if i == 2 || i == 3 {
		return 1
	}
	return 0
}
