package day1

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
		var line string = sc.Text()
		
		var digitNames = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		
		var min, max = -1, -1
		minDig, maxDig := int8(-1), int8(-1)
		for ind, str := range digitNames {
			strInd := strings.Index(line, str)
			if strInd != -1 && (strInd < min || min == -1) {
				min = strInd
				minDig = int8(ind)
			}
			strInd = strings.Index(line, strconv.Itoa(ind))
			if strInd != -1 && (strInd < min || min == -1) {
				min = strInd
				minDig = int8(ind)
			}
			strInd = strings.LastIndex(line, str)
			if strInd != -1 && (strInd > max || max == -1) {
				max = strInd
				maxDig = int8(ind)
			}
			strInd = strings.LastIndex(line, strconv.Itoa(ind))
			if strInd != -1 && (strInd > max || max == -1) {
				max = strInd
				maxDig = int8(ind)
			}
		}
		
		sum += int(minDig) * 10 + int(maxDig)
	}

	common.HandleError(sc.Err())
	fmt.Println("Day 8 Part 1 ans is ", sum)
}

func updateIfMin(s string, minStr *string, min *int , substr string) {
	if !strings.Contains(s, substr) {
		return
	}
	ind := strings.Index(s, substr)
	if *min == -1 || ind < *min {
		*min = ind
		*minStr = substr
	}
}

func updateIfMax(s string, maxStr *string, max *int , substr string) {
	if !strings.Contains(s, substr) {
		return
	}
	ind := strings.LastIndex(s, substr)
	if *max == -1 || ind > *max {
		*max = ind
		*maxStr = substr
	}
}