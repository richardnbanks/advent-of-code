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

func part1() {
	//file, err := os.Open("test1.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		addGame := true

		row := strings.Split(line, ": ")

		pattern := regexp.MustCompile(`Game (?P<number>[0-9]+)`)
		game := pattern.FindStringSubmatch(row[0])
		gameNumber, _ := strconv.Atoi(game[1])

		sets := strings.Split(row[1], "; ")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cubeCount := range cubes {
				cubeInfo := strings.Split(cubeCount, " ")
				cubeNumber, _ := strconv.Atoi(cubeInfo[0])

				if cubeInfo[1] == "red" && cubeNumber > 12 {
					addGame = false
				}

				if cubeInfo[1] == "green" && cubeNumber > 13 {
					addGame = false
				}

				if cubeInfo[1] == "blue" && cubeNumber > 14 {
					addGame = false
				}
			}
		}

		if addGame {
			sum += gameNumber
		}
	}

	fmt.Printf("Answer 1: %d\n", sum)
}

func part2() {
	//file, err := os.Open("test2.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		row := strings.Split(line, ": ")

		red := 0
		green := 0
		blue := 0

		sets := strings.Split(row[1], "; ")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cubeCount := range cubes {
				cubeInfo := strings.Split(cubeCount, " ")
				cubeNumber, _ := strconv.Atoi(cubeInfo[0])

				if cubeInfo[1] == "red" && cubeNumber > red {
					red = cubeNumber
				}

				if cubeInfo[1] == "green" && cubeNumber > green {
					green = cubeNumber
				}

				if cubeInfo[1] == "blue" && cubeNumber > blue {
					blue = cubeNumber
				}
			}
		}

		power := red * green * blue
		sum += power
	}

	fmt.Printf("Answer 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
