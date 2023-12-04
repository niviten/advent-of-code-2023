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

    sum := 0
    num := 0
    startIndex := -1
    endIndex := -1
    
    for index := 0; index <= len(currLine); index++ {
        if index < len(currLine) && unicode.IsDigit(rune(currLine[index])) {
            if startIndex == -1 {
                startIndex = index
            }
            endIndex = index
            digit := int(currLine[index] - '0')
            num = num * 10
            num = num + digit
        } else if num != 0 {
            startIndex = startIndex - 1
            endIndex = endIndex + 1
            if startIndex < 0 {
                startIndex = 0
            }
            if endIndex >= len(currLine) {
                endIndex = len(currLine) - 1
            }
            found := false
            for i := startIndex; i <= endIndex; i++ {
                found = false
                if !found && len(prevLine) > 0 {
                    found = isSymbol(prevLine[i])
                }
                if !found && len(currLine) > 0 {
                    found = isSymbol(currLine[i])
                }
                if !found && len(nextLine) > 0 {
                    found = isSymbol(nextLine[i])
                }
                if found {
                    break
                }
            }
            if found {
                sum = sum + num
            }
            num = 0
            startIndex = -1
            endIndex = -1
        }
    }

    return sum
}

func isSymbol(b byte) bool {
    return b != '.' && !unicode.IsDigit(rune(b))
}
