package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	bytes, _ := os.ReadFile("d5-1")
	str := string(bytes)
	lines := strings.Split(str, "\n")

	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedsInt := make([]int64, len(seeds))
	for i, s := range seeds {
		si, _ := strconv.ParseInt(s, 10, 64)
		seedsInt[i] = int64(si)
	}
	
	seedRanges := [][]int64{}
	for i := 0; i < len(seedsInt); i += 2 {
		rangeSet := []int64{seedsInt[i], seedsInt[i] + seedsInt[i+1], 0}
		seedRanges = append(seedRanges, rangeSet)
	}

	sections := strings.Split(str, ":")[2:]
	sectionMappings := [][][]int64{}
	for _, s := range sections {
		ls := strings.Split(s, "\n")
		mxSets := [][]int64{}
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

			// represents: start, end, offset
			mxSet := []int64{int64(srcNum), int64(srcNum) + int64(incNum), int64(desNum) - int64(srcNum)}

			// represnts list of set{start, end, offset}
			mxSets = append(mxSets, mxSet)
		}
		sectionMappings = append(sectionMappings, mxSets)
	}


	// create a new mapping of ranges
	newRanges := [][][]int64{}
	for _, r := range seedRanges {
		start := r[0]

		newSectionRanges := [][]int64{}
		for _, s := range sectionMappings {
			// for each section we have a list of range sets 

			for _, set := range s {
				setStart := set[0]
				setEnd := set[1]
				setOff := set[2]

				if start >= setStart {
					n := []int64{start, setEnd, setOff }
					newSectionRanges = append(newSectionRanges, n)
				}
			}
		}
		newRanges = append(newRanges, newSectionRanges)
	}



	var lowest int64
	lowest = math.MaxInt64
	for _, r := range seedRanges {
		minSeed := r[0]
		maxSeed := r[1]
		seedOffset := r[2]
		for s := minSeed; s < maxSeed; s++ {
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
			if next < lowest {
				lowest = next
			} 
		}
		fmt.Printf("consumed: %v\n", seedOffset)
	}

	fmt.Printf("answer: %d\n", lowest)
	fmt.Printf("took: %v\n", time.Since(start))
}
