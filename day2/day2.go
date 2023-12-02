package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type gameSet struct {
	red   int
	green int
	blue  int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file argument")
	}
	file := os.Args[1]
	games := readGames(file)

	fmt.Println("Part 1:", part1(&games))
	fmt.Println("Part 2:", part2(&games))
}

func part1(games *map[int]gameSet) int {
	bag := gameSet{red: 12, green: 13, blue: 14}
	sum := 0
	for id, cubes := range *games {
		if cubes.red <= bag.red && cubes.green <= bag.green && cubes.blue <= bag.blue {
			sum += id
		}
	}
	return sum
}

func part2(games *map[int]gameSet) int {
	sum := 0
	for _, cubes := range *games {
		sum += cubes.red * cubes.green * cubes.blue
	}
	return sum
}

func readGames(filePath string) map[int]gameSet {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make(map[int]gameSet)
	for scanner.Scan() {
		id, cubes := parseLine(scanner.Text())
		games[id] = cubes
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return games
}

func parseLine(line string) (int, gameSet) {
	var id int
	parts := strings.Split(line, ": ")

	_, err := fmt.Sscanf(parts[0], "Game %d", &id)
	if err != nil {
		log.Fatal(err)
	}

	cubes := gameSet{}
	var color string
	var num int

	for _, set := range strings.Split(parts[1], "; ") {
		for _, pair := range strings.Split(set, ", ") {
			_, err := fmt.Sscanf(pair, "%d %s", &num, &color)
			if err != nil {
				log.Fatal(err)
			}

			switch color {
			case "red":
				if num > cubes.red {
					cubes.red = num
				}
			case "green":
				if num > cubes.green {
					cubes.green = num
				}
			case "blue":
				if num > cubes.blue {
					cubes.blue = num
				}
			}
		}
	}

	return id, cubes
}
