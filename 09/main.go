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

type BlockInfo struct {
	index  int
	length int
	fileId int
}

type Disk struct {
	disk   []int
	blocks []BlockInfo
}

func parse(input string) Disk {
	files := slices.Repeat([]int{-1}, (len(input)-1)*9)
	fileIndexes := make([]BlockInfo, (len(input)+1)/2)
	currentFileIndex := 0

	for i := 0; i < len(input); i++ {
		curr, _ := strconv.Atoi(string(input[i]))

		if i%2 != 0 {
			currentFileIndex += curr
			continue
		}

		fileIndexes[i/2] = BlockInfo{
			index:  currentFileIndex,
			length: curr,
			fileId: i / 2,
		}
		for range curr {
			files[currentFileIndex] = i / 2
			currentFileIndex++
		}
	}

	return Disk{
		disk:   files[:currentFileIndex],
		blocks: fileIndexes,
	}
}

func checkSum(disk []int) int {
	checkSum := 0
	for i, fileId := range disk {
		if fileId == -1 {
			continue
		}
		checkSum += i * fileId
	}

	return checkSum
}

func runStep1(input string) int {
	files := parse(input).disk

	leftIndex := 0
	rightIndex := len(files) - 1
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

	return checkSum(files)
}

func (d *Disk) findFreeSpaceWith(n, max int) int {
	freeSpaceIndex := -1

	for i := 0; i < max; i++ {
		if d.disk[i] != -1 {
			freeSpaceIndex = -1
			continue
		}

		if freeSpaceIndex == -1 {
			freeSpaceIndex = i
		}

		if i-freeSpaceIndex+1 >= n {
			return freeSpaceIndex
		}
	}

	return -1
}

func runStep2(input string) int {
	disk := parse(input)
	files := disk.disk

	for j := len(disk.blocks) - 1; j >= 0; j-- {
		currentFileSize := disk.blocks[j].length

		firstFreeSpace := disk.findFreeSpaceWith(currentFileSize, disk.blocks[j].index)
		if firstFreeSpace == -1 {
			continue
		}

		for n := range currentFileSize {
			files[firstFreeSpace+n] = disk.blocks[j].fileId
			files[disk.blocks[j].index+n] = -1
		}
	}

	return checkSum(files)
}
