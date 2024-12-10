package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

// max value
const empty_sentinel bignum = 9223372036854775807

type bignum int64

func expandDisk(disk []rune) ([]bignum, error) {
	// runes := []rune{}

	chunks := []bignum{}

	is_file := true
	var file_id bignum = 0
	for _, item := range disk {
		if is_file {
			file_bytes_count, err := strconv.Atoi(string(item))
			if err != nil {
				return nil, err
			}

			for i := 0; i < file_bytes_count; i++ {
				chunks = append(chunks, file_id)
			}

			file_id += 1
		} else {
			gap_size, err := strconv.Atoi(string(item))
			if err != nil {
				return nil, err
			}

			for i := 0; i < gap_size; i++ {
				chunks = append(chunks, bignum(empty_sentinel))
			}

		}

		is_file = !is_file
	}

	return chunks, nil
}

func printRunes(runes []rune) {
	for _, v := range runes {
		fmt.Print(string(v))
	}
	fmt.Println("")
}

func rightmostNonGap(runes *[]bignum) bignum {
	len := len(*runes)
	var i bignum = bignum(len) - 1
	for ; i >= 0; i-- {
		if (*runes)[i] != empty_sentinel {
			return i
		}
	}

	return -1
}

func printExpanded(items []bignum) {
	for _, v := range items {
		if v == empty_sentinel {
			fmt.Print(".")
		} else {
			fmt.Print(strconv.Itoa(int(v)))
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func defragmentExpanded(runes *[]bignum) {
	for {
		to_index := bignum(slices.IndexFunc(*runes, func(r bignum) bool { return r == empty_sentinel }))
		from_index := rightmostNonGap(runes)

		// break if reached bounds of slice
		if to_index == -1 || from_index == -1 {
			return
		}

		// break if no more space on left
		if to_index > from_index {
			break
		}

		(*runes)[to_index] = (*runes)[from_index]
		(*runes)[from_index] = empty_sentinel
	}
}

func checksumExpanded(numbers *[]bignum) bignum {
	var sum bignum
	var multiplier bignum = -1
	for _, item := range *numbers {
		multiplier += 1
		if item == empty_sentinel {
			continue
		}

		item_val := item

		sum += bignum(item_val) * bignum(multiplier)

	}

	return sum
}

func dumpString(fileName string, data string) {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(data)
	file.Close()
}

func main() {
	// uuid := 0

	rows, err := aoc_lib.GrabRunesArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	row := rows[0]
	expanded, err := expandDisk(row)
	if err != nil {
		log.Fatal(err)
	}

	// printExpanded(expanded)
	defragmentExpanded(&expanded)
	// printExpanded(expanded)
	sum := checksumExpanded(&expanded)
	fmt.Println(sum)
}
