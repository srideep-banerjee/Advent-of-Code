package main

import (
	"example/2023/day3"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day3.PrintAns1(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
