package main

import (
	"example/day1"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day1.PrintAns2(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
