package main

import (
	"example/2023/day5"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day5.PrintAns2(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
