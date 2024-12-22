package main

import (
	"fmt"
	"log"
	"strconv"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type Num int64

func parseLines(filepath string) ([]int, error) {
	var lines []int = nil

	strings, err := aoc_lib.GrabLines(filepath)
	if err != nil {
		return lines, err
	}
	count := len(strings)

	lines = make([]int, 0, count)
	for _, v := range strings {
		number, err := strconv.Atoi(v)
		if err != nil {
			return lines, err
		}

		lines = append(lines, number)
	}

	return lines, nil
}

func mix(number_one Num, number_two Num) Num {
	return number_one ^ number_two

}

func prune(number_one Num) Num {
	return number_one % 16777216
}

func nextSecret(secret Num) Num {
	secret = mix(secret, secret*64)
	secret = prune(secret)

	secret = mix(secret, Num(secret/32))
	secret = prune(secret)

	secret = mix(secret, secret*2048)
	secret = prune(secret)

	return secret
}

func secret2K(secret Num) Num {
	for i := 0; i < 2_000; i++ {
		secret = nextSecret(secret)
	}

	return secret
}

func sumSecrets2k(secrets []int) Num {
	sum := Num(0)
	for _, line := range secrets {
		thousand := secret2K(Num(line))
		sum += thousand
	}

	return sum
}

func main() {
	lines, err := parseLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	secrets_sum := sumSecrets2k(lines)
	fmt.Println("secrets_sum :", secrets_sum)

}
