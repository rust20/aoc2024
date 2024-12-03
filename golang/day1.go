package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1() {

	inpFile, _ := os.ReadFile(os.Args[2])

	input := string(inpFile)
	inputLines := strings.Split(input, "\n")

	l := make([]int, 1000)
	r := make([]int, 1000)

	for _, val := range inputLines {
        if len(val) == 0 {
            continue
        }
		lr := strings.Split(val, "   ")
		res, _ := strconv.Atoi(lr[0])
		l = append(l, res)
		res, _ = strconv.Atoi(lr[1])
		r = append(r, res)
	}
	sort.Ints(l)
	sort.Ints(r)

	sum := 0
	for i := range len(l) {
		if l[i]-r[i] > 0 {
			sum += l[i] - r[i]
		} else {
			sum += r[i] - l[i]
		}
	}

	count := make(map[int]int, 1000)
	for _, val := range r {
		if n, ok := count[val]; ok {
			count[val] = n + 1
		} else {
			count[val] = 1
		}
	}
	sum2 := 0
	for _, val := range l {
		if v, ok := count[val]; ok {
			sum2 += val * v
		}
	}

	fmt.Printf("star1: %d\n", sum)
	fmt.Printf("star2: %d\n", sum2)

}
