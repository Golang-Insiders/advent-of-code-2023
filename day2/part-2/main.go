package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Game struct {
	ID      int
	Content *Content
}

func (g *Game) Power() int {
	return g.Content.Blues * g.Content.Greens * g.Content.Reds
}

type Content struct {
	Blues  int
	Greens int
	Reds   int
}

//go:embed input.txt
var inputDataB []byte

var (
	gameIdRegex, _ = regexp.Compile(`Game (\d+)`)
	greenRegex, _  = regexp.Compile(`(\d+) green`)
	blueRegex, _   = regexp.Compile(`(\d+) blue`)
	redRegex, _    = regexp.Compile(`(\d+) red`)
)

func main() {
	inputData := string(inputDataB)

	powers := 0

	lines := strings.Split(inputData, "\n")
	for i, l := range lines {
		game, err := parseGame(l)
		if err != nil {
			logrus.WithError(err).Errorf("failed to parse game %d\n", i)
		}

		powers = powers + game.Power()
	}

	fmt.Println(powers)
}

func parseGame(line string) (*Game, error) {
	splitted := strings.Split(line, ":")
	if len(splitted) != 2 {
		return nil, fmt.Errorf("invalid input line: %v", splitted)
	}

	content, err := parseContent(splitted[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}

	match := gameIdRegex.FindStringSubmatch(splitted[0])
	if len(match) != 2 {
		return nil, fmt.Errorf("invalid regex match found: %+v", match)
	}

	id, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert str to int: %w", err)
	}

	return &Game{
		ID:      id,
		Content: content,
	}, nil
}

func parseContent(line string) (*Content, error) {
	sets := strings.Split(line, ";")

	maxReds, maxGreens, maxBlues := 0, 0, 0

	for _, set := range sets {
		reds, greens, blues := 0, 0, 0

		draws := strings.Split(set, ",")
		for _, draw := range draws {
			draw = strings.TrimSpace(draw)
			if strings.Contains(draw, "green") {
				match := greenRegex.FindStringSubmatch(draw)
				if len(match) != 2 {
					return nil, fmt.Errorf("invalid number of matches: %v", match)
				}

				num, _ := strconv.Atoi(match[1])
				greens = greens + num
			} else if strings.Contains(draw, "red") {
				match := redRegex.FindStringSubmatch(draw)
				if len(match) != 2 {
					return nil, fmt.Errorf("invalid number of matches: %v", match)
				}

				num, _ := strconv.Atoi(match[1])
				reds = reds + num
			} else if strings.Contains(draw, "blue") {
				match := blueRegex.FindStringSubmatch(draw)
				if len(match) != 2 {
					return nil, fmt.Errorf("invalid number of matches: %v", match)
				}

				num, _ := strconv.Atoi(match[1])
				blues = blues + num
			} else {
				return nil, fmt.Errorf("failed to parse draw: %s", draw)
			}
		}

		maxReds = max(maxReds, reds)
		maxBlues = max(maxBlues, blues)
		maxGreens = max(maxGreens, greens)
	}

	return &Content{
		Blues:  maxBlues,
		Greens: maxGreens,
		Reds:   maxReds,
	}, nil
}
