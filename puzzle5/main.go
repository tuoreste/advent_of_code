package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func processMul(input string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	var results []int

	for _, match := range matches {
		if len(match) == 3 {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				result := x * y
				results = append(results, result)
			}
		}
	}
	return sum(results)
}

func parseFile() bool {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalSum += processMul(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	fmt.Println(totalSum)
	return true
}

func main() {
	parseFile()
}
