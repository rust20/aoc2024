package main

import (
	"fmt"
	"math/bits"
	"os"
	"strings"
	"time"
)

func day6() {
	time.Sleep(0)
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	input := string(inpFile)

	sum1 := 0
	sum2 := 0

	l := strings.Index(input, "\n")
	w := len(input) / (l + 1)
	_, _ = l, w

	visitMap := make([]int, len(input))
	willblockMap := make([]bool, len(input))

	loc := strings.Index(input, "^")

	dir := []int{
		-(l + 1), // up
		1,        // right
		(l + 1),  // down
		-1,       // left
	}

	dirflag := []int{
		0b1000, // up
		0b0100, // right
		0b0010, // down
		0b0001, // left
	}

	nexttileblockwillLoop := func(lcurr, lcdir, newblock int) bool {
		idir := lcdir
		icurr := lcurr

		lvisitMap := make([]int, len(input))

		for {
			if ((lvisitMap[icurr]))&dirflag[idir] > 0 {
				return true
			}
			lvisitMap[icurr] |= dirflag[idir]

			next := icurr + dir[idir]
			if next < 0 || next >= len(input) || input[next] == '\n' {
				break
			}

			if input[next] == '#' || next == newblock {
				idir = (idir + 1) & 3
			} else {
				icurr = next
			}
		}
		return false
	}

	cdir := 0
	curr := loc

	for {
		if visitMap[curr] == 0 {
			sum1++
		}
		visitMap[curr] |= dirflag[cdir]

		next := curr + dir[cdir]
		if next < 0 || next >= len(input) || input[next] == '\n' {
			break
		}

		if !willblockMap[next] {
			if nexttileblockwillLoop(loc, 0, next) {
				willblockMap[next] = true
				sum2++
			}
		}

		if input[next] == '#' {
			cdir = (cdir + 1) & 3
		} else {
			curr = next
		}
	}

	cdir = 0
	curr = loc

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)

	// http.HandleFunc("/day6", func(w http.ResponseWriter, r *http.Request) {
	// 	res, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Println("error: ", err)
	// 	}
	// 	fmt.Println("test", string(res))
	//
	// 	idxRaw := r.URL.Query().Get("idx")
	//
	// 	idx, _ := strconv.Atoi(idxRaw)
	// 	_ = idx
	//
	// 	_, err = fmt.Fprintf(w, drawmap(input, visitMap))
	//
	// 	if err != nil {
	// 		fmt.Println("error: ", err)
	// 	}
	// })

	// http.Handle("/", http.FileServer(http.Dir(".")))

	// http.ListenAndServe(":8080", nil)
}

func drawmap(m string, visitmapdir []int) string {
	result := ""
	for i, val := range m {
		if bits.OnesCount(uint(visitmapdir[i])) > 1 {
			result += "<span class=\"text-red-600\">+</span>"
		} else if visitmapdir[i]&0b1010 > 0 {
			result += "<span class=\"text-red-600\">|</span>"
		} else if visitmapdir[i]&0b0101 > 0 {
			result += "<span class=\"text-red-600\">-</span>"
		} else if string(val) == "\n" {
			result += "<br/>"
		} else {
			result += string(val)
		}
	}
	return result
}

func drawmap_pos(m string, visitmapdir, visitmapdir_temp []int, loc, dir int) string {
	result := make([]byte, len(visitmapdir))
	for i, val := range m {

		if bits.OnesCount(uint(visitmapdir[i])) > 1 {
			result[i] = '+'
		} else if visitmapdir[i]&0b1010 > 0 {
			result[i] = '|'
		} else if visitmapdir[i]&0b0101 > 0 {
			result[i] = '-'
		} else if bits.OnesCount(uint(visitmapdir_temp[i])) > 1 {
			result[i] = '+'
		} else if visitmapdir_temp[i]&0b1010 > 0 {
			result[i] = '|'
		} else if visitmapdir_temp[i]&0b0101 > 0 {
			result[i] = '-'
		} else if string(val) == "\n" {
			result[i] = '\n'
		} else {
			result[i] = byte(val)
		}

	}
	switch dir {
	case 0:
		result[loc] = '^'
	case 1:
		result[loc] = '>'
	case 2:
		result[loc] = 'v'
	case 3:
		result[loc] = '<'
	}

	return string(result)
}

func drawmap2(m string, visitmapdir []int, visitmapdir_temp []int) string {
	result := ""
	for i, val := range m {

		if bits.OnesCount(uint(visitmapdir[i])) > 1 {
			result += "<span class=\"text-red-600\">+</span>"
		} else if visitmapdir[i]&0b1010 > 0 {
			result += "<span class=\"text-red-600\">|</span>"
		} else if visitmapdir[i]&0b0101 > 0 {
			result += "<span class=\"text-red-600\">-</span>"
		} else if bits.OnesCount(uint(visitmapdir_temp[i])) > 1 {
			result += "<span class=\"text-blue-600\">+</span>"
		} else if visitmapdir_temp[i]&0b1010 > 0 {
			result += "<span class=\"text-blue-600\">|</span>"
		} else if visitmapdir_temp[i]&0b0101 > 0 {
			result += "<span class=\"text-blue-600\">-</span>"
		} else if string(val) == "\n" {
			result += "<br/>"
		} else {
			result += string(val)
		}

	}

	return result
}

func printMap(m string, visitmapdir []int) {
	println(m)

	for i, val := range m {
		_, _ = i, val

		if bits.OnesCount(uint(visitmapdir[i])) > 1 {
			fmt.Printf("+")
		} else if visitmapdir[i]&0b1010 > 0 {
			fmt.Printf("|")
		} else if visitmapdir[i]&0b0101 > 0 {
			fmt.Printf("-")
		} else {
			fmt.Printf("%c", val)
		}

	}

	println()

}
