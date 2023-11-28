package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	count := 0
	for sc.Scan() {
		var line string = sc.Text()
		var A = strings.Split(line, ",")[0]
		var B = strings.Split(line, ",")[1]
		var startA, startB, endA, endB int
		var err error
		startA, err = strconv.Atoi(strings.Split(A, "-")[0])
		HandleError(err)
		endA, err = strconv.Atoi(strings.Split(A, "-")[1])
		HandleError(err)
		startB, err = strconv.Atoi(strings.Split(B, "-")[0])
		HandleError(err)
		endB, err = strconv.Atoi(strings.Split(B, "-")[1])
		HandleError(err)
		
		if endA - startA > endB - startB {
			endA, endB = endB, endA
			startA, startB = startB, startA
		}
		
		if startA >= startB && endA <= endB {
			count++
		}
	}

	HandleError(sc.Err())

	fmt.Println("Day 4 Part 1 ans is ", count)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}