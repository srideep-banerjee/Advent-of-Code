package day6

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
	
	count := 0
		
	var times = ""
	var dists = ""
	
	sc.Scan()
	var line = sc.Text()[11:]
	
	for _, str := range strings.Split(line, " ") {
		if str == "" {continue}
		
		times += str
	}
	
	sc.Scan()
	line = sc.Text()[11:]
	
	for _, str := range strings.Split(line, " ") {
		if str == "" {continue}
		
		dists += str
	}
	
	time, _ := strconv.Atoi(times)
	dist, _ := strconv.Atoi(dists)
	
	for x := 1; x < time; x++ {
		if (time - x) * x > dist {
			count++
		}
	}
	
	common.HandleError(sc.Err())

	fmt.Println("Day 6 Part 2 ans is", count)
}