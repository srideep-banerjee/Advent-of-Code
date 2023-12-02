package day5

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)
	
	var stacks []Stack = []Stack{}
	
	for sc.Scan() {
		var line string = sc.Text()
		if strings.Contains(line, "1   2") {
			break
		}
		
		if len(stacks) == 0 {
			for i := 1; i<utf8.RuneCountInString(line); i+=4 {
				stacks = append(stacks, *NewStack())
			}
		}
		
		for i := 0; i * 4 + 1 < utf8.RuneCountInString(line); i++ {
			if line[i * 4 + 1] != ' ' {
				stacks[i].Push((int)(line[i * 4 + 1]))
			}
		}
	}
	
	for i := range stacks {
		stacks[i].Reverse()
	}
	
	sc.Scan()
	
	for sc.Scan() {
		var line string = sc.Text()
		var fromStrIndex = strings.Index(line, " f")
		var toStrIndex = strings.Index(line, " t")
		var moveCount, _ = strconv.Atoi(line[5:fromStrIndex])
		var from, _ = strconv.Atoi(line[fromStrIndex + 6: toStrIndex])
		var to, _ = strconv.Atoi(line[toStrIndex + 4:])
		var tempSlice = make([]int, moveCount)
		from--
		to--
		
		for i := 0; i < moveCount; i++ {
			num, _ := stacks[from].Pop()
			tempSlice[i] = num
		}
		stacks[to].Push(tempSlice...)
	}

	common.HandleError(sc.Err())
	fmt.Println("Day 5 Part 2 ans is:")
	for i := 0; i < len(stacks); i++ {
		num, _ := stacks[i].Peek()
		fmt.Print(string((byte)(num)))
	}
}