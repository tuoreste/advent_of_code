package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafe(data []int) bool {
	isIncreasing := false
	isDecreasing := false
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			isIncreasing = true
		} else if data[i] < data[i-1] {
			isDecreasing = true
		}

		if AbsInt(data[i]-data[i-1]) < 1 || AbsInt(data[i]-data[i-1]) > 3 {
			return false
		}
	}
	if isIncreasing && isDecreasing {
		return false
	}
	return true
}

func checkPoint(data []int) bool {
	if isSafe(data) {
		return true
	}

	for i := 0; i < len(data); i++ {
		var temp []int
		temp = append(temp, data[:i]...)
		temp = append(temp, data[i+1:]...)

		if isSafe(temp) {
			return true
		}
	}
	return false
}

func safetyChecker() int {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		var data []int
		line := scanner.Text()
		columns := strings.Fields(line)

		for _, value := range columns {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error parsing a figure:", err)
				return 0
			}
			data = append(data, num)
		}

		if checkPoint(data) {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	return counter
}

func main() {
	fmt.Println(safetyChecker())
}
