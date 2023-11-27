package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/srideep-banerjee/AdventOfCode/day2/part2"
)

func main() {
	file, err := os.Open("input.txt")
	handleError(err)
	var sc = bufio.NewScanner(file)
	part2.PrintAns()
	
	score := 0
	for sc.Scan() {
		var line string = sc.Text();
		var opCh byte = line[0] - 'A'
		var myCh byte = line[2] - 'X'

		if(myCh == opCh) {
			score += int(myCh) + 3 + 1
		} else if(myCh == (opCh + 1) % 3) {
			score += int(myCh) + 6 + 1
		} else {
			score += int(myCh) + 1
		}
	}
	
	handleError(sc.Err())
	handleError(file.Close())
	
	fmt.Println("Score is ", score)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
