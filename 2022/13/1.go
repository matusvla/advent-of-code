package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allPackets []ListItem
	var result1 int
	i := 0
	for scanner.Scan() {
		i++
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		a := NewListListItem(line1)
		b := NewListListItem(line2)
		cmp := compare([]ListItem{a}, []ListItem{b}) > 0
		if cmp {
			result1 += i
		}
		scanner.Scan() // skip empty line
		allPackets = append(allPackets, a, b)
	}
	fmt.Println("PART1:", result1)

	// PART 2
	div1 := NewListListItem("[[2]]")
	div2 := NewListListItem("[[6]]")
	allPackets = append(allPackets, div1, div2)

	sort.SliceStable(allPackets, func(i, j int) bool {
		return compare([]ListItem{allPackets[i]}, []ListItem{allPackets[j]}) > 0
	})

	result2 := 1
	for i, p := range allPackets {
		if p == div1 {
			result2 *= i + 1
		}
		if p == div2 {
			result2 *= i + 1
		}
	}
	fmt.Println("PART2:", result2)
}

type ListItem interface {
	Values() []ListItem
}

type IntListItem int

func (i IntListItem) Values() []ListItem {
	return []ListItem{i}
}

type ListListItem struct {
	vals []ListItem
}

func NewListListItem(s string) *ListListItem {
	var currentVal string
	var currentList *ListListItem = &ListListItem{}
	var prevListStack []*ListListItem
	for _, c := range s {
		switch c {
		case '[':
			prevListStack = append(prevListStack, currentList)
			currentList = new(ListListItem)
			currentVal = ""
		case ']':
			if currentVal != "" {
				n, err := strconv.Atoi(currentVal)
				if err != nil {
					panic(err)
				}
				currentList.vals = append(currentList.vals, IntListItem(n))
			}
			if prevListStack == nil {
				panic("unbalanced brackets")
			}
			prevList := prevListStack[len(prevListStack)-1]
			prevList.vals = append(prevList.vals, currentList)
			currentList = prevList
			prevListStack = prevListStack[:len(prevListStack)-1]
			currentVal = ""
		case ',':
			if currentVal != "" {
				n, err := strconv.Atoi(currentVal)
				if err != nil {
					panic(err)
				}
				currentList.vals = append(currentList.vals, IntListItem(n))
			}
			currentVal = ""
		default:
			currentVal += string(c)
		}
	}
	if len(prevListStack) != 0 {
		panic("unbalanced brackets 2")
	}
	return currentList
}

func (l *ListListItem) Values() []ListItem {
	return l.vals
}

func compare(aList, bList []ListItem) int {
	if len(aList) == 0 && len(bList) > 0 {
		return 1
	}
	for i, a := range aList {
		if len(bList) <= i {
			return -1
		}
		b := bList[i]
		switch a.(type) {
		case IntListItem:
			switch b.(type) {
			case IntListItem:
				if a != b {
					return int(b.(IntListItem) - a.(IntListItem))
				}
			case *ListListItem:
				if res := compare(a.Values(), b.Values()); res != 0 {
					return res
				}
			}
		case *ListListItem:
			switch b.(type) {
			case IntListItem:
				if res := compare(a.Values(), b.Values()); res != 0 {
					return res
				}
			case *ListListItem:
				if res := compare(a.Values(), b.Values()); res != 0 {
					return res
				}
			}
		default:
			panic(fmt.Sprintf("unknown type %T", a))
		}
		if i == len(aList)-1 && len(bList) > i {
			return 1
		}
	}
	return 0
}
