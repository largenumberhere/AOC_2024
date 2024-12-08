package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func stepNextFieldPos(antenna_pos Point2D, field *[][]rune) Option[Point2D] {
	if antenna_pos.x == -1 && antenna_pos.y == -1 {
		antenna_pos = Point2D{0, 0}
	} else {

		// increase counter
		antenna_pos.x += 1
	}

	// wrap around if at end of line
	if antenna_pos.x >= len(*field) {
		antenna_pos.x = 0
		antenna_pos.y += 1
	}

	// return early if out of range
	if antenna_pos.y >= len((*field)[0]) {
		return none[Point2D]()
	}

	if antenna_pos.x < 0 || antenna_pos.y < 0 {
		return none[Point2D]()
	}

	return some(antenna_pos)

}

func findNextAntennaPos(antenna_pos Point2D, field *[][]rune) Option[Point2D] {
	for {
		step := stepNextFieldPos(antenna_pos, field)
		if !step.some {
			return none[Point2D]()
		}

		step_value := step.data
		if ((*field)[step_value.y][step_value.x]) != '.' {
			return some(step_value)
		}

		antenna_pos = step_value
	}

}

func allAntennaPositions(field *[][]rune) []Point2D {
	var antennas []Point2D

	point := Point2D{-1, -1}
	for {
		a := findNextAntennaPos(point, field)
		if a.isNone() {
			break
		}

		point = a.unwrap()
		antennas = append(antennas, point)
	}

	return antennas
}

func collectAntennas(filed *[][]rune) []AntennaGroup {
	all := allAntennaPositions(filed)

	var groups []AntennaGroup

	for _, a := range all {
		marker := (*filed)[a.y][a.x]
		pos := findMarker(groups, marker)
		if pos == -1 {
			// create new group
			groups = append(groups, AntennaGroup{
				positions: []Point2D{a},
				marker:    marker,
			})

		} else {
			// append to group
			groups[pos].positions = append(groups[pos].positions, a)
		}
	}

	return groups
}

type AntennaGroup struct {
	positions []Point2D
	marker    rune
}

func (antennaGroup AntennaGroup) String() string {
	var buff bytes.Buffer
	if unicode.IsPrint(antennaGroup.marker) {
		buff.WriteRune('\'')
		buff.WriteRune(antennaGroup.marker)
		buff.WriteRune('\'')
	} else {
		a := strconv.Itoa(int(antennaGroup.marker))
		buff.WriteString(a)
	}
	buff.WriteString(": ")
	for i, v := range antennaGroup.positions {
		x := strconv.Itoa(v.x)
		y := strconv.Itoa(v.y)

		buff.WriteString("[")
		buff.WriteString(x)
		buff.WriteString(",")
		buff.WriteString(y)
		buff.WriteString("]")

		if i < len(antennaGroup.positions)-1 {
			buff.WriteString(" ")
		}

	}
	buff.WriteString(", ")
	return buff.String()
}

func findMarker(group []AntennaGroup, marker rune) int {
	for i, v := range group {
		if v.marker == marker {
			return i
		}
	}

	return -1
}

func main() {
	lines, err := grabRunesArray("sample_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	printRunes(&lines)
	antinodes_map := makeRunes(len(lines), len(lines[0]), '.')
	printRunes(&antinodes_map)

	groups := collectAntennas(&lines)
	fmt.Println(groups)

	// antenna_positions := allAntennaPositions(&lines)
	// fmt.Println(antenna_positions)
	// last_antenna_pos := Point2D{0, 0}

	// for {
	// 	a := findNextAntennaPos(last_antenna_pos, &lines)
	// 	if a.some == false {
	// 		return
	// 	}

	// 	last_antenna_pos = a.data
	// 	fmt.Println(a)

	// }
}
