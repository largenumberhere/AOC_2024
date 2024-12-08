package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func grabInput(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	str := string(bytes)

	// remove trailing and leading whitespace
	str = strings.Trim(str, " \n\r\t")

	// remove pesky carriage returns
	str = strings.Replace(str, "\r", "", -1)

	return str, nil
}

func grabLines(path string) ([]string, error) {
	file_contents, err := grabInput(path)
	if err != nil {
		return []string{}, err
	}

	var lines = strings.Split(file_contents, "\n")

	return lines, nil
}

func printRunes(runes *[][]rune) {
	for _, i := range *runes {
		for _, j := range i {
			if unicode.IsPrint(j) {
				fmt.Print(string(j))
			} else {
				fmt.Print(" ", j, " ")
			}
		}
		fmt.Println("")
	}

}

func grabRunesArray(path string) ([][]rune, error) {
	var lines_runes [][]rune

	lines, err := grabLines(path)
	if err != nil {
		return [][]rune{}, err
	}

	for _, line := range lines {
		var line_rune []rune
		for _, rune := range line {
			if rune == '\r' {
				// ignore pesky carriage returns
				continue
			}
			line_rune = append(line_rune, rune)
		}

		lines_runes = append(lines_runes, line_rune)
	}

	return lines_runes, nil
}

func makeRunes(rows int, columns int, initialized_value rune) [][]rune {
	runes := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		runes[i] = make([]rune, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			runes[i][j] = initialized_value
		}
	}

	return runes
}
