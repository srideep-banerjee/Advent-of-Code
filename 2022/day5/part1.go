package day5

import (
	"bufio"
	"errors"
	"example/common"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Stack struct {
	arr    []int
	length int
}

func NewStack() *Stack {
	var stack = Stack{[]int{}, 0}
	return &stack
}

func (s *Stack) Push(num ...int) {
	for i := len(num) - 1; i >= 0; i-- {
		if len(s.arr) == s.length {
			s.arr = append(s.arr, num[i])
			s.length++
		} else {
			s.arr[s.length] = num[i]
			s.length++
		}
	}
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("cannot pop from empty stack")
	}
	s.length--
	return s.arr[s.length], nil
}

func (s *Stack) Peek() (int, error) {
	if s.length == 0 {
		return 0, errors.New("cannot peek empty stack")
	}
	return s.arr[s.length - 1], nil
}

func (s *Stack) Reverse() {
	s2 := NewStack()
	for s.length != 0 {
		n, _ := s.Pop()
		s2.Push(n)
	}
	s.arr = s2.arr
	s.length = s2.length
}

func PrintAns1(file *os.File) {
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
		from--
		to--
		
		for i := 0; i < moveCount; i++ {
			num, _ := stacks[from].Pop()
			stacks[to].Push(num)
		}
	}

	common.HandleError(sc.Err())
	fmt.Println("Day 5 Part 1 ans is:")
	for i := 0; i < len(stacks); i++ {
		num, _ := stacks[i].Peek()
		fmt.Print(string((byte)(num)))
	}
}