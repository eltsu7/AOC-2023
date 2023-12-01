package main

import (
	day1_1 "eltsu7/aoc-2023/01/1"
	day1_2 "eltsu7/aoc-2023/01/2"
	"fmt"
	"time"
)

func measureSingle(function func() int) (string, string) {
	start := time.Now()
	result := function()
	elapsed := time.Since(start)
	return fmt.Sprint(result), elapsed.String()
}

func measureThousand(function func() int) string {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		function()
	}
	elapsed := time.Since(start) / 10000
	return elapsed.String()
}

func printResults(dayName string, function func() int) {
	res, timeSignle := measureSingle(function)
	timeThousand := measureThousand(function)
	fmt.Printf("%s, result: %s, \tSingle time: %s, \t10000x average: %s\n", dayName, res, timeSignle, timeThousand)
}

func main() {
	println()
	printResults("1a", day1_1.Execute)
	printResults("1b", day1_2.Execute)
}
