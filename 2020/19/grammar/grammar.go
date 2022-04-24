package grammar

import (
	"strings"
)

type Grammar struct {
	rules map[string][]string // this map stores rules in the following way: if 3->4 5 then rules[4 5] = {3} // todo possibly map[int][]int would be more efficient
}

func New() *Grammar {
	return &Grammar{
		rules: make(map[string][]string),
	}
}

func createKey(keyParts []string) string {
	return strings.Join(keyParts, ",")
}

func parseKey(key string) []string {
	return strings.Split(key, ",")
}

func (g *Grammar) addRule(from string, to []string) {
	key := createKey(to)
	g.rules[key] = append(g.rules[key], from)
}

func (g *Grammar) findPossibileCombinations(a, b []string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, aVal := range a {
		for _, bVal := range b {
			val, ok := g.rules[createKey([]string{aVal, bVal})]
			if !ok {
				continue
			}
			for _, v := range val {
				result[v] = struct{}{}
			}
		}
	}
	return result
}

func getTrinagleIndex(i, j, firstRowLen int) int {
	return firstRowLen*(firstRowLen+1)/2 - (firstRowLen-j)*(firstRowLen-j+1)/2 + i
}

// Works only for a grammar in Chomsky normal form
func (g *Grammar) ContainsWord(startingPoint, terminalWord string) bool {
	rowSize := len(terminalWord)
	var nonterminalTable [][]string

	// translate first row term -> nonterm
	for _, term := range terminalWord {
		nonterms, ok := g.rules[string(term)]
		if !ok {
			return false
		}
		nonterminalTable = append(nonterminalTable, nonterms)
	}

	for l := 1; l < rowSize; l++ { // traverse rows
		for k := 0; k < rowSize-l; k++ { // traverse items
			x_lk := make(map[string]struct{})
			for i := 1; i <= l; i++ {
				a_i := getTrinagleIndex(k, i-1, rowSize)
				b_i := getTrinagleIndex(k+i, l-i, rowSize)
				pc := g.findPossibileCombinations(nonterminalTable[a_i], nonterminalTable[b_i])
				for key, val := range pc {
					x_lk[key] = val
				}
			}
			var result_lk []string
			for key := range x_lk {
				result_lk = append(result_lk, key)
			}
			nonterminalTable = append(nonterminalTable, result_lk)
		}
	}
	for _, val := range nonterminalTable[len(nonterminalTable)-1] {
		if val == startingPoint {
			return true
		}
	}
	return false
}
