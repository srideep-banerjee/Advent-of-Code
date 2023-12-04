package day3

import (
	"bufio"
	"example/common"
	"fmt"
	"os"
)

func PrintAns2(file *os.File) {
	var sc = bufio.NewScanner(file)

	input := [][]byte{}
	sum := 0
	for sc.Scan() {
		var line = sc.Text()
		input = append(input, []byte(line))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '*' {
				sum += getGearRatio(j,i,len(input[0]), len(input), input)
			}
		}
	}

	common.HandleError(sc.Err())

	fmt.Println("Day 2 Part 1 ans is", sum)
}

func getGearRatio(x, y, width, height int, mat [][]byte) int {
	adjNums := []int{}
	
	if x != 0 {
		if isNumber(mat[y][x-1]) {
			adjNums = append(adjNums, extractNumber(x-1, y, width, mat))
		}
		if y != 0 && isNumber(mat[y-1][x-1]) {
			adjNums = append(adjNums, extractNumber(x-1, y-1, width, mat))
		}
		if y != height-1 && isNumber(mat[y+1][x-1]) {
			adjNums = append(adjNums, extractNumber(x-1, y+1, width, mat))
		}
	}
	if x != width-1 {
		if isNumber(mat[y][x+1]) {
			adjNums = append(adjNums, extractNumber(x+1, y, width, mat))
		}
		if y != 0 && isNumber(mat[y-1][x+1]) {
			adjNums = append(adjNums, extractNumber(x+1, y-1, width, mat))
		}
		if y != height-1 && isNumber(mat[y+1][x+1]) {
			adjNums = append(adjNums, extractNumber(x+1, y+1, width, mat))
		}
	}
	if y != 0 && isNumber(mat[y-1][x]) {
		adjNums = append(adjNums, extractNumber(x, y-1, width, mat))
	}
	if y != height-1 && isNumber(mat[y+1][x]) {
		adjNums = append(adjNums, extractNumber(x, y+1, width, mat))
	}
	
	if len(adjNums) != 2 {return 0}
	return adjNums[0] * adjNums[1]
}

// func isSymbol(b byte) bool {
// 	return !((b >= '0' && b <= '9') || b == '.')
// }

// func isNumber(b byte) bool {
// 	return b >= '0' && b <= '9'
// }

// func extractNumber(x, y, width int, mat [][]byte) int {
// 	num := int(mat[y][x] - '0')
// 	for j := x - 1; j >= 0 && isNumber(mat[y][j]); j-- {
// 		num += int(math.Pow10(x-j))*int(mat[y][j]-'0')
// 		mat[y][j] = '.'
// 	}
// 	for j := x + 1; j < width && isNumber(mat[y][j]); j++ {
// 		num *= 10
// 		num += int(mat[y][j] - '0')
// 		mat[y][j] = '.'
// 	}
// 	return num
// }

// func extractAndSumNearbyNumbers(x, y, width, height int, mat [][]byte) int {
// 	sum := 0
// 	if x != 0 {
// 		if isNumber(mat[y][x-1]) {
// 			sum += extractNumber(x-1, y, width, mat)
// 		}
// 		if y != 0 && isNumber(mat[y-1][x-1]) {
// 			sum += extractNumber(x-1, y-1, width, mat)
// 		}
// 		if y != height-1 && isNumber(mat[y+1][x-1]) {
// 			sum += extractNumber(x-1, y+1, width, mat)
// 		}
// 	}
// 	if x != width-1 {
// 		if isNumber(mat[y][x+1]) {
// 			sum += extractNumber(x+1, y, width, mat)
// 		}
// 		if y != 0 && isNumber(mat[y-1][x+1]) {
// 			sum += extractNumber(x+1, y-1, width, mat)
// 		}
// 		if y != height-1 && isNumber(mat[y+1][x+1]) {
// 			sum += extractNumber(x+1, y+1, width, mat)
// 		}
// 	}
// 	if y != 0 && isNumber(mat[y-1][x]) {
// 		sum += extractNumber(x, y-1, width, mat)
// 	}
// 	if y != height-1 && isNumber(mat[y+1][x]) {
// 		sum += extractNumber(x, y+1, width, mat)
// 	}
	
// 	return sum
// }
