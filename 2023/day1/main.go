package main

import "fmt"
import "os"
import "bufio"
import "unicode"
import "strconv"
import "strings"

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		panic(fmt.Sprintf("error reading the file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		txt := scanner.Text()
		runes := []rune(txt)

		var value int
		leftPtr := 0
		rightPtr := len(runes) - 1

		if len(runes) > 0 {
			for !unicode.IsDigit(runes[leftPtr]) && isNumericText(runes[leftPtr:rightPtr+1], false) == 0 {
				leftPtr++
			}
			for !unicode.IsDigit(runes[rightPtr]) && isNumericText(runes[leftPtr:rightPtr+1], true) == 0 {
				rightPtr--
			}
			var leftNumeric string
			if unicode.IsDigit(runes[leftPtr]) {
				leftNumeric = string(runes[leftPtr])
			} else {
				leftNumeric = fmt.Sprint(isNumericText(runes[leftPtr:rightPtr+1], false))
			}
			var rightNumeric string
			if unicode.IsDigit(runes[rightPtr]) {
				rightNumeric = string(runes[rightPtr])
			} else {
				rightNumeric = fmt.Sprint(isNumericText(runes[leftPtr:rightPtr+1], true))
			}
			value, err = strconv.Atoi(leftNumeric + rightNumeric)
			if err != nil {
				fmt.Println("value could not be parsed")
			}
		}
		total = total + value
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println("error from scanner", err)
	}
}

func isNumericText(txt []rune, isReverse bool) int {
	threeDigits := [3]string{"one", "two", "six"}
	threeDigitMap := map[string]int{"one": 1, "two": 2, "six": 6}
	fourDigits := [3]string{"four", "five", "nine"}
	fourDigitMap := map[string]int{"four": 4, "five": 5, "nine": 9}
	fiveDigits := [3]string{"three", "seven", "eight"}
	fiveDigitMap := map[string]int{"three": 3, "seven": 7, "eight": 8}

	if len(txt) < 3 {
		return 0
	}

	if len(txt) >= 3 {
		threeRunesTxt := string(txt[0:3])
		if isReverse {
			revIdx := len(txt) - 3
			if revIdx < 0 {
				revIdx = 0
			}
			threeRunesTxt = string(txt[revIdx:])
		}
		for idx, val := range threeDigits {
			if strings.Contains(val, threeRunesTxt) {
				return threeDigitMap[threeDigits[idx]]
			}
		}
	}
	if len(txt) >= 4 {
		fourRunesTxt := string(txt[0:4])
		if isReverse {
			revIdx := len(txt) - 4
			if revIdx < 0 {
				revIdx = 0
			}
			fourRunesTxt = string(txt[revIdx:])
		}
		for idx, val := range fourDigits {
			if val == fourRunesTxt {
				return fourDigitMap[fourDigits[idx]]
			}
		}
	}
	if len(txt) >= 5 {
		fiveRunesTxt := string(txt[0:5])
		if isReverse {
			revIdx := len(txt) - 5
			if revIdx < 0 {
				revIdx = 0
			}
			fiveRunesTxt = string(txt[revIdx:])
		}
		for idx, val := range fiveDigits {
			if val == fiveRunesTxt {
				return fiveDigitMap[fiveDigits[idx]]
			}
		}
	}
	return 0
}
