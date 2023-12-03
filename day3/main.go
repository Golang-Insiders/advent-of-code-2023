package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var movement = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func valid(x, y int, lines []string) bool {
	if x < 0 || y < 0 || x >= len(lines) || y >= len(lines[0]) {
		return false
	}
	return true
}

func check(lines []string, i, j int) bool {
	var char string
	for k := 0; k < len(movement); k++ {
		if valid(i+movement[k][0], j+movement[k][1], lines) {
			char = string(lines[i+movement[k][0]][j+movement[k][1]])
			if char != "." && !isInt(char) {
				return true
			}
		}
	}
	return false
}

func check2(lines []string, i, j int) (bool, bool, [2]int) {
	var char string
	for k := 0; k < len(movement); k++ {
		if valid(i+movement[k][0], j+movement[k][1], lines) {
			char = string(lines[i+movement[k][0]][j+movement[k][1]])
			if char != "." && !isInt(char) {
				return true, char == "*", [2]int{i + movement[k][0], j + movement[k][1]}
			}
		}
	}
	return false, false, [2]int{-1, -1}
}

func Solution(lines []string) int {
	var total int
	var num string
	var numInt int
	var char string
	var found bool
	for i := 0; i < len(lines); i++ {
		num = ""
		found = false
		for j := 0; j < len(lines[0]); j++ {
			char = string(lines[i][j])
			if isInt(char) {
				num += char
				if !found {
					found = check(lines, i, j)
				}
			} else {
				if found {
					numInt, _ = strconv.Atoi(num)
					total += numInt
				}
				num = ""
				found = false
			}
		}

		if found {
			numInt, _ = strconv.Atoi(num)
			total += numInt
		}

	}

	return total

}

func Solution2(lines []string) int {
	var total int
	var num string
	var numInt int
	var char string
	var found bool
	var star_index [2]int
	var star bool
	nums := make(map[[2]int][]int)
	for i := 0; i < len(lines); i++ {
		num = ""
		found = false
		for j := 0; j < len(lines[0]); j++ {
			char = string(lines[i][j])
			if isInt(char) {
				num += char
				if !found {
					found, star, star_index = check2(lines, i, j)
				}
			} else {
				if found {
					numInt, _ = strconv.Atoi(num)
					if star {
						nums[star_index] = append(nums[star_index], numInt)
					}
				}
				num = ""
				found = false
			}
		}

		if found {
			numInt, _ = strconv.Atoi(num)
			if star {
				nums[star_index] = append(nums[star_index], numInt)
			}
		}

	}

	for _, v := range nums {
		if len(v) == 2 {
			total += v[0] * v[1]
		}
	}

	return total
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input file path not provided")
		return
	}

	var lines []string
	for _, line := range readFileLineByLine(os.Args[1]) {
		lines = append(lines, line)
	}

	fmt.Println(Solution(lines))
	fmt.Println(Solution2(lines))
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
