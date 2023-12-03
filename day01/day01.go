package day01

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

func Day01() {
    readFile, err := os.Open("day01/input")

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    result1 := 0
    result2 := 0

    for fileScanner.Scan() {
        line := fileScanner.Text()
        result1 = result1 + getNumberFromDigits(line)
        result2 = result2 + getNumberFromLetters(line)
    }

    fmt.Printf("result1 -> %d\n", result1)
    fmt.Printf("result2 -> %d\n", result2)

    readFile.Close()
}

func getNumberFromDigits(inputLine string) int {
    inputLineLen := len(inputLine)

    firstIndex := 0
    lastIndex := inputLineLen - 1

    for firstIndex < inputLineLen {
        if unicode.IsDigit(rune((inputLine[firstIndex]))) {
            break
        }
        firstIndex = firstIndex + 1
    }

    for lastIndex >= 0 {
        if unicode.IsDigit(rune((inputLine[lastIndex]))) {
            break
        }
        lastIndex = lastIndex - 1
    }

    firstDigit := int(inputLine[firstIndex] - '0')
    lastDigit := int(inputLine[lastIndex] - '0')

    return (firstDigit * 10) + lastDigit
}

var digitsAsString = [...]string {
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
}

func getNumberFromLetters(inputLine string) int {
    inputLineLength := len(inputLine)

    firstDigit := -1
    for i := 0; i < inputLineLength; i++ {
        if unicode.IsDigit(rune((inputLine[i]))) {
            firstDigit = int(inputLine[i] - '0')
            break
        }
        for index, digitAsString := range digitsAsString {
            if strings.HasPrefix(inputLine[i:], digitAsString) {
                firstDigit = index
                break
            }
        }
        if firstDigit != -1 {
            break
        }
    }

    lastDigit := -1
    for i := inputLineLength-1; i >= 0; i-- {
        if unicode.IsDigit(rune((inputLine[i]))) {
            lastDigit = int(inputLine[i] - '0')
            break
        }
        for index, digitAsString := range digitsAsString {
            if strings.HasPrefix(inputLine[i:], digitAsString) {
                lastDigit = index
                break
            }
        }
        if lastDigit != -1 {
            break
        }
    }

    return (firstDigit * 10) + lastDigit
}
