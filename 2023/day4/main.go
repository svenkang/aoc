package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("d4-2")
	str := string(bytes)
	lines := strings.Split(str, "\n")
	sum1 := part1(lines)
	sum2 := part2(lines)

	fmt.Printf("answer: %d\n", sum1)
	fmt.Printf("answer2: %d\n", sum2)
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		cols := strings.Split(line, ":")
		if len(cols) > 1 {
			numCols := strings.Split(cols[1], "|")
			wins := strings.Split(numCols[0], " ")
			winMap := map[string]bool{}
			for _, w := range wins {
				winMap[string(w)] = false
			}
			rounds := strings.Split(numCols[1], " ")
			for _, r := range rounds {
				if string(r) == "" {
					continue
				}
				_, ok := winMap[string(r)]
				if ok {
					winMap[string(r)] = true
				}
			}
			i := 0
			for _, v := range winMap {
				if v {
					if i == 0 {
						i = 1
					} else {
						i = i * 2
					}
				}
			}
			sum += i
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	cardsCount := make([]int, len(lines) - 1)
	for i, _ := range cardsCount {
		cardsCount[i] = 1
	}

	for ci, line := range lines {
		cols := strings.Split(line, ":")
		if len(cols) > 1 {
			numCols := strings.Split(cols[1], "|")
			wins := strings.Split(numCols[0], " ")
			winMap := map[string]bool{}
			for _, w := range wins {
				winMap[string(w)] = false
			}
			rounds := strings.Split(numCols[1], " ")
			for _, r := range rounds {
				if string(r) == "" {
					continue
				}
				_, ok := winMap[string(r)]
				if ok {
					winMap[string(r)] = true
				}
			}
			winCounts := 0
			for _, v := range winMap {
				if v {
					winCounts++
				}
			}
			for j := 0; j < winCounts; j++ {
				holdings := cardsCount[ci]
				cardsCount[ci + j + 1] = cardsCount[ci+ j + 1] + holdings
			}
		}
	}

	for _, v := range cardsCount {
		sum = sum + v
	}

	return sum
}
