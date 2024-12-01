package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func get_rhs(file_contents string) []int {
	regex_rhs, regex_err := regexp.Compile("\\s([0-9]*)\n")
	if regex_err != nil {
		log.Fatal(regex_err)
	}

	result_rhs := regex_rhs.FindAllString(file_contents, -1)
	if result_rhs == nil {
		log.Fatal(regex_rhs)
	}

	var rhs_sorted []int

	for _, item := range result_rhs {
		stripped := strings.Trim(item, " \n\t ")

		if len(stripped) == 0 {
			continue
		}

		number, number_err := strconv.Atoi(stripped)
		if number_err != nil {
			panic(number_err)
		}

		rhs_sorted = append(rhs_sorted, number)
	}

	slices.Sort(rhs_sorted)

	return rhs_sorted
}

func get_lhs(file_contents string) []int {
	regex_rhs, regex_err := regexp.Compile("([0-9]*)\\s\\s(?:([0-9]*))")
	if regex_err != nil {
		log.Fatal(regex_err)
	}

	result_rhs := regex_rhs.FindAllString(file_contents, -1)
	if result_rhs == nil {
		log.Fatal(regex_rhs)
	}

	var rhs_sorted []int

	for _, item := range result_rhs {
		stripped := strings.Trim(item, " \n\t ")

		if len(stripped) == 0 {
			continue
		}

		number, number_err := strconv.Atoi(stripped)
		if number_err != nil {
			log.Fatal(number_err)
		}

		rhs_sorted = append(rhs_sorted, number)
	}

	slices.Sort(rhs_sorted)

	return rhs_sorted
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file_contents := string(data)

	rhs_array := get_rhs(file_contents)

	lhs_array := get_lhs(file_contents)

	fmt.Println(len(lhs_array), len(rhs_array))

	if len(lhs_array) != len(rhs_array) {
		panic("uneven slices")
	}

	var total uint64 = 0

	for i := 0; i < len(lhs_array); i++ {
		difference := rhs_array[i] - lhs_array[i]
		if difference < 0 {
			difference = -difference
		}

		total = total + uint64(difference)
	}

	fmt.Print(total)
}
