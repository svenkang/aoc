package main

import (
	"fmt"
	"math"
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

	var sum float64
	for i, l := range leftLocations {
		diff := math.Abs(float64(l - rightLocations[i]))
		sum = sum + diff
	}

	fmt.Println(int(sum))
}
