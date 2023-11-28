package day1

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	maxSum := 0
	curSum := 0
	for sc.Scan() {
		var line string = sc.Text()
		
		if utf8.RuneCountInString(line) == 0 {
			curSum = 0
		} else {
			num, err := strconv.Atoi(line)
			common.HandleError(err)
			curSum += num
			if curSum > maxSum {
				maxSum = curSum
			}
		}
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 1 Part 1 ans is", maxSum)
}
