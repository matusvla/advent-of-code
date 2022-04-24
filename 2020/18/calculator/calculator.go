package calculator

import (
	"strconv"
	"strings"
)

func plus(i, j int) int {
	return i + j
}

func times(i, j int) int {
	return i * j
}

func evaluate(operand *string, result *int, operator func(int, int) int) {
	if *operand == "" {
		return
	}
	n, err := strconv.Atoi(*operand)
	if err != nil {
		panic(err)
	}
	*result = operator(*result, n)
	*operand = ""
}

func EvaluateExpression(e string) int {
	r, _ := evaluateExpression(e)
	return r
}

func evaluateExpression(e string) (int, int) {
	var result int
	var operand string
	operator := plus
	for i := 0; i < len(e); i++ {
		switch e[i] {
		case ' ':
			evaluate(&operand, &result, operator)
		case '*':
			evaluate(&operand, &result, operator)
			operator = times
		case '+':
			evaluate(&operand, &result, operator)
			operator = plus
		case '(':
			r, newi := evaluateExpression(e[i+1:])
			sr := strconv.Itoa(r)
			evaluate(&sr, &result, operator)
			i += newi + 1
		case ')':
			evaluate(&operand, &result, operator)
			return result, i
		default:
			operand += string(e[i])
			continue
		}
	}
	evaluate(&operand, &result, operator)
	return result, -1
}

func countResult(e string) int {
	multParts := strings.Split(e, "*")
	var subRes []int
	for i, multPart := range multParts {
		subRes = append(subRes, 0)
		addParts := strings.Split(multPart, "+")
		for _, addPart := range addParts {
			n, err := strconv.Atoi(strings.TrimSpace(addPart))
			if err != nil {
				panic(n)
			}
			subRes[i] += n
		}
	}
	res := 1
	for _, r := range subRes {
		res *= r
	}
	return res
}

// todo read and implement https://en.wikipedia.org/wiki/Recursive_descent_parser
func EvaluateExpression2(e string) int {
	e = strings.ReplaceAll(e, " ", "")
	for {
		bracketEnd := strings.Index(e, ")")
		if bracketEnd == -1 {
			return countResult(e)
		}
		startBracketCount := strings.Count(e[:bracketEnd+1], "(")
		for i := 0; i < startBracketCount-1; i++ {
			bracketEnd += strings.Index(e[bracketEnd+1:], ")") + 1
			startBracketCount = strings.Count(e[:bracketEnd+1], "(")
		}

		res := EvaluateExpression2(e[strings.Index(e, "(")+1 : bracketEnd])
		e = e[:strings.Index(e, "(")] + strconv.Itoa(res) + e[bracketEnd+1:]
	}
}
