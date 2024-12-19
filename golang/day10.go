package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func day10() {
	time.Sleep(0)
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	// input := strings.Split(string(inpFile), "\n")
	input := string(inpFile)
	_ = input

	l := strings.Index(input, "\n")
	w := len(input) / (l + 1)
	_, _ = l, w

	sum1 := 0
	sum2 := 0

	trailheads := []int{}
	for i := range input {
		if input[i] == '0' {
			trailheads = append(trailheads, i)
		}
	}

	dir := []int{
		-(l + 1), // up
		1,        // right
		l + 1,    // down
		-1,       //left
	}

	var dfs func(loc map[int]int, memo []int, pos int) int
	dfs = func(loc map[int]int, memo []int, pos int) int {
		if memo[pos] > 0 {
			return memo[pos]
		}
		if input[pos] == '9' {
			loc[pos]++
			return 1
		}

		sum := 0
		for _, d := range dir {
			next := pos + d
			if next < 0 || next >= len(input) || input[next] == '\n' || input[next]-input[pos] != 1 {
				continue
			}
			sum += dfs(loc, memo, next)
		}
		memo[pos] = sum

		return sum
	}

	for _, start := range trailheads {
		memo := make([]int, len(input))
		loc := map[int]int{}
        res:=dfs(loc, memo, start)
		sum1 += len(loc)
        sum2 += res
	}

	println()

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}
