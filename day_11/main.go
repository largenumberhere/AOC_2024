package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/pkg/profile"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type Stones []string

func removeLeadingZeros(integer string) (string, error) {

	a, err := strconv.ParseInt(integer, 10, 64)
	b := strconv.FormatInt(a, 10)
	if err != nil {
		return "", err
	}

	if b == "" {
		return "0", nil
	}

	return b, nil
}

func updateStone(stone string, out1 *string, out2 *string) bool {
	if stone == "0" {
		*out1 = "1"
		return false
	} else if len(stone)%2 == 0 {
		midpoint := len(stone) / 2
		var left, right string
		var err error
		left, err = removeLeadingZeros(stone[:midpoint])
		if err != nil {
			log.Fatal(err)
		}
		right, err = removeLeadingZeros(stone[midpoint:])
		if err != nil {
			log.Fatal(err)
		}

		*out1 = left
		*out2 = right
		return true
	} else {
		stone_value, err := strconv.ParseInt(stone, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		stone_value *= 2024
		new_stone := strconv.FormatInt(stone_value, 10)
		*out1 = new_stone
		return false
	}
}

func updateStones(stones *Stones) *Stones {
	out2 := ""
	var used_out2 bool
	for i := len(*stones) - 1; i >= 0; i-- {
		used_out2 = updateStone((*stones)[i], &((*stones)[i]), &out2)
		if used_out2 {
			*stones = slices.Insert(*stones, i+1, out2)
		}
	}

	return stones
}

func main() {
	defer profile.Start().Stop()
	input, err := libaoc.GrabInput("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stones := Stones(strings.Split(input, " "))

	fmt.Println(stones)
	stones_ptr := &stones
	for i := 0; i < 25; i++ {
		stones_ptr = updateStones(stones_ptr)
		// fmt.Println(stones)
		// fmt.Println()
		fmt.Println("iteration ", i)
	}

	fmt.Println("stones count: ", len(stones))
}
