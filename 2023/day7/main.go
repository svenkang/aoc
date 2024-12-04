package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	bytes, _ := os.ReadFile("input")
	hands := strings.Split(string(bytes), "\n")

	high := make([][]string, 0)
	one := make([][]string, 0)
	two := make([][]string, 0)
	three := make([][]string, 0)
	fullhouse := make([][]string, 0)
	four := make([][]string, 0)
	five := make([][]string, 0)
	for _, h := range hands {
		hx := strings.Split(h, " ")
		if len(hx) != 2 {
			continue
		}

		hand := hx[0]
		bid := hx[1]

		handComp := []rune(hand)
		handMap := map[rune]int{}
		hasJ := false
		for _, label := range handComp {
			if unicode.IsDigit(label) || unicode.IsLetter(label) {
				c, ok := handMap[label]
				if label == 'J' {
					hasJ = true
				}
				if ok {
					handMap[label] = c + 1
				} else {
					handMap[label] = 1
				}
			}
		}

		max := 0
		var maxChar rune
		for k, v := range handMap {
			if v > max {
				maxChar = k
				max = v
			}
		}

		pairCount := 0
		if max == 2 {
			for _, v := range handMap {
				if v == 2 {
					pairCount = pairCount + 1
				}
			}
		}

		isFullHouse := false
		if max == 3 {
			for k, v := range handMap {
				if k != maxChar && v == 2 {
					isFullHouse = true
				}
			}
		}

		switch max {
		case 1:
			high = append(high, []string{hand, bid})
		case 2:
			if pairCount == 1 {
				one = append(one, []string{hand, bid})
			} else {
				two = append(two, []string{hand, bid})
			}
		case 3:
			if isFullHouse {
				fullhouse = append(fullhouse, []string{hand, bid})
			} else {
				three = append(three, []string{hand, bid})
			}
		case 4:
			four = append(four, []string{hand, bid})
		case 5:
			five = append(five, []string{hand, bid})
		default:
		}
	}

	scoreMap := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	allTypes := [][][]string{high, one, two, three, fullhouse, four, five}
	for _, t := range allTypes {
		// order A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
		sort.Slice(t, func(i int, j int) bool {
			handA := string(t[i][0])
			handB := string(t[j][0])

			for i, a := range []rune(string(handA)) {
				rHandB := []rune(string(handB))
				if scoreMap[a] == scoreMap[rHandB[i]] {
					continue
				} else if scoreMap[a] > scoreMap[rHandB[i]] {
					return false
				} else {
					return true
				}
			}
			return true
		})
	}

	total := 0
	rank := 1
	for _, t := range allTypes {
		for _, hb := range t {
			bid := hb[1]
			bidn, _ := strconv.ParseInt(bid, 10, 64)
			total = total + (int(bidn) * rank)
			rank++
		}
	}

	fmt.Println(high)
	fmt.Println()
	fmt.Println(one)
	fmt.Println()
	fmt.Println(two)
	fmt.Println()
	fmt.Println(three)
	fmt.Println()
	fmt.Println(fullhouse)
	fmt.Println()
	fmt.Println(four)
	fmt.Println()
	fmt.Println(five)
	fmt.Println()
	fmt.Println(total)
	fmt.Println()
}
