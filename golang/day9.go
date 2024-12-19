package main

import (
	"fmt"
	"os"
	"time"
)

func day9() {
	time.Sleep(0)
	inpFileName := os.Args[2]

	inpFile, _ := os.ReadFile(inpFileName)
	// input := strings.Split(string(inpFile), "\n")
	input := string(inpFile)
	_ = input

	sum1 := 0
	sum2 := 0

	l := 0
	for i := 0; i < len(input)-1; i++ {
		l += int(input[i] - '0')
	}

	disk := make([]int, l)
	disk2 := make([]int, l)

	i := 0
	id := 0
	mode := true
	for _, val := range input {
		n := int(val - '0')
		for range n {
			if mode {
				disk[i] = id
				disk2[i] = id
			} else {
				disk[i] = -1
				disk2[i] = -1
			}
			i++
		}
		mode = !mode
		if mode {
			id++
		}

	}

	println(l)
	lcur := 0
	rcur := l - 1
	for lcur < rcur {
		for lcur < l && disk[lcur] >= 0 {
			lcur++
		}
		for rcur >= 0 && disk[rcur] < 0 {
			rcur--
		}
		if lcur >= rcur {
			break
		}

		disk[lcur] = disk[rcur]
		disk[rcur] = -1
	}

	for i := 0; i < l; i++ {
		if disk[i] > 0 {
			sum1 += i * disk[i]
		}
	}

	lcur = 0
	rcur = l - 1
	for rcur >= 0 {

		if l < 1000 {
			for i := 0; i < len(disk2); i++ {
				if disk2[i] < 0 {
					print(".")
				} else {
					print(disk2[i])
				}
			}
			println()
			for range rcur {
				print(" ")
			}
			println("^")
		}

		for disk2[rcur] < 0 {
			rcur--
		}

		fileSize := 0
		filename := disk2[rcur]
		for rcur >= 0 && disk2[rcur] == filename {
			fileSize++
			rcur--
		}

		llcur := 0
		contSpace := 0
		for llcur <= rcur && contSpace < fileSize {
			for llcur <= rcur && disk2[llcur] >= 0 {
				llcur++
			}

			contSpace = 0
			for llcur <= rcur && disk2[llcur] < 0 {
				contSpace++
				llcur++
			}
		}
		if contSpace < fileSize {
			continue
		}

		for i := range fileSize {
			disk2[llcur-contSpace+i] = filename
			disk2[rcur+1+i] = -1
		}
	}

	for i := 0; i < l; i++ {
		if disk2[i] >= 0 {
			sum2 += i * disk2[i]
		}
	}

	println()

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}
