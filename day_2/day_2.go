package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isStepSizeAllowed(list []int) bool {
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

func isDecendingOrAscending(list []int) bool {
	prev := list[0]

	var ascending int
	var descending int
	var stagnating int

	for i, item := range list {
		// skip first item
		if i == 0 {
			continue
		}

		if item-prev < 0 {
			ascending += 1
		} else if item-prev > 0 {
			descending += 1
		} else {
			stagnating += 1
		}

		prev = item
	}

	is_ascending_or_descending := ascending == 0 || descending == 0
	is_stagnating := stagnating > 0

	if is_stagnating {
		return false
	}

	if is_ascending_or_descending {
		return true
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
		is_changing := isDecendingOrAscending(report)
		is_valid_steps := isStepSizeAllowed(report)

		if !is_changing || !is_valid_steps {
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
