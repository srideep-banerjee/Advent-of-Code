package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)
	sum := 0
	
	for sc.Scan() {
		var line string = sc.Text()
		var length = utf8.RuneCountInString(line)
		var compartment1 map[byte]bool = make(map[byte]bool)
		var compartment2 map[byte]bool = make(map[byte]bool)
		for i := 0; i < length / 2; i++ {
			compartment1[line[i]] = true
		}
		
		for i := length/2; i < length; i++ {
			if(!compartment2[line[i]] && compartment1[line[i]]) {
				compartment2[line[i]] = true
				if(line[i] >= 'A' && line[i] <= 'Z') {
					sum += (int)(line[i] - 'A') + 27
				} else {
					sum += (int)(line[i] - 'a') + 1
				}
			}
		}
	}
	
	HandleError(sc.Err())
	
	fmt.Println("Day 3 Part 1 ans is",sum)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}