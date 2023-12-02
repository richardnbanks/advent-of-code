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

func getNumber(s string) int {
	pattern := regexp.MustCompile(`[0-9]`)
	numbers := pattern.FindAllString(s, len(s))

	firstNumber := numbers[0]
	lastNumber := numbers[0]

	if len(numbers) > 1 {
		lastNumber = numbers[len(numbers)-1]
	}

	numberString := firstNumber + lastNumber

	number, _ := strconv.Atoi(numberString)

	return number
}

func part1() {
	//file, err := os.Open("test1.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		sum += getNumber(line)
	}

	err = file.Close()

	if err != nil {
		return
	}

	fmt.Printf("Answer 1: %d\n", sum)
}

func part2() {
	//file, err := os.Open("test2.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	words := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	for scanner.Scan() {
		line := scanner.Text()

		firstPosition := -1
		firstNumber := 0

		lastPosition := -1
		lastNumber := 0

		for number, word := range words {
			firstIndexWord := strings.Index(line, word)
			firstIndexNumber := strings.Index(line, strconv.Itoa(number))

			if firstIndexNumber >= 0 {
				if firstPosition < 0 {
					firstPosition = firstIndexNumber
					firstNumber = number
				}

				if firstIndexNumber < firstPosition {
					firstPosition = firstIndexNumber
					firstNumber = number
				}
			}

			if firstIndexWord >= 0 {
				if firstPosition < 0 {
					firstPosition = firstIndexWord
					firstNumber = number
				}

				if firstPosition < 0 && firstIndexNumber < 0 {
					firstPosition = firstIndexWord
					firstNumber = number
				}

				if firstIndexWord < firstPosition {
					firstPosition = firstIndexWord
					firstNumber = number
				}
			}

			lastIndexWord := strings.LastIndex(line, word)
			lastIndexNumber := strings.LastIndex(line, strconv.Itoa(number))

			if lastIndexNumber >= 0 {
				if lastPosition < 0 {
					lastPosition = lastIndexNumber
					lastNumber = number
				}

				if lastIndexNumber > lastPosition {
					lastPosition = lastIndexNumber
					lastNumber = number
				}
			}

			if lastIndexWord >= 0 {
				if lastPosition < 0 {
					lastPosition = lastIndexWord
					lastNumber = number
				}

				if lastPosition < 0 && lastIndexNumber < 0 {
					lastPosition = lastIndexWord
					lastNumber = number
				}

				if lastIndexWord > lastPosition {
					lastPosition = lastIndexWord
					lastNumber = number
				}
			}
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, lastNumber))

		sum += number
	}

	err = file.Close()

	if err != nil {
		return
	}

	fmt.Printf("Answer 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
