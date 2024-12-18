// note to self: update module immedaitely to latest version with `env GOPROXY="direct" go get -u`

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type Queue[T any] struct {
	buffer *[]T
}

func (queue *Queue[T]) popLeft() (T, bool) {
	// if queue.buffer == nil {
	// 	var default_value T
	// 	return default_value, false
	// }

	if len(*queue.buffer) < 1 {
		var result T
		return result, false
	}

	item := (*queue.buffer)[0]
	(*queue.buffer) = (*queue.buffer)[1:]

	return item, true
}

func (queue *Queue[T]) pushRight(t T) {
	// if queue.buffer == nil {
	// 	buff := make([]T, 0, 8)
	// 	queue.buffer = buff
	// }

	*queue.buffer = append(*queue.buffer, t)

	// fmt.Println("appended to queue ", queue.buffer, "value", t)
}

func makeQueue[T comparable]() Queue[T] {
	buffer := make([]T, 0, 8)

	return Queue[T]{
		buffer: &buffer,
	}
}

func (queue *Queue[T]) count() int {
	return len(*queue.buffer)
}

func (queue *Queue[T]) isEmpty() bool {
	return queue.count() == 0
}

type Cell struct {
	pos   aoc_lib.Point2D
	value rune
}

func pushCell(queue *Queue[Cell], x int, y int, value rune) {
	point := aoc_lib.MakePoint2D(x, y)
	var cell Cell = Cell{pos: point, value: value}
	queue.pushRight(cell)
}

func popCell(queue *Queue[Cell]) (cell Cell, b bool) {
	return queue.popLeft()
}

func contains(queue *Queue[Cell], item Cell) bool {
	for _, v := range *queue.buffer {
		if v == item {
			return true
		}
	}

	return false
}

func (cell *Cell) neigbours(max_bounds aoc_lib.Point2D, field [][]rune) []Cell {
	x, y := cell.pos.GetXY()
	max_x, max_y := max_bounds.GetXY()

	neighbours := make([]Cell, 0, 4)
	has_left := x > 0 && x < max_x
	has_top := y > 0 && y < max_y
	has_right := x+1 < max_x
	has_down := y+1 < max_x

	if has_top {
		// up
		pos := aoc_lib.MakePoint2D(x, y-1)
		cell := Cell{pos: pos, value: field[y-1][x]}

		neighbours = append(neighbours, cell)

	}

	if has_left {
		//left
		pos := aoc_lib.MakePoint2D(x-1, y)
		cell := Cell{pos: pos, value: field[y][x-1]}

		neighbours = append(neighbours, cell)

	}

	if has_right {
		// right
		pos := aoc_lib.MakePoint2D(x+1, y)
		cell := Cell{pos: pos, value: field[y][x+1]}

		neighbours = append(neighbours, cell)

	}

	if has_down {
		// down
		pos := aoc_lib.MakePoint2D(x, y+1)
		cell := Cell{pos: pos, value: field[y+1][x]}

		neighbours = append(neighbours, cell)
	}

	return neighbours
}

func findBestPath(field [][]rune, destination aoc_lib.Point2D) []aoc_lib.Point2D {
	// a god awful but functional implementation of breadth-first search

	frontier := makeQueue[Cell]()
	origin := aoc_lib.MakePoint2D(0, 0)
	pushCell(&frontier, 0, 0, field[0][0])

	came_from := make(map[Cell]Cell, 0)
	came_from[Cell{pos: origin, value: field[0][0]}] = Cell{pos: aoc_lib.MakePoint2DUninit()}

	_, max_bounds := aoc_lib.RunesBounds(field)

	for {
		current, some := popCell(&frontier)
		if !some {
			break
		}

		for _, next := range current.neigbours(max_bounds, field) {
			_, is_in_history := came_from[next]
			if !is_in_history && next.value != '#' {
				x, y := next.pos.GetXY()
				pushCell(&frontier, x, y, next.value)
				came_from[next] = current
			}
		}

	}

	fmt.Println("", came_from)
	goal_pos := destination

	points := make([]aoc_lib.Point2D, 0)
	dest_x, dest_y := goal_pos.GetXY()
	current := Cell{pos: goal_pos, value: field[dest_y][dest_x]}
	ok := true
	aoc_lib.PrintRunes(field)

	i := 0
	for {
		i += 1
		if current.pos == aoc_lib.MakePoint2DUninit() {
			break
		}

		points = append(points, current.pos)

		fmt.Println(current)
		current, ok = came_from[current]

		if !ok {
			panic("Unable to find any path")
		}
	}

	return points
}

func main() {
	if aoc_lib.VersionCount < 5 {
		log.Fatal("At least version 0.0.5 of lib_aoc required, got ", aoc_lib.Version, ".")
	}

	positions, err := aoc_lib.GrabLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	positions_points := make([]aoc_lib.Point2D, 0)
	for _, point := range positions {
		parts := strings.Split(point, ",")
		left := parts[0]
		right := parts[1]

		left_int, err := strconv.Atoi(left)
		if err != nil {
			log.Fatal(err)
		}
		right_int, err := strconv.Atoi(right)
		if err != nil {
			log.Fatal(err)
		}

		point2d := aoc_lib.MakePoint2D(left_int, right_int)

		positions_points = append(positions_points, point2d)
	}

	field := aoc_lib.MakeRunes(71, 71, '.')

	i := 0
	for _, pos := range positions_points {
		x, y := pos.GetXY()
		fmt.Println(x, y)

		field[y][x] = '#'

		i += 1
		if i >= 1024 {
			break
		}
	}

	aoc_lib.PrintRunes(field)

	path := findBestPath(field, aoc_lib.MakePoint2D(70, 70))
	fmt.Println(path)
	fmt.Println(len(path) - 1)

}

// type Point2D struct {
// 	X, Y int
// }

/*
todo: make a search grid function from this
func canRobotMove(field [][]rune, direction Direction, is_gap bool, gap_pos Point2D) bool {
        robot_pos := findRobot(field)

        offset := direction.directionOffset()

        robot_pos = robot_pos.AddComponents(offset)
        // if the position is a wall, false
        if field[robot_pos.y][robot_pos.x] == '#' {
                return false
        }

        // check if there are boxes that can be moved
        if !is_gap {
                return false
        }

        // else the position is empty, true
        return true

}

*/
