package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseFile() [][]int {
	//file, err := os.Open("../data/test.txt")
	file, err := os.Open("../data/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	data := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		row := strings.Fields(line)
		numbers := []int{}

		for _, v := range row {
			number, _ := strconv.Atoi(v)
			numbers = append(numbers, number)
		}

		data = append(data, numbers)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func isEqual(data []int8) bool {
	counter := make(map[int8]int)

	for _, row := range data {
		counter[row]++
	}

	for _, count := range counter {
		if count != len(data) {
			return false
		}
	}

	return true
}

func isEqual2(data []int8) bool {
	counter := make(map[int8]int)

	for _, row := range data {
		counter[row]++
	}

	if len(counter) > 2 {
		return false
	}

	for _, count := range counter {
		if count != len(data) {
			if count == len(data)-1 || count == 1 {
				return true
			}

			return false
		}
	}

	return true
}

func diffRange(data []int) bool {
	for _, v := range data {
		if v > 3 {
			return false
		}
	}

	return true
}

func part1() {
	data := parseFile()
	count := 0

	for _, row := range data {
		var scale []int8
		var diff []int

		rowLength := len(row)
		for index, result := range row {
			if index == rowLength-1 {
				continue
			}

			searchIndex := index + 1

			if result == row[searchIndex] {
				scale = append(scale, 0)
				diff = append(diff, 0)
			} else if result > row[searchIndex] {
				scale = append(scale, 1)
				diff = append(diff, result-row[searchIndex])
			} else if result < row[searchIndex] {
				scale = append(scale, -1)
				diff = append(diff, row[searchIndex]-result)
			}
		}

		equal := isEqual(scale)
		lowRange := diffRange(diff)

		if equal && lowRange {
			count++
		}
	}

	fmt.Println("Safe Report Count:", count)
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func checkLevels(data []int, skipIndex int, dampen bool) bool {
	safe := true
	var checkData []int

	if skipIndex >= 0 {
		checkData = removeIndex(data, skipIndex)
	} else {
		checkData = data
	}

	increasing := checkData[0] < checkData[1]

	rowLength := len(checkData)
	for index := range rowLength - 1 {
		diff := checkData[index+1] - checkData[index]
		diff = int(math.Abs(float64(diff)))

		if diff > 3 || diff < 1 {
			safe = false
		} else if increasing && (checkData[index] > checkData[index+1]) {
			safe = false
		} else if !increasing && (checkData[index] < checkData[index+1]) {
			safe = false
		}

		if dampen && safe == false {
			skipIndex = index
			break
		}
	}

	if dampen && skipIndex >= 0 {
		safe = checkLevels(data, skipIndex, false) || checkLevels(data, skipIndex+1, false) || checkLevels(data, skipIndex-1, false)
	}

	return safe
}

func part2() {
	data := parseFile()
	count := 0

	for _, row := range data {
		safe := checkLevels(row, -1, true)

		if safe {
			count++
		}
	}

	fmt.Println("Safe Report Count:", count)
}

func main() {
	part1()
	part2()
}
