package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	res := getWinsMultiplied(fullInput)
	fmt.Println(res)
}

func getWinsMultiplied(input string) int {
	splitInput := strings.Split(input, "\n")
	regex, _ := regexp.Compile("\\d+")
	times := regex.FindAllString(splitInput[0], -1)
	distances := regex.FindAllString(splitInput[1], -1)

	multiplied := 1

	for i := range times {
		timeInt, _ := strconv.Atoi(times[i])
		distanceInt, _ := strconv.Atoi(distances[i])

		multiplied *= getWaysToWin(timeInt, distanceInt)
	}

	return multiplied
}

func getWaysToWin(time int, distance int) int {
	speed := 0
	wins := 0
	for i := 1; i < time; i++ {
		speed += 1
		remainingTime := time - i
		travelled := speed * remainingTime
		if travelled > distance {
			wins++
		}
	}
	return wins
}

var sampleInput string = `Time:      7  15   30
Distance:  9  40  200`

var fullInput string = `Time:        53     89     76     98
Distance:   313   1090   1214   1201`
