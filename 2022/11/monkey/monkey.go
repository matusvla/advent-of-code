package monkey

import (
	"fmt"
	"math/big"
)

type Monkey struct {
	id              int
	items           map[string]int // string is big int
	operation       func(*big.Int) *big.Int
	test            func(int2 *big.Int) bool
	successTestDest *Monkey
	failTestDest    *Monkey

	monkeyBusinessScore int
}

type Monkeys []*Monkey

func NewMonkey(items []int, operation func(*big.Int) *big.Int, test func(*big.Int) bool) *Monkey {
	itemMap := make(map[string]int)
	for _, it := range items {
		itemMap[fmt.Sprint(it)]++
	}
	return &Monkey{
		items:     itemMap,
		operation: operation,
		test:      test,
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
	for _, m := range ms {
		//ms.Print()
		for item, count := range m.items {
			m.monkeyBusinessScore += count
			i := big.Int{}
			i.SetString(item, 10)
			//fmt.Printf("Monkey %v observes %v\n", m.id, item)
			worryLevel := m.operation(&i)
			//fmt.Printf("Monkey %v worries about %v\n", m.id, worryLevel.String())
			//worryLevel = new(big.Int).Div(worryLevel, big.NewInt(3)) //only for part 1
			//fmt.Printf("Monkey %v stopped worying about %v\n", m.id, worryLevel.String())
			if worryLevel.Cmp(big.NewInt(0)) == 0 {
				panic("worry level is zero")
			}
			if m.test(worryLevel) {
				//fmt.Printf("Monkey %v really throws %v to %v\n", m.id, worryLevel.String(), m.successTestDest.id)
				m.successTestDest.items[worryLevel.String()] += count
			} else {
				//fmt.Printf("Monkey %v really throws %v to %v\n", m.id, worryLevel.String(), m.failTestDest.id)
				m.failTestDest.items[worryLevel.String()] += count
			}
			delete(m.items, item)
			//ms.Print()
		}
		//ms.Print()
		//fmt.Println()
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
