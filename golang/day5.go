package main

import (
	"os"

	"sort"
	"strconv"
	"strings"
)

func day5() {
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	input := string(inpFile)

	inputParts := strings.Split(input, "\n\n")
	rulesRaw := strings.Split(inputParts[0], "\n")
	booksRaw := strings.Split(inputParts[1], "\n")

	rules := make(map[pair]bool, 1000)

	for _, val := range rulesRaw {
		rule := strings.Split(val, "|")

		l, _ := strconv.Atoi(rule[0])
		r, _ := strconv.Atoi(rule[1])
		rules[pair{l, r}] = true

	}

	funcB := func(b []int) int {
		return b[len(b)/2]
	}

	sum := 0
	sum2 := 0

	books := [][]int{}
	for _, val := range booksRaw {
		pages := strings.Split(val, ",")
		if val == "" {
			continue
		}

		book := []int{}
		for _, page := range pages {
			p, _ := strconv.Atoi(page)
			book = append(book, p)
		}

		books = append(books, book)

		isLess := func(i, j int) bool {
			_, ok := rules[pair{book[j], book[i]}]
			return !ok
		}

		if sort.SliceIsSorted(book, isLess) {
			sum += funcB(book)
		} else {
			sort.SliceStable(book, isLess)
			sum2 += funcB(book)
		}

	}

	println("sum1 ", sum)
	println("sum2 ", sum2)
}
