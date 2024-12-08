package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("Error to open the file", err)
		return
	}
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		x, errorX := strconv.Atoi(columns[0])
		y, errorY := strconv.Atoi(columns[1])

		if errorX != nil || errorY != nil {
			fmt.Println("Error parsing a figures:", errorX, errorY)
			return
		}
		
		col1 = append(col1, x)
		col2 = append(col2, y)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}
	fmt.Println("column1", col1)
	fmt.Println("column2", col2)
}
