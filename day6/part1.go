package day6

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
)

func PrintAns1(file *os.File) {
	var sc = bufio.NewScanner(file)

	res := -1
	
	sc.Scan()
	var line string = sc.Text()

		var freq = make(map[uint16]uint16)
		for i := uint16(0); i < uint16(len(line)); i++ {
			freq[uint16(line[i])]++
			
			if i >= 4 {
				if freq[uint16(line[i - 4])] == 1 {
					delete(freq, uint16(line[i - 4]))
				} else {
					freq[uint16(line[i - 4])]--
				}
			}
			
			if(len(freq) == 4) {
				res = int(i) + 1
				break
			}
		}

	common.HandleError(sc.Err())
	fmt.Println("Day 6 Part 1 ans is ", res)
}
