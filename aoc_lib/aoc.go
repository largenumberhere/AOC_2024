package libaoc

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// mroww

/*
 *	Read the aoc file and remove any spurious whitespace
 */
func GrabInput(path string) (string, error) {
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

/*
 *	Parse the lines of an aoc file
 */
func GrabLines(path string) ([]string, error) {
	file_contents, err := GrabInput(path)
	if err != nil {
		return []string{}, err
	}

	var lines = strings.Split(file_contents, "\n")

	return lines, nil
}

/*
 *	Print a 2d array of runes
 */
func PrintRunes(runes [][]rune) {
	for _, i := range runes {
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

/*
*	Parse an aoc file into a 2d array of runes
*
 */
func GrabRunesArray(path string) ([][]rune, error) {
	var lines_runes [][]rune

	lines, err := GrabLines(path)
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

/*
 *	Initalize a 2d array of runes
 */
func MakeRunes(rows int, columns int, initialized_value rune) [][]rune {
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

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// destructively modifies the input and returns a new array as a result
func Remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
