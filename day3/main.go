package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	isNumberOrDotRx regexp.Regexp = *regexp.MustCompile(`([0-9]|\.)`)
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input file path not provided")
		return
	}

	input, err := os.ReadFile(os.Args[1])
	must(err)

	res1, err := dayThreePartOne(string(input))
	must(err)
	res2, err := dayThreePartTwo(string(input))
	must(err)

	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dayThreePartOne(input string) (string, error) {
	lines := splitLines(input)
	matrix := createMatrixFromLines(lines)
	numbers, symbols, err := walkMatrix(matrix)
	if err != nil {
		return "", err
	}

	results := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		for _, s := range symbols {
			if !numbers[i].checked && numbers[i].isNear(s.Int()) {
				results = append(results, numbers[i].no)
			}
		}
	}

	sum := sumIntSlice(results)
	return fmt.Sprintf("%d", sum), nil
}

type point struct {
	x int
	y int
}

func (p point) Int() []int {
	return []int{p.x, p.y}
}

type gear struct {
	count  int
	values []int
}

func (g gear) ratio() int {
	if len(g.values) != 2 {
		log.Fatal("invalid gear")
	}
	return g.values[0] * g.values[1]
}

func dayThreePartTwo(input string) (string, error) {
	lines := splitLines(input)
	matrix := createMatrixFromLines(lines)
	numbers, symbols, err := walkMatrix(matrix)
	if err != nil {
		return "", err
	}

	gears := make(map[point]gear)
	for i := 0; i < len(numbers); i++ {
		for _, s := range symbols {
			if s.s == "*" {
				if !numbers[i].checked && numbers[i].isNear(s.Int()) {
					if _, ok := gears[s.point]; ok {
						g := gears[s.point]
						g.count++
						g.values = append(g.values, numbers[i].no)
						gears[s.point] = g
					} else {
						gears[s.point] = gear{
							count:  1,
							values: []int{numbers[i].no},
						}
					}
				}
			}
		}
	}

	results := make([]int, 0)
	for _, g := range gears {
		if g.count == 2 {
			results = append(results, g.ratio())
		}
	}

	sum := sumIntSlice(results)
	return fmt.Sprintf("%d", sum), nil
}

func walkMatrix(matrix [][]string) ([]number, []symbol, error) {

	buf := ""
	numbers := make([]number, 0)
	symbols := make([]symbol, 0)
	lastPos := point{x: -1, y: -1}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			cursor := matrix[y][x]
			if isNumber(cursor) {
				buf += cursor
				lastPos = point{x: x, y: y}
				if !(y == len(matrix)-1 && x == len(matrix[0])-1) {
					continue
				}
			}

			if len(buf) > 0 {
				n, err := strconv.Atoi(buf)
				if err != nil {
					return nil, nil, err
				}

				numbers = append(numbers, number{
					no:  n,
					len: len(buf),
					pos: []int{lastPos.x - (len(buf) - 1), lastPos.y}, // start position of number
				})
				buf = ""
			}

			if isSymbol(cursor) {
				symbols = append(symbols, symbol{
					point: point{x: x, y: y},
					s:     cursor,
				})
			}
		}
	}

	return numbers, symbols, nil
}

func isNumber(no string) bool {
	_, err := strconv.Atoi(no)
	return err == nil
}

func isSymbol(c string) bool {
	return !isNumberOrDotRx.MatchString(c)
}

type symbol struct {
	point
	s string
}

type number struct {
	no      int
	pos     []int
	len     int
	checked bool
}

func (n *number) isNear(pos []int) bool {
	for i := 0; i < n.len; i++ {
		dx := n.pos[0] + i - pos[0]
		dy := n.pos[1] - pos[1]
		if math.Abs(float64(dx)) <= 1 && math.Abs(float64(dy)) <= 1 {
			n.checked = true
			return true
		}
	}
	return false
}

func splitLines(input string) []string {
	input = strings.TrimSuffix(input, "\n")
	return strings.Split(input, "\n")
}

func createMatrixFromLines(lines []string) [][]string {
	matrix := make([][]string, len(lines))

	for y := 0; y < len(lines); y++ {
		matrix[y] = make([]string, len(lines[y]))
		for x, c := range lines[y] {
			matrix[y][x] = string(c)
		}
	}

	return matrix
}

func sumIntSlice(ints []int) int {
	var sum int
	for _, i := range ints {
		sum += i
	}

	return sum
}
