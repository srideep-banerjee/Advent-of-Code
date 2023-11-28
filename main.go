package main

import (
	"example/day4"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day4.PrintAns1(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
