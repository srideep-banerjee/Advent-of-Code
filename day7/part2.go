package day7

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sizeSlice = []int{}

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)
	
	res := -1

	sc.Scan()
	used := GetDirectorySize2(sc)
	unused := 70000000 - used
	required := 30000000 - unused
	
	sort.Slice(sizeSlice, func(i, j int) bool {
		return i < j
	})
	
	for i := 0; i < len(sizeSlice); i++ {
		if sizeSlice[i] >= required {
			res = sizeSlice[i]
			break
		}
	}

	common.HandleError(sc.Err())
	fmt.Println("Day 7 Part 2 ans is ", res)
}

func GetDirectorySize2(sc *bufio.Scanner) int {
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
			sum += GetDirectorySize2(sc)
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
	
	sizeSlice = append(sizeSlice, sum)
	
	return sum
}