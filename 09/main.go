package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
)

//go:embed input.txt
var s string

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s)
	} else {
		ans = runStep2(s)
	}

	fmt.Println("Output: ", ans)
}

func runStep1(input string) int {
	files := slices.Repeat([]int{-1}, (len(input)-1)*9)
	currentFileIndex := int64(0)

	for i := 0; i < len(input); i++ {
		curr, _ := strconv.ParseInt(string(input[i]), 10, 32)

		if i%2 != 0 {
			currentFileIndex += curr
			continue
		}

		for range curr {
			files[currentFileIndex] = i / 2
			currentFileIndex++
		}
	}

	leftIndex := int64(0)
	rightIndex := currentFileIndex - 1
	for {
		if leftIndex >= rightIndex {
			break
		}

		if files[leftIndex] != -1 {
			leftIndex++
			continue
		}

		if files[rightIndex] == -1 {
			rightIndex--
			continue
		}

		files[leftIndex] = files[rightIndex]
		files[rightIndex] = -1

		leftIndex++
		rightIndex--
	}

	checkSum := 0
	for i, fileId := range files {
		if fileId == -1 {
			break
		}

		checkSum += i * fileId
	}

	return checkSum
}

func runStep2(input string) int {
	return 0
}
