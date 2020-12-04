package passport

import (
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func Process(s string) bool {
	result := true
	for _, field := range requiredFields {
		result = result && strings.Contains(s, field)
	}
	return result
}

func checkRange(s string, lb, ub int) bool {
	i, _ := strconv.Atoi(s)
	return i >= lb && i <= ub
}

var hclRe = regexp.MustCompile(`^#[a-f0-9]{6}$`)
var pidRe = regexp.MustCompile(`^\d{9}$`)

var eclValues = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

var validationRules = map[string]func(s string) bool{
	"byr": func(s string) bool { return checkRange(s, 1920, 2002) },
	"iyr": func(s string) bool { return checkRange(s, 2010, 2020) },
	"eyr": func(s string) bool { return checkRange(s, 2020, 2030) },
	"hgt": func(s string) bool {
		endIndex := len(s) - 2
		if endIndex < 0 {
			return false
		}
		switch s[endIndex:] {
		case "cm":
			return checkRange(s[:endIndex], 150, 193)
		case "in":
			return checkRange(s[:endIndex], 59, 76)
		}
		return false
	},
	"hcl": func(s string) bool { return hclRe.MatchString(s) },
	"ecl": func(s string) bool {
		for _, val := range eclValues {
			if s == val {
				return true
			}
		}
		return false
	},
	"pid": func(s string) bool { return pidRe.MatchString(s) },
	"cid": func(s string) bool { return true },
}

func ProcessAndValidate(s string) bool {
	for _, fld := range strings.Fields(s) {
		val := strings.Split(fld, ":")
		fn, ok := validationRules[val[0]]
		if !ok {
			continue
		}
		if !fn(val[1]) {
			return false
		}
	}
	return true
}
