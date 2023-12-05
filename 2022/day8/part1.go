package day8

import (
	"bufio"
	"fmt"
	"os"
	"example/common"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)
	var input = []string{}

	count := 0
	for sc.Scan() {
		var line string = sc.Text()
		input = append(input, line)
	}
	
	width, height := len(input[0]), len(input)

	var output = make([][]bool, height)
	for i := range output {
		output[i] = make([]bool, width)
	}
	
	for i := 0; i < height; i++ {
		maxH,maxHR := -1,-1
		for j := 0; j < width; j++ {
			if int(input[i][j]) > maxH {
				output[i][j] = true
				maxH = int(input[i][j])
			}
			if int(input[i][width - j -1]) > maxHR {
				output[i][width - j-1] = true
				maxHR=int(input[i][width-j-1])
			}
		}
	}
	
	for j := 0; j < width; j++ {
		maxV,maxVR := -1,-1
		for i := 0; i < height; i++ {
			if int(input[i][j]) > maxV {
				output[i][j] = true
				maxV = int(input[i][j])
			}
			if int(input[height-i-1][j]) > maxVR {
				output[height-i-1][j] = true
				maxVR=int(input[height-i-1][j])
			}
		}
	}
	
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if output[i][j] {
				count++
			}
		}
	}
	
	common.HandleError(sc.Err())
	fmt.Println("Day 8 Part 1 ans is ", count)
}