package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var spelledDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	file, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += getNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func part2() int {
	file, err := os.Open("./day01/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := processSpelledNumbers(scanner.Text())
		total += getNumber(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func processSpelledNumbers(line string) string {
	for spelledDigit, digit := range spelledDigits {
		line = strings.ReplaceAll(line, spelledDigit,
			spelledDigit[:1]+digit+spelledDigit[len(spelledDigit)-1:])
	}
	return line
}

func getNumber(line string) int {
	digits := []rune{}
	for _, char := range line {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	number, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
	if err != nil {
		log.Fatal(err)
	}

	return number
}
