package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	max := map[string]int{"red": 12, "green": 13, "blue": 14}
	f, _ := os.Open(os.Args[1])
	s := bufio.NewScanner(f)
	a := 0
	line:
	for s.Scan() {
		text := s.Text()
		game := strings.Split(text, ":")
		idx, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		sets := strings.Split(game[1], ";")
		for _, s := range sets {
			cubes := strings.Split(s, ",")
			for _, c := range cubes {
				d := strings.Split(c, " ")
				count, _ := strconv.Atoi(d[1])
				color := d[2]
				m := max[color]
				if count > m {
					continue line
				}
			}
		}
		a += idx
	}
	fmt.Println(a)

}

func part2() {
	f, _ := os.Open(os.Args[1])
	s := bufio.NewScanner(f)
	a := 0
	for s.Scan() {
		min := map[string]int{"red": 0, "green": 0, "blue": 0}
		text := s.Text()
		game := strings.Split(text, ":")
		sets := strings.Split(game[1], ";")
		for _, s := range sets {
			cubes := strings.Split(s, ",")
			for _, c := range cubes {
				d := strings.Split(c, " ")
				count, _ := strconv.Atoi(d[1])
				color := d[2]
				m := min[color]
				if count > m {
					min[color] = count
				}
			}
		}
		a += (min["red"] * min["green"] * min["blue"])
	}
	fmt.Println(a)
}
