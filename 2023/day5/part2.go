package day5

import (
	"bufio"
	"example/common"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)
	var inputs = [][2]int{}
	var outputs [][2]int = nil
	
	sc.Scan()
	for ind,str := range strings.Split(sc.Text()[7:], " ") {
		num, _ := strconv.Atoi(str)
		
		if ind % 2 == 0 {
			inputs = append(inputs, [2]int{num})
		} else {
			inputs[ind / 2][1] = num + inputs[ind / 2][0] - 1
		}
	}
	// fmt.Println(inputs)
	
	sc.Scan();sc.Scan()
		
	for ii := 0; ii < 7; ii++ {
		// fmt.Println("Iteration", ii,"------------------------\n")
		// fmt.Println(inputs)
		outputs = [][2]int{}
		// fmt.Println(inputs)
		var inputRangeMapped = [][2]int{}
		
		for sc.Scan() {
			var line = sc.Text()
			
			if line == "" {sc.Scan(); break}
			var lineData = strings.Split(line, " ")
			
			var inputRangeStart, _ = strconv.Atoi(lineData[1])
			var outputRangeStart, _ = strconv.Atoi(lineData[0])
			
			var rang, _ = strconv.Atoi(lineData[2])              //Not a typo
			
			var inputRangeEnd = inputRangeStart + rang - 1
			var outputRangeEnd = outputRangeStart + rang - 1
			
			//Looking for overlap between mapped source/input
			for i := 0; i < len(inputs); i++ {
				// fmt.Println("Checking overlap",inputs[i][0],inputs[i][1], inputRangeStart, inputRangeEnd)
				if CheckOverlap(inputs[i][0],inputs[i][1], inputRangeStart, inputRangeEnd) {
					// fmt.Println("Overlap")
					//Finding intersection of mapped source/input and required source/input
					intersection := GetIntersection(inputs[i][0], inputs[i][1], inputRangeStart, inputRangeEnd)
					
					//Storing input range that are mapped to do add the unmapped but required ranges later
					inputRangeMapped = append(inputRangeMapped, intersection)
					
					//Finding and storing the mapped destination ranges corresponding to the required ranges
					startOffset := intersection[0] - inputRangeStart
					endOffset := intersection[1] - inputRangeEnd
					
					outputs = append(outputs, [2]int{outputRangeStart + startOffset})
					outputs[len(outputs) - 1][1] = outputRangeEnd + endOffset
				}
			}
		}
		
		//Adding the unmapped but required source ranges to the destination
		sort.SliceStable(inputRangeMapped, func(i, j int) bool {
			return inputRangeMapped[i][1] < inputRangeMapped[j][1]
		})
		
		sort.SliceStable(inputRangeMapped, func(i, j int) bool {
			return inputRangeMapped[i][0] < inputRangeMapped[j][0]
		})
		
		// fmt.Println(inputRangeMapped)
		
		lastEnd := -1
		
		for _, rang := range inputRangeMapped {
			if lastEnd + 1 < rang[0] {
				// fmt.Println("Gap", lastEnd + 1, rang[0] - 1)
				
				for _, inp := range inputs {
					intersection := GetIntersection(inp[0], inp[1], lastEnd + 1, rang[0] - 1)
					
					if intersection[0] <= intersection[1] {
						outputs = append(outputs, intersection)
					}
					// fmt.Println("Checking",inp, "in range",lastEnd + 1,"to",rang[0] - 1)
					// fmt.Println(outputs,"\n")
				}
			}
			
			if rang[1] > lastEnd {
				lastEnd = rang[1]
			}
		}

		for _, inp := range inputs {
			if CheckOverlap(lastEnd + 1, math.MaxInt, inp[0], inp[1]) {
				outputs = append(outputs, GetIntersection(inp[0], inp[1], lastEnd + 1, math.MaxInt))
			}
		}

		inputs = outputs
		}
	
	min := -1
	for _, in := range inputs {
		// fmt.Println(len(in))
		if min == -1 || in[0] < min {
			min = in[0]
		}
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", min)
}

func CheckOverlap(startA, endA, startB, endB int) bool {
	return startA <= endB && startB <= endA
}

func GetIntersection(startA, endA, startB, endB int) [2]int {
	start := int(math.Max(float64(startA), float64(startB)))
	end := int(math.Min(float64(endA), float64(endB)))
	return [2]int{start, end}
}