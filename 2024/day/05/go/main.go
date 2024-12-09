package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func parseFile() ([][]string, [][]string) {
	//file, err := os.Open("../data/test.txt")
	file, err := os.Open("../data/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var rules [][]string
	var pages [][]string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")

			rules = append(rules, rule)
		}

		if strings.Contains(line, ",") {
			page := strings.Split(line, ",")

			pages = append(pages, page)
		}
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	return rules, pages
}

func checkRule(rule []string, list []string) bool {
	return slices.Contains(list[slices.Index(list, rule[0]):], rule[1]) && slices.Contains(list[:slices.Index(list, rule[1])], rule[0])
}

func part1() {
	rules, pages := parseFile()

	sum := 0

	for _, list := range pages {
		rulePass := true

		for _, rule := range rules {
			if slices.Contains(list, rule[0]) && slices.Contains(list, rule[1]) {
				if checkRule(rule, list) == false {
					rulePass = false
				}
			}
		}

		if rulePass {
			page, _ := strconv.Atoi(list[len(list)/2])

			sum += page
		}
	}

	fmt.Println("Part 1:", sum)
}

func reorder(rule []string, list []string) []string {
	swap := reflect.Swapper(list)
	swap(slices.Index(list, rule[1]), slices.Index(list, rule[0]))

	return list
}

func reprocessRules(rules [][]string, list []string) []string {
	for _, rule := range rules {
		if slices.Contains(list, rule[0]) && slices.Contains(list, rule[1]) {
			if checkRule(rule, list) == false {
				reorder(rule, list)
				reprocessRules(rules, list)
			}
		}
	}

	return list
}

func processRules(rules [][]string, list []string) bool {
	for _, rule := range rules {
		if slices.Contains(list, rule[0]) && slices.Contains(list, rule[1]) {
			if checkRule(rule, list) == false {
				return false
			}
		}
	}

	return true
}

func part2() {
	rules, pages := parseFile()
	var failed [][]string

	sum := 0

	for _, list := range pages {
		if processRules(rules, list) == false {
			failed = append(failed, list)
		}
	}

	for _, list := range failed {
		newList := reprocessRules(rules, list)

		page, _ := strconv.Atoi(newList[len(newList)/2])

		sum += page
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	part1()
	part2()
}
