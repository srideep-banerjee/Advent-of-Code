package day3

import (
	"bufio"
	"fmt"
	"os"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)
	sum := 0
	lineNo := 0
	var commonRunes map[byte]bool = make(map[byte]bool)
	
	for sc.Scan() {
		var line string = sc.Text()
		
		if lineNo == 0 {
			for _, by := range []byte(line) {
				commonRunes[by] = true
			}
		} else {
			currentLineRunes := make(map[byte]bool)
			for _, by := range []byte(line) {
				currentLineRunes[by] = true
			}
			
			for by := range commonRunes {
				if !currentLineRunes[by] {
					delete(commonRunes, by)
				}
			}
		}
		
		lineNo++
		lineNo %= 3

		if(lineNo == 0) {
			for k := range commonRunes {
				if(k>='A' && k<='Z') {
					sum += (int) (k - 'A') + 27
				} else {
					sum += (int) (k - 'a') + 1
				}
				delete(commonRunes, k)
			}
		}
	}
	
	HandleError(sc.Err())
	
	fmt.Println("Day 3 Part 2 ans is",sum)
}