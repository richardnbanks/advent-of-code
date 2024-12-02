package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseList(filename string) [][]int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var list1 = []int{}
	var list2 = []int{}

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	var lists = [][]int{list1, list2}

	return lists
}

func part1() {
	//lists := parseList("../input/test.txt")
	lists := parseList("../input/input.txt")

	var distance = []int{}
	totalDistance := 0

	sort.Ints(lists[0])
	sort.Ints(lists[1])

	for index, number := range lists[0] {
		diff := 0

		if number > lists[1][index] {
			diff = number - lists[1][index]
		} else {
			diff = lists[1][index] - number
		}

		distance = append(distance, diff)
		totalDistance += diff
	}

	fmt.Println("Distance:", totalDistance)
}

func part2() {
	//lists := parseList("../input/test.txt")
	lists := parseList("../input/input.txt")

	similarity := 0

	for _, number := range lists[0] {
		count := 0

		for _, number2 := range lists[1] {
			if number == number2 {
				count++
			}
		}

		similarityScore := number * count

		similarity += similarityScore
	}

	fmt.Println("Similarity:", similarity)
}

func main() {
	part1()
	part2()
}
