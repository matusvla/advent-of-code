package answers

import (
	"math/bits"
	"strings"
)

func Union(s string) int {
	usedLetters := make(map[int32]struct{})
	answers := strings.Split(s, "\n")
	for _, answer := range answers {
		for _, subanswer := range answer {
			usedLetters[subanswer] = struct{}{}
		}
	}
	return len(usedLetters)
}

func Intersection(s string) int {
	var usedLetters uint = 1<<27 - 1
	answers := strings.Split(s, "\n")
	for _, answer := range answers {
		if answer == "" {
			continue
		}
		var numAnswer uint = 0
		for _, subanswer := range answer {
			numAnswer += 1 << (subanswer - 'a')
		}
		usedLetters &= numAnswer
	}
	return bits.OnesCount(usedLetters)
}
