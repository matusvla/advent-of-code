package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := ioutil.ReadFile("./testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)
	data := strings.Split(input,"\n")
	time, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	var min, bestLine int
	min = 9000000000
	busLines := strings.Split(data[1], ",")
	for _, line := range busLines {
		if line == "x" {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if time % id == 0 || id - time % id < min {
			min = id - time % id
			bestLine = id
		}
	}
	fmt.Println(min * bestLine)

	// Prepare variables for Chinese remainder theorem
	//var x int64
	var a, n []int
	for i, line := range busLines {
		id, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		a = append(a, i) // todo modulo n
		n = append(n, id)
	}

	//for i := 0; i < len(n) - 1; i++ {
	//	_, bezoutCoefa, bezoutCoefb := extendedEucleid(n[i],n[i+1])
	//}

	fmt.Println(extendedEucleid(7,13))
}

func extendedEucleid(a,b int) (int, int, int) {
	r_i := a
	r_ip1 := b
	s_i := 1
	s_ip1 := 0
	t_i := 0
	t_ip1 := 1
	for r_ip1 != 1 {
		q_ip1 := r_ip1 / r_i
		fmt.Println(r_i, s_i, t_i, q_ip1)
		pom := r_ip1
		r_ip1 = r_i - q_ip1 * r_ip1
		r_i = pom
		s_ip1, s_i = s_i - q_ip1 * s_ip1,s_ip1
		t_ip1, t_i = t_i - q_ip1 * t_ip1,t_ip1
		time.Sleep(time.Second)
	}
	return r_i, s_i, t_i
}