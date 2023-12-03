package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang-insiders/advent-of-code-2023/mohammed_ali_mohaali482/day3"
)

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(day3.Solution(lines))
}
