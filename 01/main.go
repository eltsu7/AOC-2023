package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func readInput() *bufio.Scanner {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func main() {
	start := time.Now()
	scanner := readInput()

	sum := 0

	for scanner.Scan() {
		first, last := getFirstAndLastDigit(scanner.Text())
		digit, _ := strconv.Atoi(first + last)
		sum += digit
	}
	println("Final sum:", sum)
	elapsed := time.Since(start)
	println("Execution time:", elapsed.String())
}

func getFirstAndLastDigit(line string) (string, string) {
	first := "-1"
	last := "-1"
	for i := 0; i < len(line); i++ {
		if 47 < int(line[i]) && int(line[i]) < 58 {
			first = string(line[i])
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if 47 < int(line[i]) && int(line[i]) < 58 {
			last = string(line[i])
			break
		}
	}
	return first, last
}
