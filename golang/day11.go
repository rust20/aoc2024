package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day11() {
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	// input := strings.Split(string(inpFile), "\n")
	input := string(inpFile)
	_ = input

	sum1 := 0
	sum2 := 0

	raw := strings.TrimSpace(input)
	raw2 := strings.Split(raw, " ")
	stones := []int{}
	for i := range raw2 {
		res, _ := strconv.Atoi(raw2[i])
		stones = append(stones, res)
	}

	digit := func(n int) int {
		if n == 0 {
			return 1
		}
		res := 0
		for n > 0 {
			res++
			n /= 10
		}
		return res
	}

	split := func(n int) (int, int) {

		d := digit(n) / 2
		t := int(math.Pow10(d))

		return n / t, n % t
	}

	memo := map[pair]int{}
	var fn func(stones int, level int) int
	fn = func(stone int, level int) int {
		if level == 0 {
			return 1
		}

		if memoval, ok := memo[pair{stone, level}]; ok {
			return memoval
		}

		sum := 0
		if stone == 0 {
			res := fn(1, level-1)
			sum += res
		} else if digit(stone)&1 == 0 {
			d1, d2 := split(stone)
			sum += fn(d1, level-1)
			sum += fn(d2, level-1)
		} else {
			sum += fn(stone*2024, level-1)
		}
		memo[pair{stone, level}] = sum
		return sum
	}

    for i := range stones {
        sum1 += fn(stones[i], 25)
    }
    for i := range stones {
        sum2 += fn(stones[i], 75)
    }

	println()

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}
