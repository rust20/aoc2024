package main

import (
	"fmt"
	"time"
)

func main() {
    // duration(day1)
    // duration(day4)
    // duration(day5)
    // duration(day6)
    // duration(day7)
    // duration(day8)
    // duration(day9)
    // duration(day10)
    duration(day11)
    // mainother_day6()
}

func duration(f func()) {
    start := time.Now()
    f()
    duration := time.Since(start)
    fmt.Printf("duration: %+v\n", duration)

}
