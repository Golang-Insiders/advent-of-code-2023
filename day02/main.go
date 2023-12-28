package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input file path not provided")
		return
	}

	var sum int64
	for _, line := range readFileLineByLine(os.Args[1]) {
		sum += calculateLine(line)
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

type gameSet struct {
	Blue  int64
	Green int64
	Red   int64
}

func calculateLine(input string) int64 {
	// used in problem 1
	// numString := input[5:strings.Index(input, ":")]
	// num, err := strconv.ParseInt(numString, 10, 12)
	// if err != nil {
	// 	panic(err)
	// }

	var build strings.Builder
	var localNum int64
	result := []gameSet{}
	local := gameSet{}
	iterateOver := input[strings.Index(input, ":"):]
	var err error
	for i, char := range iterateOver {
		if char >= 48 && char <= 57 {
			build.WriteRune(char)
		}

		if char == ' ' && build.Len() != 0 {
			numS := build.String()
			localNum, err = strconv.ParseInt(numS, 10, 12)
			if err != nil {
				panic(err)
			}
			build.Reset()
		}

		if char == 'r' && iterateOver[i-1] == ' ' {
			local.Red = localNum
			localNum = 0
		}

		if char == 'b' && iterateOver[i-1] == ' ' {
			local.Blue = localNum
			localNum = 0
		}

		if char == 'g' && iterateOver[i-1] == ' ' {
			local.Green = localNum
			localNum = 0
		}

		if char == ';' {
			result = append(result, local)
			local = gameSet{}
		}

	}

	result = append(result, local)

	var highestBlue int64
	var highestGreen int64
	var highestRed int64
	for _, i := range result {
		if i.Blue > highestBlue {
			highestBlue = i.Blue
		}

		if i.Red > highestRed {
			highestRed = i.Red
		}

		if i.Green > highestGreen {
			highestGreen = i.Green
		}
	}
	return highestBlue * highestGreen * highestRed
}
