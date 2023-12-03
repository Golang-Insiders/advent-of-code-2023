package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	grid [][]byte

	symMap = map[byte]bool{
		'@': true,
		'#': true,
		'$': true,
		'%': true,
		'&': true,
		'*': true,
		'+': true,
		'-': true,
		'=': true,
		'/': true,
	}

	numMap = map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

type point struct {
	r, c int
}

func main() {
	loadGrid()
	partSum := 0
	gearRatioSum := 0
	gears := make(map[point][]int)
	for r := 0; r < len(grid); r++ {
		c := 0
		for c < len(grid[r]) {
			if numMap[grid[r][c]] {
				stPt := point{r, c}
				endPt := endPoint(stPt)
				if gear, ok := isPart(stPt, endPt); ok {
					n := partNum(stPt, endPt)
					partSum += n
					gears[gear] = append(gears[gear], n)
				}
				c = endPt.c
			}
			c++
		}
	}
	fmt.Println("PartSum", partSum)
	for _, gear := range gears {
		if len(gear) == 2 {
			gearRatioSum += gear[0] * gear[1]
		}
	}
	fmt.Println("GearRatioSum", gearRatioSum)
}

func endPoint(start point) point {
	i := start.c
	for i < len(grid[start.r]) {
		if !numMap[grid[start.r][i]] {
			break
		}
		if i == len(grid[start.r])-1 { //number at end of line
			return point{start.r, i}
		}
		i++
	}
	return point{start.r, i - 1}
}

func isPart(stPt, endPt point) (point, bool) {
	tl := topLeft(stPt)
	br := botRight(stPt)
	for r := tl.r; r <= br.r; r++ {
		for c := tl.c; c <= br.c; c++ {
			if symMap[grid[r][c]] {
				return point{r, c}, true
			}
		}
	}

	tl = topLeft(endPt)
	br = botRight(endPt)
	for r := tl.r; r <= br.r; r++ {
		for c := tl.c; c <= br.c; c++ {
			if symMap[grid[r][c]] {
				return point{r, c}, true
			}
		}
	}
	return point{}, false
}

func topLeft(pt point) point  { return point{max(pt.r-1, 0), max(pt.c-1, 0)} }
func botRight(pt point) point { return point{min(pt.r+1, len(grid)-1), min(pt.c+1, len(grid[pt.r])-1)} }

func partNum(stPt, endPt point) int {
	num := 0
	r := stPt.r
	for c := stPt.c; c <= endPt.c; c++ {
		num *= 10
		num += int(grid[r][c] - '0')
	}
	return num
}

func loadGrid() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	for _, line := range strings.Split(string(b), "\n") {
		grid = append(grid, []byte(line))
	}
}
