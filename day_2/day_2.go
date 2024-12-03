package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func testFloorStepSize(list []int) bool {
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

func firstStagnating(list []int) int {
	prev := list[0]

	for i, item := range list {
		if i == 0 {
			continue
		}

		difference := item - prev

		if difference == 0 {
			return i
		}

		prev = item
	}

	return -1
}

func firstDecreasing(list []int) int {
	prev := list[0]

	for i, item := range list {
		if i == 0 {
			continue
		}

		difference := item - prev

		if difference < 0 {
			return i
		}

		prev = item
	}

	return -1
}

func firstIncreasing(list []int) int {
	prev := list[0]

	for i, item := range list {
		if i == 0 {
			continue
		}
		difference := item - prev

		if difference > 0 {
			return i
		}

		prev = item
	}

	return -1
}

// func testFloorsAgain(list []int) bool {
// 	// find the offender
// 	for i := 0; i < len(list); i++ {

// 	}
// }

func testFloorDirection(list []int) bool {
	ascending_count := countIncreasing(list)
	descending_count := countDecreasing(list)
	stagnating_count := countStagnating(list)

	is_ascending_or_descending := (ascending_count == 0 || descending_count == 0)
	is_stagnating := stagnating_count > 0

	if !is_stagnating && is_ascending_or_descending {
		return true
	}

	return false
}

func isReportValid(list []int) bool {
	if !testFloorDirection(list) || !testFloorStepSize(list) {
		return false
	}

	return true
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// destructively modifies the input and returns the result
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func floorsAreAllowed(list []int) bool {
	if isReportValid(list) {
		return true
	}

	// retry with a floor removed
	for i := 0; i < len(list); i++ {
		list2 := make([]int, len(list))
		copy(list2, list)

		list2 = remove(list2, i)

		if isReportValid(list2) {
			// fmt.Println("2nd try was safe ", list)
			return true
		} else {
			// fmt.Println("2nd try was unsafe ", list, " -> ", list2)
		}
	}

	return false
}

func parse_reports(file string) [][]int {
	each_line := strings.Split(file, "\n")

	// parse the lines of 'reports'
	var all_reports [][]int

	for _, line := range each_line {
		fmt.Println(line)

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

			fmt.Println(number)

			report_summary = append(report_summary, number)
		}

		if len(report_summary) == 0 {
			break
		}
		all_reports = append(all_reports, report_summary)

	}

	fmt.Println("reports: ", all_reports)

	return all_reports
}

func main() {
	file_contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(file_contents)

	// process the reports
	all_reports := parse_reports(lines)

	safe_count := 0

	for _, report := range all_reports {
		report_safe := floorsAreAllowed(report)

		if !report_safe {
			fmt.Println("report is unsafe ", report)
		} else {
			fmt.Println("report is safe ", report)
			safe_count += 1
		}
	}

	fmt.Println("safe count: ", safe_count)
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}

	return value
}
