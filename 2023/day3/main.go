package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

func main() {

	bytes, _ := ioutil.ReadFile("input2.txt")
	runes := []rune(string(bytes))

	colSize := 0
	for _, r := range runes {
		colSize++
		if r == '\n' {
			break
		}
	}

	numMap := make(map[int]bool)
	currNumber := []rune{}
	isPart := make([]bool, len(runes))
	sum := 0
	for i, r := range runes {

		ps := getPerimeters(i, colSize)
		// look up all parameters
		// if any of them is a symbol then this number is a part
		for _, p := range ps {
			if p >= 0 && p < len(runes) {
				x := runes[p]
				if !unicode.IsLetter(x) && !unicode.IsDigit(x) && x != '.' && x != '\n' {
					isPart[i] = true
				}
			}
		}

		if unicode.IsDigit(r) {
			currNumber = append(currNumber, r)
		} else {
			if len(currNumber) != 0 {
				num, _ := strconv.ParseInt(string(currNumber), 10, 64)
				numMap[int(num)] = false

				// go back one to the last digit of number
				// minus loop through length of digit and check if any of them is part

				// 0 1 2 3 4 5 6 7 8
				// . . . 4 8 6 . . .

				for j := i - 1; j >= i-len(currNumber); j-- {
					if j >= 0 && isPart[j] {
						numMap[int(num)] = true
					}
				}
				if numMap[int(num)] {
					sum += int(num)
				}
			}
			currNumber = []rune{}
		}
	}

	fmt.Printf("answer: %d\n", sum)
}

func getPerimeters(idx int, colSize int) []int {
	return []int{
		idx - colSize,
		idx - colSize + 1,
		idx + 1,
		idx + colSize + 1,
		idx + colSize,
		idx + colSize - 1,
		idx - 1,
		idx - colSize - 1,
	}
}
