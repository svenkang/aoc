package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	bytes, _ := ioutil.ReadFile("d5-1")
	str := string(bytes)
	lines := strings.Split(str, "\n")

	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedsInt := make([]int, len(seeds))
	for i, s := range seeds {
		si, _ := strconv.ParseInt(s, 10, 64)
		seedsInt[i] = int(si)
	}
	fmt.Println(seeds)

	sections := strings.Split(str, ":")[2:]
	maps := make([]map[int]int, len(sections))
	for i, s := range sections {
		ls := strings.Split(s, "\n")
		nwMap := map[int]int{}
		for _, l := range ls {
			values := strings.Split(l, " ")
			if !(len(values[0]) >= 1 && unicode.IsDigit([]rune(values[0])[0])) {
				continue
			}
			des := values[0]
			src := values[1]
			inc := values[2]
			desNum, _ := strconv.ParseInt(des, 10, 64)
			srcNum, _ := strconv.ParseInt(src, 10, 64)
			incNum, _ := strconv.ParseInt(inc, 10, 64)
			for m := 0; m < int(incNum); m++ {
				nwMap[int(srcNum)+m] = int(desNum) + m
			}
		}
		maps[i] = nwMap
	}

	locs := []int{}
	for _, s := range seedsInt {
		next := s
		for _, m := range maps {
			_, ok := m[next]
			if ok {
				next = m[next]
			}
		}
		locs = append(locs, next)
	}

	lowest := math.MaxInt
	for _, l := range locs {
		if l < lowest {
			lowest = l
		}
		fmt.Printf("locations %d\n", l)
	}
	fmt.Printf("answer: %d\n", lowest)
}
