package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findUseRemoveCurrentMin(nums *[]int) int {
	if len(*nums) == 0 {
		fmt.Println("Error: slice is empty")
		return 0
	}

	minIndex := 0
	for i, num := range *nums {
		if num < (*nums)[minIndex] {
			minIndex = i
		}
	}
	minValue := (*nums)[minIndex]
	*nums = append((*nums)[:minIndex], (*nums)[minIndex + 1:]...)
	return minValue
}

// func	findMin(nums []int) int {
// 	if len(nums) == 0 {
// 		fmt.Println("Empty slice")
// 		return 0
// 	}

// 	min := nums[0]
// 	for _, num := range nums[1:]{
// 		if num < min {
// 			min = num
// 		}
// 	}

// 	return min
// }

func dataInSlices() ([]int, []int) {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("Error to open the file", err)
		return nil, nil
	}
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) != 2 {
			fmt.Println("Unexpected number of columns(only 2)")
			return nil, nil
		}
		x, errorX := strconv.Atoi(columns[0])
		y, errorY := strconv.Atoi(columns[1])

		if errorX != nil || errorY != nil {
			fmt.Println("Error parsing a figures:", errorX, errorY)
			return nil, nil
		}
		
		col1 = append(col1, x)
		col2 = append(col2, y)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return nil, nil
	}
	return col1, col2
}

func main() {
	var sum = 0
	col1, col2 := dataInSlices()
	for len(col1) > 0 && len(col2) > 0 {
		minA := findUseRemoveCurrentMin(&col1)
		minB := findUseRemoveCurrentMin(&col2)
		if minA < minB {
			sum += minB - minA
		} else {
			sum += minA - minB
		}
	}

	// fmt.Println("column1", col1)
	// fmt.Println("column2", col2)

	// minA := findMin(col1)
	// minB := findMin(col2)

	// fmt.Println("Minimum number", minA)
	// fmt.Println("Minimum number", minB)
	
	fmt.Println("SUM OF DIFFERENCES", sum)
}
