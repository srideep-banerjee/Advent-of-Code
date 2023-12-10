package day6

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

	mult := 1
	
	var times = []int{}
	var dists = []int{}
	
	sc.Scan()
	var line = sc.Text()[11:]
	
	for _, str := range strings.Split(line, " ") {
		if str == "" {continue}
		
		time, _ := strconv.Atoi(str)
		times = append(times, time)
	}
	
	sc.Scan()
	line = sc.Text()[11:]
	
	for _, str := range strings.Split(line, " ") {
		if str == "" {continue}
		
		dist, _ := strconv.Atoi(str)
		dists = append(dists, dist)
	}
	
	for i := range times {
		count := 0
		for x := 1; x < times[i]; x++ {
			if (times[i] - x) * x > dists[i] {
				count++
			}
		}
		mult*=count
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 6 Part 1 ans is", mult)
}