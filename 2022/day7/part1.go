package day7

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var res = 0

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	sc.Scan()
	GetDirectorySize(sc)

	common.HandleError(sc.Err())
	fmt.Println("Day 7 Part 1 ans is ", res)
}
func GetDirectorySize(sc *bufio.Scanner) int {
	sc.Scan()
	
	var contentArray = []string{}
	for sc.Scan() {
		line := sc.Text()
		
		if line[:4] == "$ cd" {
			break
		}
		
		contentArray = append(contentArray, line)
	}
	var sum = 0
	for i := 0; i < len(contentArray); i++ {
		if contentArray[i][:3] == "dir" {
			sum += GetDirectorySize(sc)
		} else {
			num, err := strconv.Atoi(contentArray[i][:strings.Index(contentArray[i], " ")])
			if err != nil {
				for j:=0;j<len(contentArray);j++{
					fmt.Println(contentArray[j])
				}
				common.HandleError(err)
			}
			sum += num
		}
	}
	
	sc.Scan()
	
	if sum <= 100000 {
		res += sum
	}
	
	return sum
}