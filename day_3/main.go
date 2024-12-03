package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evlauateMul(regex_result []string) int {
	// unpack the result
	var arg1 int
	var arg2 int
	var err error

	arg1, err = strconv.Atoi(regex_result[1])
	if err != nil {
		log.Fatal(err)
	}
	arg2, err = strconv.Atoi(regex_result[2])
	if err != nil {
		log.Fatal(err)
	}

	sum := arg1 * arg2

	return sum
}

func main() {
	// read input
	file_contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(file_contents)

	// capture the numbers in mul(1,2) as seperate groups
	// regex_str := `mul\((\d*),(\d*)\)`

	/**
		This regex captures each occourence of:
			- mul(a,b) with a and b being integers
			- don't()
			- do()
	**/
	regex_str := `mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`

	regex, err := regexp.Compile(regex_str)
	if err != nil {
		log.Fatal(err)
	}

	// result is an array of strings, each item like "mul(2,4)", "2", "4"
	results := regex.FindAllStringSubmatch(lines, -1)

	// var flag bool = true
	var total uint64 = 0
	var ignore_mul bool = false
	for _, result := range results {
		if strings.HasPrefix(result[0], "mul") {
			if !ignore_mul {
				total += uint64(evlauateMul(result))
				fmt.Println("added mul: ", result[0])
			} else {
				fmt.Println("skipped mul: ", result[0])
			}
		} else if strings.HasPrefix(result[0], "don't") {
			ignore_mul = true
			fmt.Println("ignoring muls")
		} else if strings.HasPrefix(result[0], "do") {
			ignore_mul = false
			fmt.Println("parsing muls")
		} else {
			fmt.Println("unexpected symbol in", result)
			continue
		}

	}

	fmt.Println("total ", total)

}
