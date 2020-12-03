package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var lineRe = regexp.MustCompile(`^(\d+)-(\d+) (.): (.*)$`)

func OldPolicyRegExp(from, to int, key, value string) bool {
	return regexp.MustCompile(
		fmt.Sprintf(`^[^%v]*([%v][^%v]*){%v,%v}$`, key, key, key, from, to),
	).MatchString(value)
}

func OldPolicy(from, to int, key, value string) bool {
	contains := strings.Count(value, key)
	return contains >= from && contains <= to
}

func NewPolicy(pos1, pos2 int, key, value string) bool {
	if value[pos1-1] != value[pos2-1] && (value[pos1-1] == key[0] || value[pos2-1] == key[0]) {
		return true
	}
	return false
}

func ValidateLine(passLine string, policy func(int, int, string, string) bool) (bool, error) {
	submatch := lineRe.FindStringSubmatch(passLine)
	if len(submatch) < 5 {
		return false, errors.New("invalid line")
	}
	from, err := strconv.Atoi(submatch[1])
	if err != nil {
		return false, err
	}
	to, err := strconv.Atoi(submatch[2])
	if err != nil {
		return false, err
	}
	return policy(from, to, submatch[3], submatch[4]), nil
}
