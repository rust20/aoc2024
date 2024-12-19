package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func day8() {
	time.Sleep(0)
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	// input := strings.Split(string(inpFile), "\n")
	input := string(inpFile)
	_ = input

	l := strings.Index(input, "\n")
	w := len(input) / (l + 1)
	_, _ = l, w

	toLoc := func(n int) pair {
		return pair{n % (l + 1), n / (l + 1)}
	}
	_ = toLoc

	inmap := func(p pair) bool {
		return 0 <= p.first && p.first < l &&
			0 <= p.second && p.second < w
	}

    incpair := func(p pair, i pair) pair {
        return pair{
            p.first + i.first,
            p.second + i.second,
        }
    }

	sum1 := 0
	sum2 := 0

	antennaList := map[rune][]pair{}

	for i, val := range input {
		if ('0' <= val && val <= '9') ||
			('a' <= val && val <= 'z') ||
			('A' <= val && val <= 'Z') {
			antennaList[val] = append(antennaList[val], toLoc(i))
		}
	}

	antinode := make([]bool, len(input))
	antinode2 := make([]bool, len(input))
	for _, val := range antennaList {
		for i := range val {
			for j := range val {
				if i == j {
					continue
				}

				a := val[i]
				b := val[j]

				distx := a.first - b.first
				disty := a.second - b.second
                
                pu := pair{distx, disty}
                pd := pair{-distx, -disty}

				an1 := incpair(a, pu)
				an2 := incpair(b, pd)

				if inmap(an1) {
					antinode[an1.second*(l+1)+an1.first] = true
				}
				if inmap(an2) {
					antinode[an2.second*(l+1)+an2.first] = true
				}

                for bn1 := a; inmap(bn1); bn1 = incpair(bn1, pu) {
					antinode2[bn1.second*(l+1)+bn1.first] = true
                }
                for bn2 := b; inmap(bn2); bn2 = incpair(bn2, pd) {
					antinode2[bn2.second*(l+1)+bn2.first] = true
                }

			}
		}
	}

	for i := range antinode {
		if antinode[i] {
			sum1++
		}
	}
	for i := range antinode2 {
		if antinode2[i] {
			sum2++
		}
	}

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}
