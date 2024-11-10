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
	fmt.Printf("answer: %d", sum)
}
