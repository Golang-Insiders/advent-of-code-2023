package day1

import (
	"regexp"
	"strconv"
)

func Solution(line string) (output string, start int, end int) {
	i := 0
	j := len(line) - 1
	var err error
	for ; i < len(line); i++ {
		_, err = strconv.Atoi(string(line[i]))
		if err == nil {
			output += string(line[i])
			start = i
			break
		}
	}
	for ; j >= 0; j-- {
		_, err = strconv.Atoi(string(line[j]))
		if err == nil {
			output += string(line[j])
			end = j
			break
		}
	}
	return output, start, end
}

func Solution2(line string) string {
	digits := map[string]string{
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
	first, i, j := Solution(line)
	var result [][]int
	_min := len(line)
	_max := -1
	_min_ch := ""
	_max_ch := ""

	for key := range digits {
		reg := regexp.MustCompile(key)
		result = reg.FindAllStringIndex(line, -1)
		if len(result) == 0 {
			continue
		}
		if result[0][0] < _min {
			_min = result[0][0]
			_min_ch = key
		}

		if result[len(result)-1][0] > _max {
			_max = result[len(result)-1][0]
			_max_ch = key

		}
	}
	second := ""
	if _min >= 0 && _min < len(line) {
		second += digits[_min_ch]
	}
	if _max >= 0 && _max < len(line) {
		second += digits[_max_ch]
	}
	output := ""
	if len(second) == 2 && len(first) == 2 {
		if i < _min {
			output += string(first[0])
		} else {
			output += string(second[0])
		}
		if j > _max {
			output += string(first[1])
		} else {
			output += string(second[1])
		}

		return output
	} else if len(first) == 2 {
		return first
	} else if len(second) == 2 {
		return second
	}
	return ""
}
