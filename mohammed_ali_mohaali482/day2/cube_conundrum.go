package day2

import (
	"strconv"
	"strings"
)

const (
	blue  = 14
	green = 13
	red   = 12
)

func Solution(line string) int {
	splitted_line := strings.Split(line, " ")
	games := splitted_line[2:]
	game, _ := strconv.Atoi(splitted_line[1][:len(splitted_line[1])-1])
	var b, g, r int
	var num int
	for i := 0; i < len(games)-1; i += 2 {
		if games[i+1][0] == 'b' {
			num, _ = strconv.Atoi(games[i])
			b = num
		}
		if games[i+1][0] == 'r' {
			num, _ = strconv.Atoi(games[i])
			r = num
		}
		if games[i+1][0] == 'g' {
			num, _ = strconv.Atoi(games[i])
			g = num
		}
		if games[i+1][len(games[i+1])-1] == ';' {
			if b > blue || g > green || r > red {
				return 0
			}
			b, g, r = 0, 0, 0
		}
	}

	if b > blue || g > green || r > red {
		return 0
	}

	return game
}

func Solution2(line string) int {
	splitted_line := strings.Split(line, " ")
	games := splitted_line[2:]
	var b, g, r int
	var num int
	for i := 0; i < len(games)-1; i += 2 {
		if games[i+1][0] == 'b' {
			num, _ = strconv.Atoi(games[i])
			b = max(num, b)
		}
		if games[i+1][0] == 'r' {
			num, _ = strconv.Atoi(games[i])
			r = max(num, r)
		}
		if games[i+1][0] == 'g' {
			num, _ = strconv.Atoi(games[i])
			g = max(num, g)
		}
	}

	return r * g * b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
