package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isFloorStepSizeAllowed(list []int) bool {
	prev := list[0]

	for i, item := range list {
		// skip first item
		if i == 0 {
			continue
		}

		diff := abs(prev - item)
		if diff > 3 || diff < 1 {
			return false
		}

		prev = item
	}

	return true
}

func countIncreasing(list []int) int {
	prev := list[0]

	ascending_count := 0

	for i, item := range list {
		if i == 0 {
			continue
		}
		difference := item - prev

		if difference > 0 {
			ascending_count += 1
		}

		prev = item
	}

	return ascending_count
}

func countDecreasing(list []int) int {
	prev := list[0]

	descending_count := 0

	for i, item := range list {
		if i == 0 {
			continue
		}

		difference := item - prev

		if difference < 0 {
			descending_count += 1
		}

		prev = item
	}

	return descending_count
}

func countStagnating(list []int) int {
	prev := list[0]

	stagnating_count := 0

	for i, item := range list {
		if i == 0 {
			continue
		}

		difference := item - prev

		if difference == 0 {
			stagnating_count += 1
		}

		prev = item
	}

	return stagnating_count
}

// check if the floors are only decreasing or increasing
func testFloorDirection(list []int) bool {
	ascending_count := countIncreasing(list)
	descending_count := countDecreasing(list)
	stagnating_count := countStagnating(list)

	// does not handle list size of zero correctly but that didn't seem to matter
	is_ascending_or_descending := (ascending_count == 0 || descending_count == 0)
	is_stagnating := stagnating_count > 0

	if !is_stagnating && is_ascending_or_descending {
		return true
	}

	return false
}

// validate a single report
func isReportValid(list []int) bool {
	if !testFloorDirection(list) || !isFloorStepSizeAllowed(list) {
		return false
	}

	return true
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// destructively modifies the input and returns a new array as a result
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// determine if report has a combinartion of reports that is considered 'safe'
func isFloorsSafe(list []int, dampener_enabled bool) bool {
	if isReportValid(list) {
		return true
	}

	if dampener_enabled {
		// brute force each combination with one less floor
		for i := 0; i < len(list); i++ {
			list2 := make([]int, len(list))
			copy(list2, list)

			list2 = remove(list2, i)

			if isReportValid(list2) {
				return true
			}
		}
	}

	return false
}

// turn the inputs into an array of reports. Each report contains several values
func parse_reports(file string) [][]int {
	each_line := strings.Split(file, "\n")

	// parse the lines of 'reports'
	var all_reports [][]int

	for _, line := range each_line {
		var report_summary []int

		reports := strings.Split(line, " ")
		for _, report := range reports {
			report := strings.Trim(report, "\r\n ")

			if len(report) == 0 {
				break
			}
			number, err := strconv.Atoi(report)
			if err != nil {
				log.Fatal(err)
			}

			report_summary = append(report_summary, number)
		}

		if len(report_summary) == 0 {
			break
		}
		all_reports = append(all_reports, report_summary)

	}

	return all_reports
}

func main() {
	is_part_2 := true

	// read input
	file_contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(file_contents)

	// process the reports
	all_reports := parse_reports(lines)

	// calculate result
	safe_count := 0

	for _, report := range all_reports {
		report_safe := isFloorsSafe(report, is_part_2)

		if report_safe {
			safe_count += 1
		}
	}

	fmt.Println("safe count: ", safe_count)
}

// the 'absolute magnitude' of the input
func abs(value int) int {
	if value < 0 {
		value = -value
	}

	return value
}
