package day1

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)
	sum := 0

	for sc.Scan() {
		var line string = sc.Text()
		digits := []byte{}
		
		for _, ch := range line {
			if ch>='0' && ch<='9' {
				digits = append(digits, byte(ch))
			}
		}
		
		sum += int(digits[0] - '0') * 10 + int(digits[len(digits) - 1] - '0')
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 1 Part 1 ans is", sum)
}
