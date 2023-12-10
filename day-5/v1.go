package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getLowestLocationNumber(input string) int {
	splitInput := strings.Split(input, "map:")
	maps := []map[int]int{}

	for i := 1; i < len(splitInput); i++ {
		maps = append(maps, convertMap(splitInput[i]))
		fmt.Println("created map ", i)
	}

	regex, _ := regexp.Compile("\\d+")
	seeds := regex.FindAllString(splitInput[0], -1)

	lowest := getLocationNumber(seeds[0], maps)

	for i := 1; i < len(seeds); i++ {
		curr := getLocationNumber(seeds[i], maps)
		if curr < lowest {
			lowest = curr
		}
	}

	return lowest
}

func convertMap(input string) map[int]int {
	regex, _ := regexp.Compile("((\\d+) (\\d+) (\\d+))")
	mapRanges := regex.FindAllStringSubmatch(input, -1)

	corresponds := map[int]int{}

	for _, mapRange := range mapRanges {
		fmt.Println("creating map ", mapRange[0])
		destRange, _ := strconv.Atoi(mapRange[2])
		sourceRange, _ := strconv.Atoi(mapRange[3])
		count, _ := strconv.Atoi(mapRange[4])
		for i := sourceRange; i < sourceRange+count; i++ {
			corresponds[i] = destRange
			destRange++
		}
	}

	return corresponds
}

func getLocationNumber(seed string, maps []map[int]int) int {
	seedInt, _ := strconv.Atoi(seed)

	curr := seedInt
	ok := true

	for _, seedMap := range maps {
		_, ok = seedMap[curr]
		if !ok {
			continue
		}
		curr = seedMap[curr]
	}

	return curr
}
