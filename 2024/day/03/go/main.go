package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseFile() []string {
	//file, err := os.Open("../data/test.txt")
	//file, err := os.Open("../data/test2.txt")
	file, err := os.Open("../data/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, line)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func part1() {
	data := parseFile()
	sum := 0

	for _, line := range data {

		re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))`)
		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			match = strings.Replace(match, "mul(", "", 1)
			match = strings.Replace(match, ")", "", 1)

			numbers := strings.Split(match, ",")

			first, _ := strconv.Atoi(numbers[0])
			second, _ := strconv.Atoi(numbers[1])

			multiplier := first * second

			sum += multiplier
		}
	}

	fmt.Println("Part 1: Sum of multiplications:", sum)
}

func part2() {
	data := parseFile()
	sum := 0

	skip := false
	for _, line := range data {
		re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			if strings.Contains(match, "do") {
				if match == "do()" {
					skip = false
				} else {
					skip = true
				}

				continue
			}

			if skip {
				continue
			}

			match = strings.Replace(match, "mul(", "", 1)
			match = strings.Replace(match, ")", "", 1)

			numbers := strings.Split(match, ",")

			first, _ := strconv.Atoi(numbers[0])
			second, _ := strconv.Atoi(numbers[1])

			multiplier := first * second

			sum += multiplier
		}
	}

	fmt.Println("Part 2: Sum of multiplications:", sum)
}

func main() {
	part1()
	part2()
}
