package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var mapDigits map[string]int = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("input file path not provided")
		return
	}

	var sum int
	for _, line := range readFileLineByLine(os.Args[1]) {
		sum += fetchNumberFromString(line)
	}

	fmt.Println(sum)
}

func readFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func fetchNumberFromString(input string) int {
	var intSlice []int

	// find all digit occurrances
	for index, char := range input {
		if char >= 48 && char <= 57 {
			intSlice = append(intSlice, index)
		}
	}

	var charSlice []int
	charMap := make(map[int]int)

	// find all string occurrances
	for stringDigit, intDigit := range mapDigits {
		start := strings.Index(input, stringDigit)
		last := strings.LastIndex(input, stringDigit)
		if start != -1 {
			charSlice = append(charSlice, start)
			charMap[start] = intDigit
		}

		if last != -1 {
			charSlice = append(charSlice, last)
			charMap[last] = intDigit
		}
	}

	sort.Ints(charSlice)

	var first, last int

	if len(intSlice) == 0 && len(charSlice) == 0 {
		return 0
	}

	if (len(intSlice) == 0) || (len(charSlice) != 0 && charSlice[0] < intSlice[0]) {
		first = charMap[charSlice[0]]
	} else {
		first = int(input[intSlice[0]]) - 48
	}

	if (len(intSlice) == 0) || (len(charSlice) != 0 && charSlice[len(charSlice)-1] > intSlice[len(intSlice)-1]) {
		last = charMap[charSlice[len(charSlice)-1]]
	} else {
		last = int(input[intSlice[len(intSlice)-1]]) - 48
	}

	return (first * 10) + last
}
