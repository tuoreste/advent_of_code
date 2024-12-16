package main

import (
	"fmt"
	"bufio"
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

func	checkPoint(data []int) (bool){
	isIncreasing := 0
	isDecreasing := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i - 1] {
			isIncreasing++
		} else if data[i] < data[i - 1] {
			isDecreasing++
		}
		if AbsInt(data[i] - data[i - 1]) < 1 || AbsInt(data[i] - data[i - 1]) > 3 {
			return false
		}
	}
	if (isIncreasing > 0 && isDecreasing > 0) {
		return false
	}
	return true
}

func safetyChecker() int {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return 0 // Return 0 if file cannot be opened
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

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