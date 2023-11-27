package day2

import (
	"bufio"
	"fmt"
	"os"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	score := 0
	for sc.Scan() {
		var line string = sc.Text()
		var opCh byte = line[0] - 'A'
		var myCh byte = line[2] - 'X'

		if myCh == opCh {
			score += int(myCh) + 3 + 1
		} else if myCh == (opCh+1)%3 {
			score += int(myCh) + 6 + 1
		} else {
			score += int(myCh) + 1
		}
	}

	HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", score)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}