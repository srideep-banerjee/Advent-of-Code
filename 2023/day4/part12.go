package day4

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strings"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)
	
	sum := 0
	commanCount := []int{}
	for sc.Scan() {
		var line = sc.Text()
		line = line[strings.Index(line, ":") + 2:]
		input := strings.Split(line, " | ")
		set := make(map[string]bool)
		count := 0
		
		for _, str := range strings.Split(input[0], " ") {
			if str != "" {
				set[str] = true
			}
		}
		
		for _, str := range strings.Split(input[1], " ") {
			if set[str] {
				count++
			}
		}
		
		commanCount = append(commanCount, count)
	}
	
	cardCount := []int{}
	for _,_ = range commanCount {
		cardCount = append(cardCount, 1)
	}

	for i:=0;i<len(cardCount);i++ {
		sum += cardCount[i]
		for j := i + 1; j <= i + commanCount[i]; j++ {
			cardCount[j] += cardCount[i]
		}
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", sum)
}