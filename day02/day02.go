package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const redCubesLimit = 12
const greenCubesLimit = 13
const blueCubesLimit = 14

func Day02() {
    readFile, err := os.Open("day02/input")

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    result1 := 0
    result2 := 0

    for fileScanner.Scan() {
        line := fileScanner.Text()
        result1 = result1 + part1(line)
        result2 = result2 + part2(line)
    }

    fmt.Printf("result1 -> %d\n", result1)
    fmt.Printf("result2 -> %d\n", result2)
}

func part1(inputLine string) int {
    parts := strings.Split(inputLine, ":")

    gameId, err := strconv.Atoi(strings.Split(parts[0], " ")[1])

    if err != nil {
        fmt.Println(err)
        return 0
    }

    gameRecordsAsStr := strings.Split(parts[1], ";")

    for _, gameRecordAsStr := range gameRecordsAsStr {
        gameRecordParts := strings.Split(gameRecordAsStr, " ")
        numberOfParts := len(gameRecordParts)

        for i := 3; i <= 7 && i <= numberOfParts; i = i + 2 {
            count, err := strconv.Atoi(gameRecordParts[i-2])
            if err != nil {
                fmt.Println(err)
                return 0
            }

            color := gameRecordParts[i-1][0]
            if color == 'r' && count > redCubesLimit {
                return 0
            } else if color == 'g' && count > greenCubesLimit {
                return 0
            } else if color == 'b' && count > blueCubesLimit {
                return 0
            }
        }
    }

    return gameId
}

func part2(inputLine string) int {
    parts := strings.Split(inputLine, ":")
    gameRecordsAsStr := strings.Split(parts[1], ";")

    minRedCubesNeeded := 0
    minGreenCubesNeeded := 0
    minBlueCubesNeeded := 0

    for _, gameRecordAsStr := range gameRecordsAsStr {
        gameRecordParts := strings.Split(gameRecordAsStr, " ")
        numberOfParts := len(gameRecordParts)

        red := 0
        green := 0
        blue := 0

        for i := 3; i <= 7 && i <= numberOfParts; i = i + 2 {
            count, err := strconv.Atoi(gameRecordParts[i-2])
            if err != nil {
                fmt.Println(err)
                break
            }

            color := gameRecordParts[i-1][0]
            if color == 'r' && count > red {
                red = count
            } else if color == 'g' && count > green {
                green = count
            } else if color == 'b' && count > blue {
                blue = count
            }
        }

        if red > minRedCubesNeeded {
            minRedCubesNeeded = red
        }
        if green > minGreenCubesNeeded {
            minGreenCubesNeeded = green
        }
        if blue > minBlueCubesNeeded {
            minBlueCubesNeeded = blue
        }
    }

    return minRedCubesNeeded * minGreenCubesNeeded * minBlueCubesNeeded
}
