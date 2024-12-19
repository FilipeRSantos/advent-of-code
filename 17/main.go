package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string
var a, b, c int
var out chan int
var numbers []int
var pointer int
var outputIndex int

const (
	Corrupted = -1
	Success   = 0
	Halted    = 1
	Jump      = 2
)

func main() {
	var ans string
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s)
	} else {
		ans = strconv.Itoa(runStep2(s))
	}

	fmt.Println("Output: ", ans)
}

func getComboOperator(operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}

	panic("Should not be here")
}

func adv(operand int) int {
	return a / (maths.Pow(2, getComboOperator(operand)))
}

func execute(part1 bool) int {
	if pointer >= len(numbers)-1 {
		if part1 {
			close(out)
		}
		return Halted
	}

	instruction := numbers[pointer]
	operand := numbers[pointer+1]

	switch instruction {
	case 0:
		a = adv(operand)
	case 1:
		b = b ^ operand
	case 2:
		v := getComboOperator(operand)
		b = v % 8
	case 3:
		if a == 0 {
			return Success
		}

		pointer = operand
		return Jump
	case 4:
		b = b ^ c
	case 5:
		v := getComboOperator(operand) % 8

		if part1 {
			out <- v
		} else {
			if numbers[outputIndex] != v {
				return Corrupted
			}
			outputIndex++
		}

	case 6:
		b = adv(operand)
	case 7:
		c = adv(operand)
	}

	return Success
}

func process(part1 bool) bool {
	for {
		status := execute(part1)
		switch status {
		case Success:
			pointer += 2
		case Halted:
			return outputIndex == len(numbers)
		case Jump:
			continue
		case Corrupted:
			return false
		}
	}
}

func runStep1(input string) string {
	out = make(chan int)
	pointer = 0
	parse(input)

	go process(true)

	separator := ""
	var buffer bytes.Buffer
	for v := range out {
		buffer.WriteString(separator)
		buffer.WriteString(strconv.Itoa(v))
		separator = ","
	}

	return buffer.String()
}

func runStep2(input string) int {
	parse(input)

	initialA := 0
	originalB := b
	originalC := c

	logProgress := func() {
		for range time.Tick(time.Second * 5) {
			fmt.Println(initialA)
		}
	}

	go logProgress()

	for {
		a = initialA
		b = originalB
		c = originalC
		pointer = 0
		outputIndex = 0

		if process(false) {
			break
		}

		initialA++
	}

	return initialA
}

func parse(input string) {
	lines := strings.Split(input, "\n")

	a = maths.ParseInt(strings.Split(lines[0], ": ")[1])
	b = maths.ParseInt(strings.Split(lines[1], ": ")[1])
	c = maths.ParseInt(strings.Split(lines[2], ": ")[1])

	operations := strings.Split(strings.Split(lines[4], ": ")[1], ",")

	numbers = make([]int, len(operations))
	for i := range len(operations) {
		numbers[i] = maths.ParseInt(operations[i])
	}
}
