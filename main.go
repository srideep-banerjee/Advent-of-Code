package main

import (
	"example/day7"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if(err != nil) {
		panic(err)
	}
	
	day7.PrintAns2(file)

	err = file.Close()
	if(err != nil) {
		panic(err)
	}
}
