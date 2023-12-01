package main

import (
	day1_1 "eltsu7/aoc-2023/01/1"
	day1_2 "eltsu7/aoc-2023/01/2"
	"fmt"
	"time"
)

func measureTime(day int, function func() int) {
	start := time.Now()
	result := function()
	fmt.Printf("Day %d,\tResult: %d,\tExecution time: %s\n", day, result, time.Since(start).String())
}

func main() {
	measureTime(1, day1_1.Execute)
	measureTime(2, day1_2.Execute)
}
