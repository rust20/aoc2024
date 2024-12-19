package main

import (
	"fmt"
	"io"
	"math/bits"
	"net/http"
	"os"
	"strconv"
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

	blockloclist := []int{}
	startloclist := []int{}
	dirloclist := []int{}

	// have normal tracing
	// create some memo to count solution for part 1
	// for part 2
	// count every single possible block positioning that will cause a loop
	// 1. detect loop
	// 2. ensure no double loop detection
	n_isloop := 0

	isloop := func(start, cdir, newblock int) bool {
		// fmt.Println("check isloop", n_isloop)
		n_isloop++
		idir := cdir
		icurr := start

		count := 0
		countturn := 0

		lvisitMap := make([]int, len(input))
		_ = lvisitMap

		for {
			// if count%100 == 0 {
			// 	res := drawmap_pos(input, visitMap, lvisitMap, icurr, idir)
			// 	println(res)
			// 	time.Sleep(100 * time.Millisecond)
			// }
			// println("w", ndir, icurr, count)
			// println("w", cdir, stalrt)
			if icurr == start && idir == cdir && count > 0 {
				if count < 10 {
					println("whatt??", count, countturn)
				}
				// println("return")
				return true
			}
			lvisitMap[icurr] |= dirflag[idir]

			inext := icurr + dir[idir]
			if inext < 0 || inext >= len(input) || input[inext] == '\n' {
				break
			}

			if input[inext] == '#' || inext == newblock {
				idir = (idir + 1) % 4
				countturn++
			} else {
				icurr = inext
				count++
			}
		}
		// println("return")
		return false
	}
	_ = isloop

	fpcount := 0
	_ = fpcount

	nexttileblockwillLoop := func(lcurr, lcdir, newblock int) (bool, int, int) {
		// idir := (lcdir + 1) & 3
		idir := lcdir
		icurr := lcurr

		lvisitMap := make([]int, len(input))
		_ = lvisitMap

		count := 0
		_ = count

		for {
			// drawmap_pos(input, visitMap, lvisitMap, curr, ndir)

			// if count%100 == 0 {
			// 	res := drawmap_pos(input, visitMap, lvisitMap, icurr, idir)
			// 	println(res)
			// 	time.Sleep(100 * time.Millisecond)
			// }
			//          count++

			if (lvisitMap[icurr])&dirflag[idir] > 0 {
				// if !isloop(icurr, idir, newblock) {
				// 	fmt.Println("false positive?")
				// 	fpcount++
				// 	// return false
				// }
				return true, icurr, idir
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
		return false, 0, 0
	}

	cdir := 0
	curr := loc

	sum3 := 0
	willblockMap2 := make([]bool, len(input))
	_ = willblockMap2

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
			if ok, a, b := nexttileblockwillLoop(loc, 0, next); ok {
				willblockMap[next] = true
				sum2++

				blockloclist = append(blockloclist, next)
				startloclist = append(startloclist, a)
				dirloclist = append(dirloclist, b)

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

	cnt_willblockmap := 0
	for i := range willblockMap {
		if willblockMap[i] {
			cnt_willblockmap++
		}
	}
	fmt.Println("checking, willblockmap: ", cnt_willblockmap)
	fmt.Println("checking, willblockmap: ", sum3)

	fmt.Println("checking, fp count: ", fpcount)
	cntfail := 0
	cntright1 := 0
	cntright2 := 0
	for i := range blockloclist {
		if ok, _, _ := nexttileblockwillLoop(loc, 0, blockloclist[i]); !ok {
			// fmt.Println("fail", i)
			cntfail++
		}
		if isloop(startloclist[i], dirloclist[i], blockloclist[i]) {
			// fmt.Println("fail", i)
			cntright1++
		}
		if ok, _, _ := nexttileblockwillLoop(startloclist[i], dirloclist[i], blockloclist[i]); ok {
			// fmt.Println("fail", i)
			cntright2++
		}

		if input[blockloclist[i]] == '#' {
			println("heyy")
		}
		if input[blockloclist[i]] == '^' {
			println("heyy")
		}

	}

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2 (should be true): ", sum2)
	fmt.Println("sum2 (should true): ", len(blockloclist))
	fmt.Println("sum2 (should be 0): ", cntfail)
	fmt.Println("sum2 (should true): ", sum2-cntfail)
	fmt.Println("check 2 (should be true): ", cntright1)
	fmt.Println("check 3 (should be true): ", cntright2)

	http.HandleFunc("/day6", func(w http.ResponseWriter, r *http.Request) {
		res, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println("test", string(res))

		idxRaw := r.URL.Query().Get("idx")

		idx, _ := strconv.Atoi(idxRaw)
		_ = idx

		_, err = fmt.Fprintf(w, drawmap(input, visitMap))

		if err != nil {
			fmt.Println("error: ", err)
		}
	})

	http.Handle("/", http.FileServer(http.Dir(".")))

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
			// result += "<span class=\"text-red-600\">|</span>"
			result[i] = '|'
		} else if visitmapdir[i]&0b0101 > 0 {
			// result += "<span class=\"text-red-600\">-</span>"
			result[i] = '-'
		} else if bits.OnesCount(uint(visitmapdir_temp[i])) > 1 {
			// result += "<span class=\"text-blue-600\">+</span>"
			result[i] = '+'
		} else if visitmapdir_temp[i]&0b1010 > 0 {
			// result += "<span class=\"text-blue-600\">|</span>"
			result[i] = '|'
		} else if visitmapdir_temp[i]&0b0101 > 0 {
			// result += "<span class=\"text-blue-600\">-</span>"
			result[i] = '-'
		} else if string(val) == "\n" {
			// result += "<br/>"
			result[i] = '\n'
		} else {
			// result += string(val)
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

func printMap(m string, visitmapdir []int, visitmap []bool) {
	_ = visitmapdir
	_ = visitmap

	println(m)

	for i, val := range m {
		_, _ = i, val

		// fmt.Printf("%v\n", visitmapdir[i])
		// _ = bits.LeadingZeros

		if bits.OnesCount(uint(visitmapdir[i])) > 1 {
			fmt.Printf("+")
		} else if visitmapdir[i]&0b1010 > 0 {
			fmt.Printf("|")
		} else if visitmapdir[i]&0b0101 > 0 {
			fmt.Printf("-")
		} else {
			fmt.Printf("%c", val)
		}

		// if visitmap[i] {
		// 	fmt.Printf("O")
		// } else {
		// 	fmt.Printf("%c", val)
		// }

	}

	println()

}
