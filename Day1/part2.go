package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode/utf8"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)
	
	sums := []int{0}
	curSum := 0
	
	for sc.Scan() {
		var line string = sc.Text()
		
		if utf8.RuneCountInString(line) == 0 {
			sums = append(sums, curSum)
			curSum = 0
		} else {
			num, err := strconv.Atoi(line)
			HandleError(err)
			curSum += num
		}
	}
	if curSum != 0 {
		sums = append(sums, curSum)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i]>sums[j]
	})
		
	HandleError(sc.Err())

	fmt.Println("Day 1 Part 2 ans is", sums[0] + sums[1] + sums[2])
}
