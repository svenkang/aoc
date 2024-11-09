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
	sum := part1(runes)
	sum2 := part2(runes)

	fmt.Printf("answer part 1: %d\n", sum)
	fmt.Printf("answer part 2: %d\n", sum2)
}

func getPerimeters(idx int, colSize int) []int {
	return []int{
		idx - colSize + 1,
		idx + 1,
		idx + colSize + 1,
		idx + colSize,
		idx + colSize - 1,
		idx - 1,
		idx - colSize - 1,
		idx - colSize,
	}
}

func part1(runes []rune) int {
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
	return sum
}

func part2(runes []rune) int {
	colSize := 0
	for _, r := range runes {
		colSize++
		if r == '\n' {
			break
		}
	}
	numMap := make(map[int]bool)
	currNumber := []rune{}
	sum := 0
	for i, r := range runes {
		if r == '*' {
			gears := []int{}
			visited := make([]int, len(runes))
			// look through perimeters
			// if more than 2 numbers then is gear
			ps := getPerimeters(i, colSize)
			for _, p := range ps {
				if p >= 0 && p < len(runes) {
					char := runes[p]
					if visited[p] == 0 && unicode.IsDigit(char) {
						// go all the way left until
						// it is not digit or new line
						num := []rune{}
						numLen := 1
						j := p
						for {
							if j-1 >= 0 && visited[j-1] == 0 && unicode.IsDigit(runes[j-1]) {
								numLen++
								visited[j-1] = 1
								j--
							} else {
								break
							}
						}
						if numLen > 0 {
							for i := j; i < j+numLen || unicode.IsDigit(runes[i]); i++ {
								visited[i] = 1
								num = append(num, runes[i])
							}
						}
						if len(num) > 0 {
							num, _ := strconv.ParseInt(string(num), 10, 64)
							gears = append(gears, int(num))
							if len(gears) == 2 {
								sum = sum + (gears[0] * gears[1])
							}
						}
					}
				}
			}
			currNumber = append(currNumber, r)
		} else {
			if len(currNumber) != 0 {
				num, _ := strconv.ParseInt(string(currNumber), 10, 64)
				numMap[int(num)] = false
			}
			currNumber = []rune{}
		}
	}
	return sum
}
