package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day04() {
	fmt.Println("Day 04")

	readFile, err := os.Open("day04/input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result1 := 0
	result2 := 0

	var scoreCardsCount []int
	index := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		val, count := part1(line)
		result1 = result1 + val
		if index >= len(scoreCardsCount) {
			scoreCardsCount = append(scoreCardsCount, 0)
		}
		scoreCardsCount[index] = scoreCardsCount[index] + 1
		for i := 1; i <= count; i++ {
			for (index + i) >= len(scoreCardsCount) {
				scoreCardsCount = append(scoreCardsCount, 0)
			}
			scoreCardsCount[index+i] = scoreCardsCount[index+i] + scoreCardsCount[index]
		}
		index = index + 1
	}

	for i := 0; i < index; i++ {
		result2 = result2 + scoreCardsCount[i]
	}

	fmt.Printf("result1 -> %d\n", result1)
	fmt.Printf("result2 -> %d\n", result2)
}

func part1(line string) (int, int) {
	parts := strings.Split(line, ":")
	parts1 := strings.Split(parts[1], "|")
	winingNumbersAsString := parts1[0]
	winingNumbersAsString = strings.TrimSpace(winingNumbersAsString)
	numbersWeHaveAsString := parts1[1]
	numbersWeHaveAsString = strings.TrimSpace(numbersWeHaveAsString)

	var winingNumbers []int
	var numbersWeHave []int

	for _, str := range strings.Split(winingNumbersAsString, " ") {
		numberVal, err := strconv.Atoi(str)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		winingNumbers = append(winingNumbers, numberVal)
	}

	for _, str := range strings.Split(numbersWeHaveAsString, " ") {
		numberVal, err := strconv.Atoi(str)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		numbersWeHave = append(numbersWeHave, numberVal)
	}

	points := 0
	count := 0

	for _, winingNumber := range winingNumbers {
		for _, numberWeHave := range numbersWeHave {
			if winingNumber == numberWeHave {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
				count = count + 1
				break
			}
		}
	}

	return points, count
}
