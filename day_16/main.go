package main

import (
	"bytes"
	"fmt"
	"log"

	lib_aoc "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type Cell struct {
	x         int
	y         int
	neigbours []*Cell
	value     rune
}

type CellMap [][]Cell

func (cell Cell) String() string {
	buff := bytes.Buffer{}
	str := fmt.Sprintf("Cell {[%d][%d]", cell.x, cell.y)
	buff.WriteString(str)
	buff.WriteString(" neighbors = ")
	if cell.neigbours != nil {
		for _, cell_ptr := range cell.neigbours {
			x := (*cell_ptr).x
			y := (*cell_ptr).y

			str2 := fmt.Sprintf("[%d][%d], ", x, y)
			buff.WriteString(str2)
		}
	}

	buff.WriteByte('}')

	return buff.String()
}

func (cells CellMap) String() string {
	string_builder := bytes.Buffer{}

	for _, row := range cells {
		for _, item := range row {
			string_builder.WriteRune(rune(item.value))
		}
		string_builder.WriteRune('\n')
	}
	return string_builder.String()

}

type Point2D struct {
	x int
	y int
}

func (field *CellMap) getAtPoint(pos Point2D) *Cell {
	x := pos.x
	y := pos.y
	return &(*field)[y][x]
}

func parseFile(path string) (CellMap, error) {

	runes, err := lib_aoc.GrabRunesArray(path)
	if err != nil {
		return nil, err
	}

	// construct the cells
	field := make(CellMap, 0, len(runes))
	for _, row := range runes {
		line := make([]Cell, 0, len(runes[0]))
		for _, item := range row {
			line = append(line, Cell{value: item})
		}

		field = append(field, line)
	}

	// assign each cell's neighbours
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			cell_ptr := field.getAtPoint(Point2D{x: x, y: y})

			has_left := x > 0 && x < len(field)
			has_top := y > 0 && y < len(field[0])
			has_right := x+1 < len(field)
			has_down := y+1 < len(field)

			if has_left || has_top || has_right || has_down {
				(*cell_ptr).neigbours = make([]*Cell, 0, 4)
				if has_left {
					pos := Point2D{x: x - 1, y: y}
					cell2 := field.getAtPoint(pos)
					(*cell_ptr).neigbours = append((*cell_ptr).neigbours, cell2)
				}
				if has_top {
					pos := Point2D{x: x, y: y - 1}
					cell2 := field.getAtPoint(pos)
					(*cell_ptr).neigbours = append((*cell_ptr).neigbours, cell2)
				}
				if has_right {
					pos := Point2D{x: x + 1, y: y}
					cell2 := field.getAtPoint(pos)
					(*cell_ptr).neigbours = append((*cell_ptr).neigbours, cell2)
				}
				if has_down {
					pos := Point2D{x: x, y: y + 1}
					cell2 := field.getAtPoint(pos)
					(*cell_ptr).neigbours = append((*cell_ptr).neigbours, cell2)
				}
			} else {
				(*cell_ptr).neigbours = nil
			}
		}
	}

	// assign each cell's coordinates
	// assign each cell's neighbours
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			cell_ptr := field.getAtPoint(Point2D{x, y})

			(*cell_ptr).x = x
			(*cell_ptr).y = y
		}
	}

	return field, nil
}

func (cell *Cell) printNeighbours() {
	if cell.neigbours != nil {
		for _, neigbour := range cell.neigbours {
			fmt.Println(neigbour)
		}
	}
}

func main() {
	if lib_aoc.VersionCount < 2 {
		log.Fatal("version of lib_aoc too old")
	}

	cells, err := parseFile("sample_input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	cell := cells.getAtPoint(Point2D{x: 1, y: 1})

	fmt.Println(cell)

	fmt.Println("meow")
}
