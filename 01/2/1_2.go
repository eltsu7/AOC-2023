package day1_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var digitList = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

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
	for lineIndex := 0; lineIndex < len(line); lineIndex++ {
		if 47 < int(line[lineIndex]) && int(line[lineIndex]) < 58 {
			first = string(line[lineIndex])
			break
		}
		match := false

		for digitListIndex, digit := range digitList {
			found := true
			for digitIndex := range digit {
				if line[lineIndex+digitIndex] != digit[digitIndex] {
					found = false
					break
				}
			}
			if found {
				first = fmt.Sprint(digitListIndex + 1)
				match = true
				break
			}
		}
		if match {
			break
		}
	}
	for lineIndex := len(line) - 1; lineIndex >= 0; lineIndex-- {
		if 47 < int(line[lineIndex]) && int(line[lineIndex]) < 58 {
			last = string(line[lineIndex])
			break
		}
		match := false

		for digitListIndex, digit := range digitList {
			found := true
			if len(line)-lineIndex < len(digit) {
				continue
			}
			for digitIndex := range digit {
				if line[lineIndex+digitIndex] != digit[digitIndex] {
					found = false
					break
				}
			}
			if found {
				last = fmt.Sprint(digitListIndex + 1)
				match = true
				break
			}
		}
		if match {
			break
		}
	}
	return first, last
}
