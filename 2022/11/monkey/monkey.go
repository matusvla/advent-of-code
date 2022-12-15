package monkey

import (
	"fmt"
	"sort"
)

type Monkey struct {
	id              int
	items           map[int]int
	operation       func(int) int
	testDivisor     int
	successTestDest *Monkey
	failTestDest    *Monkey

	monkeyBusinessScore int
}

type Monkeys []*Monkey

func NewMonkey(items []int, operation func(int) int, testDivisor int) *Monkey {
	itemMap := make(map[int]int)
	for _, it := range items {
		itemMap[it] = 1
	}
	return &Monkey{
		items:       itemMap,
		operation:   operation,
		testDivisor: testDivisor,
	}
}

func NewMonkeys(monkeys []*Monkey, links [][2]int) Monkeys {
	for i, link := range links {
		monkeys[i].Link(monkeys[link[0]], monkeys[link[1]])
	}
	for i, m := range monkeys {
		m.id = i
	}
	return monkeys
}

func (m *Monkey) Link(a, b *Monkey) {
	m.successTestDest = a
	m.failTestDest = b
}

func (ms Monkeys) ThrowAll() {
	commonDivisor := 1
	for _, m := range ms {
		commonDivisor *= m.testDivisor
	}

	for _, m := range ms {
		for item, count := range m.items {
			m.monkeyBusinessScore += count
			worryLevel := m.operation(item)
			// only for part 1
			//worryLevel = worryLevel / 3
			if rem := worryLevel % m.testDivisor; rem == 0 {
				m.successTestDest.items[worryLevel%commonDivisor] += count
			} else {
				m.failTestDest.items[worryLevel%commonDivisor] += count
			}
			delete(m.items, item)
		}
	}
}

func (m *Monkey) MonkeyBusinessScore() int {
	return m.monkeyBusinessScore
}

func (ms *Monkeys) Print() {
	for _, m := range *ms {
		fmt.Printf("%v : %v : %v : %v\n", m.monkeyBusinessScore, m.items, m.successTestDest.id, m.failTestDest.id)
	}
}

func (ms *Monkeys) MonkeyBusiness() int {
	// sort monkeys monkeyBusinessScore
	var scores []int
	for _, m := range *ms {
		scores = append(scores, m.monkeyBusinessScore)
	}
	sort.Ints(scores)
	return scores[len(scores)-1] * scores[len(scores)-2]
}
