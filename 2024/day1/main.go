package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input")
	lines := strings.Split(string(bytes), "\n")
	
	leftLocations := []int{}
	rightLocations := []int{}
	for _, l := range lines {
		pair := strings.Split(l, "   ")
		if len(pair) == 2 {
			leftLocation, _ := strconv.Atoi(pair[0])
			rightLocation, _ := strconv.Atoi(pair[1])
			leftLocations = append(leftLocations, int(leftLocation))
			rightLocations = append(rightLocations, int(rightLocation))
		}
	}
	sort.Slice(leftLocations, func(i, j int) bool {
		return leftLocations[i] < leftLocations[j]
	})

	sort.Slice(rightLocations, func(i, j int) bool {
		return rightLocations[i] < rightLocations[j]
	})

	locMap := map[int]int{}
	for _, l := range leftLocations {
		v, ok := locMap[l]
		if ok {
			locMap[l] = v + 1
		} else {
			locMap[l] = 1
		}
	}
	locMapRight := map[int]int{}
	for _, l := range rightLocations {
		v, ok := locMapRight[l]
		if ok {
			locMapRight[l] = v + 1
		} else {
			locMapRight[l] = 1
		}
	}

	var sum2 int
	for _, l := range leftLocations {
		vr, ok := locMapRight[l]
		mul := 0
		if ok {
			mul = l * vr
		}	
		sum2 = sum2 + mul
	}

	fmt.Println(int(sum2))
}
