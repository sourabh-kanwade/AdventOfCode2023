package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	inputFile, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(inputFile)
	var total int64
	for scanner.Scan() {
		line := scanner.Text()
		var firstNum rune
		var lastNum rune
		for _, char := range line {
			if unicode.IsNumber(char) {
				if firstNum == 0 {
					firstNum = char
				}
				lastNum = char
			}
		}
		sum := string(firstNum) + string(lastNum)
		num, err := strconv.ParseInt(sum, 10, 64)
		check(err)
		total += num
	}
	fmt.Println(total)
}
