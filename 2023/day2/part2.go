package day2

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		var line = sc.Text()

		maxR, maxG, maxB := -1, -1, -1
		for _, str := range strings.Split(line[strings.Index(line, ":") + 2:], "; ") {
			for _, str2 := range strings.Split(str, ", ") {
				ind := strings.Index(str2, " ")
				num,_ := strconv.Atoi(str2[:ind])
				switch str2[ind + 1] {
				case 'r' :
					if num > maxR {maxR = num}
				case 'g' :
					if num > maxG {maxG = num}
				case 'b' :
					if num > maxB {maxB = num}
				}
			}
		}
		
		sum += maxR * maxG * maxB
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", sum)
}