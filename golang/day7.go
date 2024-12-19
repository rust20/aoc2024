package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func rec(arr []string, n, i, target int) bool {
	if i == len(arr) {
		return n == target
	}

	num, _ := strconv.Atoi(arr[i])

	a := n + num
	b := n * num

	return rec(arr, a, i+1, target) || rec(arr, b, i+1, target)
}

func day7() {
	time.Sleep(0)
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	input := strings.Split(string(inpFile), "\n")

	sum1 := 0
	sum2 := 0

	var rec func(arr []string, n, i, target int) bool
	rec = func(arr []string, n, i, target int) bool {
		if i == len(arr) {
			return n == target
		}

		num, _ := strconv.Atoi(arr[i])

		a := n + num
		b := n * num

		return rec(arr, a, i+1, target) || rec(arr, b, i+1, target)
	}

	var rec2 func(arr []string, n, i, target int) bool
	rec2 = func(arr []string, n, i, target int) bool {
		if i == len(arr) {
			return n == target
		}

		num, _ := strconv.Atoi(arr[i])

		shift := int(math.Pow10(len(arr[i])))
		a := n + num
		b := n * num
		c := n*shift + num

		return rec2(arr, a, i+1, target) || rec2(arr, b, i+1, target) || rec2(arr, c, i+1, target)
	}

	for i := range input {
		if input[i] == "" {
			continue
		}
		t := strings.Split(input[i], ": ")
		total, _ := strconv.Atoi(t[0])

		arr := strings.Split(t[1], " ")
		n, _ := strconv.Atoi(arr[0])

		if rec(arr, n, 1, total) {
			sum1 += total
		}
		if rec2(arr, n, 1, total) {
			sum2 += total
		}

	}

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}
