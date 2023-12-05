package main

import (
	"regexp"
	"strconv"
	"strings"
)

// func main() {
// 	fmt.Println(validGamesTotal(fullInput))
// }

func validGamesTotal(input string) int {
	games := strings.Split(input, "\n")
	output := 0

	for _, game := range games {
		if isGameValid(game) {
			regex, _ := regexp.Compile("Game (\\d+)")
			matches := regex.FindStringSubmatch(game)
			gameNumber, _ := strconv.Atoi(matches[1])
			output += gameNumber
		}
	}

	return output
}

func isGameValid(input string) bool {
	rounds := strings.Split(input, ";")

	for _, round := range rounds {
		if !isRoundValid(round) {
			return false
		}
	}

	return true
}

func isRoundValid(input string) bool {
	regex, _ := regexp.Compile("(\\d+) (blue|red|green)")
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		count := match[1]
		color := match[2]

		countInt, _ := strconv.Atoi(count)
		if countInt > cubes[color] {
			return false
		}
	}

	return true
}

var cubes map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}
