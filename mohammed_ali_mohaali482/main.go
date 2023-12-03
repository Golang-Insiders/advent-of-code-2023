package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang-insiders/advent-of-code-2023/mohammed_ali_mohaali482/day2"
)

func main() {
	// trebuchet_solution()
	// cube_conundrum()
	cube_conundrum2()
}

// func trebuchet_solution() {
// 	file, err := os.Open("day1/input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	ans := 0
// 	var num int
// 	for scanner.Scan() {
// 		num, _ = strconv.Atoi(day1.Solution2(scanner.Text()))
// 		ans += num
// 	}

// 	fmt.Println(ans)
// }

// func cube_conundrum() {
// 	file, err := os.Open("day2/input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	ans := 0
// 	var num int
// 	for scanner.Scan() {
// 		num = day2.Solution(scanner.Text())
// 		ans += num
// 	}

// 	fmt.Println(ans)
// }

func cube_conundrum2() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	var num int
	for scanner.Scan() {
		num = day2.Solution2(scanner.Text())
		ans += num
	}

	fmt.Println(ans)
}
