package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func parseFile() ([]string, int, int) {
	//file, err := os.Open("../data/test.txt")
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

	width := len(data[0])
	height := len(data)

	return data, width, height
}

func searchHorizontally(data []string, search string) int {
	re := regexp.MustCompile(`(` + search + `)`)
	count := 0

	for _, line := range data {
		matches := re.FindAllString(line, -1)

		count += len(matches)
	}

	return count
}

func searchVertically(data []string, width int, height int, search string) int {
	count := 0
	queryLength := len(search)

	stop := height - queryLength + 1

	for row := range stop {
		for column := range width {
			var potentialWord string

			for letter := range queryLength {
				potentialWord += string([]rune(data[row+letter])[column])
			}

			if potentialWord == search {
				count++
			}
		}
	}

	return count
}

func searchDiagonally(data []string, width int, height int, direction bool, search string) int {
	count := 0
	queryLength := len(search)

	stop := height - queryLength + 1

	for row := range stop {
		for column := range width {
			var potentialWord string

			for letter := range queryLength {
				queryColumn := 0

				if direction {
					queryColumn = column + queryLength - letter - 1
				} else {
					queryColumn = column + letter
				}

				if queryColumn > width-1 {
					continue
				}

				potentialWord += string([]rune(data[row+letter])[queryColumn])
			}

			if potentialWord == search {
				count++
			}
		}
	}

	return count
}

func searchXMas(data []string, width int, height int) int {
	count := 0
	queryLength := 3

	stop := height - queryLength + 1

	for row := range stop {
		for column := range width {
			var potentialWord1 string
			var potentialWord2 string

			for letter := range queryLength {
				queryColumnLeft := column + letter
				queryColumnRight := column + queryLength - letter - 1

				if queryColumnLeft > width-1 || queryColumnRight > width-1 {
					continue
				}

				potentialWord1 += string([]rune(data[row+letter])[queryColumnLeft])
				potentialWord2 += string([]rune(data[row+letter])[queryColumnRight])
			}

			if (potentialWord1 == "MAS" || potentialWord1 == "SAM") && (potentialWord2 == "MAS" || potentialWord2 == "SAM") {
				count++
			}
		}
	}

	return count
}

func part1() {
	data, width, height := parseFile()

	count := 0
	for _, search := range []string{"XMAS", "SAMX"} {
		count += searchHorizontally(data, search)
		count += searchVertically(data, width, height, search)
		count += searchDiagonally(data, width, height, false, search)
		count += searchDiagonally(data, width, height, true, search)
	}

	fmt.Println("Part 1: Total matches:", count)
}

func part2() {
	data, width, height := parseFile()

	count := searchXMas(data, width, height)

	fmt.Println("Part 2: Total matches:", count)
}

func main() {
	part1()
	part2()
}
