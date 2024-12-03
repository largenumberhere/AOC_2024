package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// read input
	file_contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(file_contents)

	// capture the numbers in mul(1,2) as seperate groups
	regex_str := `mul\((\d*),(\d*)\)`

	regex, err := regexp.Compile(regex_str)
	if err != nil {
		log.Fatal(err)
	}

	// result is an array of strings, each item like "mul(2,4)", "2", "4"
	results := regex.FindAllStringSubmatch(lines, -1)

	var total uint64 = 0
	for _, result := range results {
		// unpack the result
		var arg1 int
		var arg2 int
		var all string = result[0]

		arg1, err = strconv.Atoi(result[1])
		if err != nil {
			log.Fatal(err)
		}
		arg2, err = strconv.Atoi(result[2])
		if err != nil {
			log.Fatal(err)
		}

		sum := arg1 * arg2
		fmt.Println("evaluating expression ", all, " = ", sum)

		total += uint64(sum)
	}

	fmt.Println("total ", total)

}
