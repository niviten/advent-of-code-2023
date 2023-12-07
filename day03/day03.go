package day03

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Day03() {
	readFile, err := os.Open("day03/input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result1 := 0

	var prevLine string
	var currLine string
	var nextLine string

	if fileScanner.Scan() {
		nextLine = fileScanner.Text()
	}

	for len(nextLine) > 0 {
		currLine = nextLine
		if fileScanner.Scan() {
			nextLine = fileScanner.Text()
		} else {
			nextLine = ""
		}
		result1 = result1 + part1(prevLine, currLine, nextLine)
		prevLine = currLine
	}

	fmt.Printf("result1 -> %d\n", result1)
}

func part1(prevLine string, currLine string, nextLine string) int {
	if len(currLine) == 0 {
		return 0
	}

	prevLineNumberSequences := parseLine(prevLine)
	currLineNumberSequences := parseLine(currLine)
	nextLineNumberSequences := parseLine(nextLine)

	sum := 0

	for index, ch := range currLine {
		if isSymbol(byte(ch)) {
			sum = sum + processNumberSequences(prevLineNumberSequences, index)
			sum = sum + processNumberSequences(currLineNumberSequences, index)
			sum = sum + processNumberSequences(nextLineNumberSequences, index)
		}
	}

	return sum
}

func processNumberSequences(numberSequences []NumberSequence, index int) int {
	sum := 0
	var hit bool
	var hitIndex int
	hit, hitIndex = isHit(numberSequences, index)
	if hit {
		numberSequences[hitIndex].isVisited = true
		sum = sum + numberSequences[hitIndex].value
	} else {
		hit, hitIndex = isHit(numberSequences, index-1)
		if hit {
			numberSequences[hitIndex].isVisited = true
			sum = sum + numberSequences[hitIndex].value
		}
		hit, hitIndex = isHit(numberSequences, index+1)
		if hit {
			numberSequences[hitIndex].isVisited = true
			sum = sum + numberSequences[hitIndex].value
		}
	}
	return sum
}

func isHit(numberSequences []NumberSequence, symbolIndex int) (bool, int) {
	for index, numberSequence := range numberSequences {
		if !numberSequence.isVisited {
			withinRange := (symbolIndex >= numberSequence.from) && (symbolIndex <= numberSequence.to)
			if withinRange {
				return true, index
			}
		}
	}
	return false, -1
}

func isSymbol(b byte) bool {
	return b != '.' && !unicode.IsDigit(rune(b))
}

type NumberSequence struct {
	from      int
	to        int
	value     int
	isVisited bool
}

func parseLine(line string) []NumberSequence {
	var numberSequences []NumberSequence

	num := 0
	startIndex := -1
	endIndex := -1

	for index, ch := range line {
		if unicode.IsDigit(ch) {
			digit := int(byte(ch) - '0')
			num = num * 10
			num = num + digit
			if startIndex == -1 {
				startIndex = index
			}
			endIndex = index
		} else if num != 0 {
			numberSequence := NumberSequence{
				startIndex, endIndex, num, false,
			}
			numberSequences = append(numberSequences, numberSequence)
			startIndex = -1
			endIndex = -1
			num = 0
		}
	}

	if num != 0 {
		numberSequence := NumberSequence{startIndex, endIndex, num, false}
		numberSequences = append(numberSequences, numberSequence)
	}

	return numberSequences
}
