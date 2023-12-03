package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Draws []Draw
}

type Draw struct {
	Red   int
	Blue  int
	Green int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []*Game
	j := 1
	for scanner.Scan() {
		turnsStr := strings.Split(scanner.Text(), ":")[1]
		turns := strings.Split(turnsStr, ";")
		g := Game{
			Id: j,
		}
		for _, turn := range turns {
			var red, blue, green int
			for _, draw := range strings.Split(turn, ",") {
				draw = strings.TrimSpace(draw)
				switch {
				case strings.HasSuffix(draw, "red"):
					i, err := strconv.Atoi(strings.TrimSuffix(draw, " red"))
					if err != nil {
						panic(err)
					}
					red += i
				case strings.HasSuffix(draw, "blue"):
					i, err := strconv.Atoi(strings.TrimSuffix(draw, " blue"))
					if err != nil {
						panic(err)
					}
					blue += i
				case strings.HasSuffix(draw, "green"):
					i, err := strconv.Atoi(strings.TrimSuffix(draw, " green"))
					if err != nil {
						panic(err)
					}
					green += i
				}
			}
			g.Draws = append(g.Draws, Draw{
				Red:   red,
				Blue:  blue,
				Green: green,
			})
		}
		games = append(games, &g)
		j++
	}

	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	var result int
	for _, g := range games {
		isPossible := true
		for _, d := range g.Draws {
			if d.Red > redLimit || d.Blue > blueLimit || d.Green > greenLimit {
				isPossible = false
				break
			}
		}
		if isPossible {
			result += g.Id
		}
	}

	fmt.Println("part1: ", result)

	var result2 int
	for _, g := range games {
		var maxRed, maxBlue, maxGreen int
		for _, d := range g.Draws {
			if d.Red > maxRed {
				maxRed = d.Red
			}
			if d.Blue > maxBlue {
				maxBlue = d.Blue
			}
			if d.Green > maxGreen {
				maxGreen = d.Green
			}
		}
		result2 += maxRed * maxBlue * maxGreen
	}
	fmt.Println("part2: ", result2)
}
