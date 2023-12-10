package day5

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
	var inputs = []int{}
	var outputs = []int{}
	
	sc.Scan()
	for _,str := range strings.Split(sc.Text()[7:], " ") {
		num, _ := strconv.Atoi(str)
		inputs = append(inputs, num)
		outputs = append(outputs, -1)
	}
	
	sc.Scan();sc.Scan()
	
	for ii := 0; ii < 7; ii++ {
		
		for i := 0; i < len(inputs); i++ {
			outputs[i] = -1
		}
		
		for sc.Scan() {
			var line = sc.Text()
			
			if line == "" {sc.Scan(); break}
			
			var lineData = strings.Split(line, " ")
			var inputRangeStart, _ = strconv.Atoi(lineData[1])
			var outputRangeStart, _ = strconv.Atoi(lineData[0])
			var rang, _ = strconv.Atoi(lineData[2])              //Not a typo
			
			for i, input := range inputs {
				if input >= inputRangeStart && input <= inputRangeStart + rang - 1 {
					offset := input - inputRangeStart
					outputs[i] = outputRangeStart + offset
				}
			}
		}
		
		for i, output := range outputs {
			if output != -1 {
				inputs[i] = output
			}
		}
	}
	
	min := -1
	for _, in := range inputs {
		if min == -1 || in < min {
			min = in
		}
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", min)
}