package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("d6-1")
	lines := strings.Split(string(file), "\n")

	bestTimes := []int{}
	bestTimesLine := lines[0]
	lineParts := strings.Split(bestTimesLine, ":")
	values := strings.Fields(lineParts[1])
	for _, v := range values {
		vi, _ := strconv.ParseInt(v, 10, 64)
		bestTimes = append(bestTimes, int(vi))
	}

	bestDistances := []int{}
	bestDistancesLine := lines[1]
	blineParts := strings.Split(bestDistancesLine, ":")
	bvalues := strings.Fields(blineParts[1])
	for _, v := range bvalues {
		vi, _ := strconv.ParseInt(v, 10, 64)
		bestDistances = append(bestDistances, int(vi))
	}

	bestTime := ""
	for _, b := range bestTimes {
		bestTime = bestTime + strconv.FormatInt(int64(b), 10) 
	}
	bestTimeI, _ := strconv.ParseInt(bestTime, 10, 64)

	bestDistance := ""
	for _, b := range bestDistances {
		bestDistance = bestDistance + strconv.FormatInt(int64(b), 10) 
	}
	bestDistanceI, _ := strconv.ParseInt(bestDistance, 10, 64)

	fmt.Println(bestTime)
	fmt.Println(bestDistance)

	bestTimes = []int{int(bestTimeI)}
	bestDistances = []int{int(bestDistanceI)}

	wins := 1
	for i, t := range bestTimes {
		dis := bestDistances[i]

		w := 0
		for i := 0; i < t; i ++ {
			score := i * (t - i) 
			if score > dis {
				w++
			}
		}
		wins = wins * w
	}

	fmt.Printf("answer: %d\n", wins) 
}
