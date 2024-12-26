package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib"
)

func removeLeadingZeros(integer *string) {
	zeros := 0
	for i := 0; i < len(*integer); i++ {
		if (*integer)[i] != '0' {
			break
		} else {
			zeros += 1
		}
	}

	if zeros > 0 {
		*integer = (*integer)[zeros:]
	}

	if *integer == "" {
		*integer = "0"
	}
}

func updateStone(stone *string, out1 *string, out2 *string) bool {
	if *stone == "0" {
		*out1 = "1"
		return false
	} else if len(*stone)%2 == 0 {
		midpoint := len(*stone) / 2

		left := (*stone)[:midpoint]
		right := (*stone)[midpoint:]

		removeLeadingZeros(&left)
		removeLeadingZeros(&right)

		*out1 = left
		*out2 = right
		return true
	} else {
		stone_value, err := strconv.ParseInt(*stone, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		stone_value *= 2024
		new_stone := strconv.FormatInt(stone_value, 10)
		*out1 = new_stone
		return false
	}
}

func updateStoneConcurrent(stone *string, id int) UpdateResult {
	var out1 string
	var out2 string
	used_out2 := updateStone(stone, &out1, &out2)

	if used_out2 {
		return UpdateResult{out1, out2, id}
	} else {
		return UpdateResult{out1, "", id}
	}
}

type UpdateResult struct {
	stone1     string
	new_stone2 string
	id         int
}

func createUpdateStoneRoutine(stone *string, id int, output chan UpdateResult) {
	output <- updateStoneConcurrent(stone, id)
}

func updateStones(stones *[]string) *[]string {
	outputs := make(chan UpdateResult, len(*stones))

	for i := len(*stones) - 1; i >= 0; i-- {
		go createUpdateStoneRoutine(&(*stones)[i], i, outputs)
	}

	results := make([]UpdateResult, len(*stones))
	for i := 0; i < len(*stones); i++ {
		out := <-outputs
		results[out.id] = out
	}

	slices.SortFunc(results, func(a UpdateResult, b UpdateResult) int { return a.id - b.id })

	for i := len(*stones) - 1; i >= 0; i-- {
		(*stones)[i] = results[i].stone1
		if results[i].new_stone2 != "" {
			*stones = slices.Insert(*stones, i+1, results[i].new_stone2)
		}
	}

	// fmt.Println("result ", *stones)

	return stones
}

func main() {
	// defer profile.Start().Stop()
	input, err := libaoc.GrabInput("sample_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stones := strings.Split(input, " ")

	fmt.Println(stones)
	stones_ptr := &stones
	for i := 0; i < 40; i++ {
		stones_ptr = updateStones(stones_ptr)
		fmt.Println("iteration ", i)
	}

	fmt.Println("stones count: ", len(stones))
}
