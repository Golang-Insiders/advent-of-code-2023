package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

type partNumber struct {
	number    int
	row       int
	col       int
	end       int
	hasSymbol bool
}

type symbol struct {
	symbol        rune
	row           int
	col           int
	adjacentParts []*partNumber
}
type symbolTab map[int]map[int]*symbol

func hasSymbol(tab symbolTab, row, col int, p partNumber) bool {
	if cols, ok := tab[row]; ok {
		s, ok := cols[col]
		if ok {
			s.adjacentParts = append(s.adjacentParts, &p)
			cols[col] = s
		}
		return ok
	}
	return false
}
func scanSymbol(tab symbolTab, p partNumber) bool {
	rowStart, rowEnd := max(p.row-1, 0), p.row+1
	colStart, colEnd := max(p.col-1, 0), p.end+1

	for i := rowStart; i <= rowEnd; i++ {
		for j := colStart; j <= colEnd; j++ {
			ok := hasSymbol(tab, i, j, p)
			// fmt.Println(p.number, i, j, ok)
			if ok {
				return true
			}
		}
	}
	return false
}

func main() {
	buf := bytes.NewReader(input)
	scan := bufio.NewScanner(buf)

	m := [][]rune{}
	parts := []partNumber{}
	symbols := make(symbolTab)
	symbolList := []*symbol{}

	for scan.Scan() {
		text := scan.Text()
		row := len(m)
		m = append(m, []rune(text))

		slice := make([]rune, 0, 3)
		var col int
		for i, a := range text {
			col = i
			if a >= '0' && a <= '9' {
				slice = append(slice, a)
				continue
			}
			if v, err := strconv.Atoi(string(slice)); err == nil {
				parts = append(parts, partNumber{number: v, row: row, col: col - len(slice), end: col - 1})
				slice = slice[:0]
			}

			if a != '.' {
				cols, ok := symbols[row]
				if !ok {
					cols = make(map[int]*symbol)
				}
				s := &symbol{row: row, col: col, symbol: a}
				cols[col] = s
				symbols[row] = cols
				symbolList = append(symbolList, s)
			}
		}
		if v, err := strconv.Atoi(string(slice)); err == nil {
			parts = append(parts, partNumber{number: v, row: row, col: col - len(slice), end: col - 1})
			slice = slice[:0]
		}
	}

	sum := 0
	for i, p := range parts {
		ok := scanSymbol(symbols, p)
		parts[i].hasSymbol = ok
		if ok {
			sum += p.number
		}
	}

	sumGears := 0
	for _, s := range symbolList {
		if s.symbol == '*' && len(s.adjacentParts) == 2 {
			sumGears += s.adjacentParts[0].number * s.adjacentParts[1].number
		}
	}

	fmt.Println(m)
	fmt.Println(parts)
	fmt.Println(symbols)
	fmt.Println(symbolList)
	fmt.Println(sum)
	fmt.Println(sumGears)
}
