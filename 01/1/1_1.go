package day1_1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput() *bufio.Scanner {
	file, err := os.Open("./01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func Execute() int {
	scanner := readInput()

	sum := 0

	for scanner.Scan() {
		first, last := getFirstAndLastDigit(scanner.Text())
		digit, _ := strconv.Atoi(first + last)
		sum += digit
	}
	return sum
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
