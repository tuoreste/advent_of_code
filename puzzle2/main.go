package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func multiply(x int, y int) int {
	return (x * y)
}

func sum(nums [] int) int {
	total := 0
	for _,num := range(nums) {
		total += num;
	}
	return total
}

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

func findScore() int {
	var valueScore []int
	col1, col2 := dataInSlices();
	for _,num := range(col1) {
		counter := 0;
		for _,value := range(col2) {
			if (num == value) {
				counter++;
			}
		}
		valueScore = append(valueScore, multiply(num, counter))
	}
	return (sum(valueScore))
}

func main() {
	fmt.Println(findScore())
}
