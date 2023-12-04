package day2

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	sum := 0
	gameNo := 0
	GameLoop:
	for sc.Scan() {
		var line = sc.Text()
		gameNo++

		for _, str := range strings.Split(line[strings.Index(line, ":") + 2:], "; ") {
			for _, str2 := range strings.Split(str, ", ") {
				ind := strings.Index(str2, " ")
				num,_ := strconv.Atoi(str2[:ind])
				switch str2[ind + 1] {
				case 'r' :
					if num > 12 {continue GameLoop}
				case 'g' :
					if num > 13 {continue GameLoop}
				case 'b' :
					if num > 14 {continue GameLoop}
				}
			}
		}
		
		sum += gameNo
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", sum)
}