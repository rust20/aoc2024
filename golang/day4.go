package main

import (
	"os"
	"strings"
)

func day4() {
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	input := string(inpFile)

	l := strings.Index(input, "\n")
	h := len(input) / l - 1

	check_h := func(x, y int) int {
		count := 0
        l := l+1
		pos := x*l + y
		// hori
		// ihori
		if input[pos:pos+4] == "XMAS" {
			count++
		} else if input[pos:pos+4] == "SAMX" {
			count++
		}

		return count
	}
	check_v := func(x, y int) int {
		count := 0
        l := l+1
		pos := x*l + y

		s := string([]byte{
            input[pos],
			input[pos+l],
			input[pos+2*l],
			input[pos+3*l],
		})

		// hori
		// ihori
		if s == "XMAS" {
			count++
		} else if s == "SAMX" {
			count++
		}

		return count
	}
	check_d := func(x, y int) int {
		count := 0
        l := l+1
		pos := x*(l) + y

		s := string([]byte{
            input[pos      ],
			input[pos+  l+1],
			input[pos+2*l+2],
			input[pos+3*l+3],
		})

		// diagd
		// idiagd
		if s == "XMAS" {
			count++
		} else if s == "SAMX" {
			count++
		}

		s = string([]byte{
            input[pos    +3],
			input[pos+1*l+2],
			input[pos+2*l+1],
			input[pos+3*l+0],
		})

		// diagu
		// idiagu
		if s == "XMAS" {
			count++
		} else if s == "SAMX" {
			count++
		}

		return count
	}


	check_xmas := func(x, y int) int {
		count := 0
        l := l+1
		pos := x*(l) + y

		s := string([]byte{
            input[pos      ],
			input[pos+   +2],
			input[pos+  l+1],
			input[pos+2*l  ],
			input[pos+2*l+2],
		})

		// diagd
		// idiagd
		if s == "SSAMM" {
			count++
		} else if s == "SMASM" {
			count++
		} else if s == "MMASS" {
			count++
		} else if s == "MSAMS" {
			count++
		}

		return count
	}

	sum := 0
	for i := range l {
		for j := range h - 3 {
			sum += check_h(i, j)
		}
	}
	for i := range l - 3 {
		for j := range h {
			sum += check_v(i, j)
		}
	}
	for i := range l - 3 {
		for j := range h - 3 {
			sum += check_d(i, j)
		}
	}

    sum2 := 0
	for i := range l - 2 {
		for j := range h - 2 {
			sum2 += check_xmas(i, j)
		}
	}

	println("star1: ", sum)
	println("star2: ", sum2)
}
