package main

import (
	"example/2022/day8"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day8.PrintAns1(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
