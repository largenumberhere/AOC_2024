package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	inputs []int
	output int
}

// chat gpt code :(
// generateString is a recursive function that generates all strings of length n
func generateString(n int, current string) []string {
	// Base case: if the current string length equals n, return it in a slice
	if len(current) == n {
		return []string{current}
	}

	// Recursively generate strings with '*' and '+'
	var result []string
	result = append(result, generateString(n, current+"*")...)
	result = append(result, generateString(n, current+"+")...)
	result = append(result, generateString(n, current+"|")...)

	return result
}

func generateStrings(n int) []string {
	return generateString(n, "")
}

// func cat(one string, two string) string {
// 	var b bytes.Buffer
// 	b.WriteString(one)
// 	b.WriteString(two)

// 	return b.String()
// }

func catInt(left int, right int) int {
	// left2 := strconv.Itoa(left)
	// right2 := strconv.Itoa(right)

	// all := cat(left2, right2)

	// out, err := strconv.Atoi(all)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// implementation based on https://stackoverflow.com/questions/12700497/how-to-concatenate-two-integers-in-c
	pow := 10
	for right >= pow {
		pow *= 10
	}

	return left*pow + right
}

func solveEquation(e Equation) bool {
	to_try := generateStrings(len(e.inputs))
	for _, line := range to_try {
		var sum int
		for i, symbol := range line {
			if symbol == '*' {
				sum *= e.inputs[i]
			} else if symbol == '+' {
				sum += e.inputs[i]
			} else if symbol == '|' {
				sum = catInt(sum, e.inputs[i])
			}
		}

		if sum == e.output {
			return true
		}
	}

	return false
}

func parseEquations(file_contents string) []Equation {
	lines := strings.Split(file_contents, "\n")

	var equations []Equation

	for _, line := range lines {
		// ignore empty lines
		if strings.Trim(line, " \r\n") == "" {
			continue
		}

		var inputs []int
		var output int

		// parse output
		left := strings.Trim(strings.Split(line, ":")[0], " \r\n")
		output, err := strconv.Atoi(left)
		if err != nil {
			log.Fatal(err)
		}

		// parse inputs
		for i, num := range strings.Split(line, " ") {
			if i == 0 {
				continue
			}

			num_trimmed := strings.Trim(num, " \r\n")
			if num_trimmed == "" {
				continue
			}

			num_int, err := strconv.Atoi(num_trimmed)
			if err != nil {
				log.Fatal(err)
			}

			inputs = append(inputs, num_int)
		}

		for _, in := range inputs {
			fmt.Println("input:", in)
		}
		fmt.Println("output:", output)

		equations = append(equations, Equation{
			inputs: inputs,
			output: output,
		})
	}

	return equations
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file_contents := strings.Trim(string(bytes), " \n")

	equations := parseEquations(file_contents)

	fmt.Println(equations)

	var tally int64
	for _, e := range equations {
		result := solveEquation(e)
		if result {
			tally += int64(e.output)
			fmt.Println("solvable :", e)
		} else {
			fmt.Println("not solveable", e)
		}
	}

	fmt.Println("tally: ", tally)
}
