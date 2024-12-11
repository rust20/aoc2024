package main

import (
	"fmt"
	"time"
)

func main() {
    // duration(day1)
    // duration(day4)
    duration(day5)
}

func duration(f func()) {
    start := time.Now()
    f()
    duration := time.Since(start)
    fmt.Printf("duration: %+v\n", duration)

}
