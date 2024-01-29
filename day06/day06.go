package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day06() {
	fmt.Println("Day 06")

	// readFile, err := os.Open("day06/sample_input")
	readFile, err := os.Open("day06/input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var inputs []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		inputs = append(inputs, line)
	}

	result1 := part1(inputs)
	result2 := part2(inputs)

	fmt.Printf("result1 -> %d\n", result1)
	fmt.Printf("result2 -> %d\n", result2)
}

func part1(inputs []string) int64 {
	if len(inputs) < 2 {
		fmt.Println("input lines are lesser than 2")
		return -1
	}

	var times []int
	var distances []int

	for _, input := range inputs {
		if strings.HasPrefix(input, "Time") {
			times = parseLine(input)
		} else if strings.HasPrefix(input, "Distance") {
			distances = parseLine(input)
		}
	}

	var result int64 = 1

	for i, time := range times {
		distance := distances[i]

		t := 0
		for t <= time {
			d := (time - t) * t
			if d > distance {
				tmp := int64(time - (2 * t) + 1)
				result = result * tmp
				break
			}
			t = t + 1
		}
	}

	return result
}

func part2(inputs []string) int64 {
	if len(inputs) < 2 {
		fmt.Println("input lines are lesser than 2")
		return -1
	}

	var time int64
	var distance int64

	for _, input := range inputs {
		if strings.HasPrefix(input, "Time") {
			time = parseLineAndJoinNumbers(input)
		} else if strings.HasPrefix(input, "Distance") {
			distance = parseLineAndJoinNumbers(input)
		}
	}

	var result int64 = 1

	t := int64(0)
	for t <= time {
		d := (time - t) * t
		if d > distance {
			result = int64(time - (2 * t) + 1)
			break
		}
		t = t + 1
	}

	return result
}

func parseLine(line string) []int {
	valuesString := strings.Split(line, ":")[1]
	valuesString = strings.TrimSpace(valuesString)
	valuesStringList := strings.Fields(valuesString)
	values := make([]int, len(valuesStringList))
	for i, v := range valuesStringList {
		value, _ := strconv.ParseInt(v, 10, 32)
		values[i] = int(value)
	}
	return values
}

func parseLineAndJoinNumbers(line string) int64 {
	str := strings.Split(line, ":")[1]
	fields := strings.Fields(str)
	valueStr := strings.Join(fields, "")
	value, _ := strconv.ParseInt(valueStr, 10, 64)
	return value
}
