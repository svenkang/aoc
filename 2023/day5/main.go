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

	sections := strings.Split(str, ":")[2:]
	sectionMappings := [][][]int{}
	for _, s := range sections {
		ls := strings.Split(s, "\n")
		mxSets := [][]int{}
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

			mxSet := []int{int(srcNum), int(srcNum) + int(incNum), int(desNum) - int(srcNum)}
			mxSets = append(mxSets, mxSet)
		}
		sectionMappings = append(sectionMappings, mxSets)
	}

	locs := []int{}
	for _, s := range seedsInt {
		next := s
		for _, sets := range sectionMappings {
			mapFound := false
			for _, s := range sets {
				min := s[0]
				max := s[1]
				offset := s[2]
				if !mapFound && next >= min && next <= max {
					next = next + offset
					mapFound = true
				}
			}
		}
		locs = append(locs, next)
	}

	lowest := math.MaxInt
	for _, l := range locs {
		if l < lowest {
			lowest = l
		}
	}
	fmt.Printf("answer: %d\n", lowest)
}
