package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type Bit bool

type WireKey string

type WireOperation int

const (
	WireOperationNone WireOperation = iota
	WireOperationAnd
	WireOperationOr
	WireOperationXOR
)

type WireExpression struct {
	left      WireKey
	right     WireKey
	operation WireOperation
}

type Wire struct {
	key        WireKey
	isConstant bool

	// only valid if hasConstant is true
	constantValue Bit

	// only valid if hasConstant is false
	expression WireExpression
	wasSolved  bool
}

func (wire Wire) String() string {
	sb := bytes.Buffer{}
	sb.WriteString("Wire {  ")
	sb.WriteString("key: ")
	sb.WriteString(string(wire.key))
	sb.WriteString(", ")

	sb.WriteString("value: ")
	if wire.isConstant {
		if !wire.constantValue {
			sb.WriteRune('0')
		} else {
			sb.WriteRune('1')
		}
	} else {

		sb.WriteString(string(wire.expression.left))
		sb.WriteString(" ")

		if wire.expression.operation == WireOperationAnd {
			sb.WriteString("AND")
		} else if wire.expression.operation == WireOperationOr {
			sb.WriteString("Or")
		} else if wire.expression.operation == WireOperationXOR {
			sb.WriteString("XOR")
		}

		sb.WriteString(" ")
		sb.WriteString(string(wire.expression.right))

	}
	sb.WriteString("  }  ")
	return sb.String()
}

func parseWires(file_lines []string) ([]Wire, error) {
	input_wires := make([]Wire, 0)

	for line_number, line := range file_lines {

		if len(strings.Trim(line, "\r\n ")) == 0 {
			continue
		}

		fmt.Println(line)

		is_constant := strings.Contains(line, ":")
		parts := strings.Split(line, " ")
		if !is_constant {
			var wire Wire

			// is a calculated wire
			arg_0_wire := parts[0]
			operation_wire := parts[1]
			arg_1_wire := parts[2]
			_ = parts[3] // contains `->` which has no meaning
			destination_wire := parts[4]

			wire.isConstant = false
			wire.key = WireKey(destination_wire)
			wire.expression.left = WireKey(arg_0_wire)
			wire.expression.right = WireKey(arg_1_wire)

			if operation_wire == "AND" {
				wire.expression.operation = WireOperationAnd
			} else if operation_wire == "XOR" {
				wire.expression.operation = WireOperationXOR
			} else if operation_wire == "OR" {
				wire.expression.operation = WireOperationOr
			} else {
				err := fmt.Errorf("%s. offending line: %d '%s'", "invalid operation in lines. Expected 'XOR', 'OR' or 'AND'", line_number, line)
				return nil, err
			}

			input_wires = append(input_wires, wire)

		} else {
			var wire Wire

			// is a constant wire
			destination_wire := parts[0][0:3]
			value := strings.Trim(parts[1], "\r\n ")

			wire.isConstant = true

			if value == "1" {
				wire.constantValue = true
			} else if value == "0" {
				wire.constantValue = false
			} else {
				err := fmt.Errorf("invalid bit in line number %d. Expected 0 or one, given:'%s'", line_number, value)
				return nil, err
			}

			wire.key = WireKey(destination_wire)

			input_wires = append(input_wires, wire)
		}

	}

	return input_wires, nil
}

func canBeSolved(wire Wire, wires []Wire) bool {
	if wire.isConstant {
		// it is already solved
		return false
	}

	contains_left := slices.ContainsFunc(wires, func(w Wire) bool { return w.isConstant && w.key == wire.expression.left })
	if !contains_left {
		return false
	}

	contains_right := slices.ContainsFunc(wires, func(w Wire) bool { return w.isConstant && w.key == wire.expression.right })
	if !contains_right {
		return false
	}

	return true
}

func getWireFromKey(value WireKey, wires *[]Wire) *Wire {
	for i := 0; i < len(*wires); i++ {
		if (*wires)[i].key == value {
			return &(*wires)[i]
		}
	}

	return nil
}

func getWireValueFromKey(value WireKey, wires []Wire) (Bit, bool) {
	wire_ptr := getWireFromKey(value, &wires)

	if wire_ptr.isConstant {
		return wire_ptr.constantValue, true
	} else {
		return Bit(false), false
	}
}

func solve(wire_ptr *Wire, wires []Wire) {
	fmt.Print("solved ", *wire_ptr)

	// parse fields
	bit_left, ok := getWireValueFromKey(wire_ptr.expression.left, wires)
	if !ok {
		panic("meow")
	}

	var left int
	if bit_left == false {
		left = 0
	} else if bit_left == true {
		left = 1
	}

	bit_right, ok := getWireValueFromKey(wire_ptr.expression.right, wires)
	if !ok {
		panic("meow")
	}

	var right int
	if bit_right == false {
		right = 0
	} else if bit_right == true {
		right = 1
	}

	operation := wire_ptr.expression.operation

	// solve
	var result int
	switch operation {
	case WireOperationAnd:
		result = left & right
	case WireOperationOr:
		result = left | right
	case WireOperationXOR:
		result = left ^ right
	}

	// update
	var result_bit Bit
	if result == 0 {
		result_bit = Bit(false)
	} else if result == 1 {
		result_bit = Bit(true)
	} else {
		panic("invalid result")
	}

	wire_ptr.constantValue = result_bit
	wire_ptr.isConstant = true
	wire_ptr.wasSolved = true
}

func trySolveNext(wires *[]Wire) bool {
	for i := 0; i < len(*wires); i++ {
		wire_ptr := &(*wires)[i]
		if canBeSolved(*wire_ptr, *wires) {
			solve(wire_ptr, *wires)
			return true
		}
	}

	return false

}

func onlyWiresStartingWithZ(wires *[]Wire) {
	for i := len(*wires) - 1; i >= 0; i-- {
		key := (*wires)[i].key
		if !strings.HasPrefix(string(key), "z") {
			*wires = aoc_lib.Remove(*wires, i)
		}
	}
}

func wireToBinary(wires []Wire) string {
	sb := bytes.Buffer{}
	for i := len(wires) - 1; i >= 0; i-- {
		if wires[i].constantValue {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	return sb.String()
}

func wireToDecimal(wires []Wire) (int64, error) {
	binary := wireToBinary(wires)
	value, err := strconv.ParseInt(binary, 2, 64)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func main() {
	v := aoc_lib.VersionCount
	if v < 5 {
		log.Fatal("version 5 or greater of aoc_lib required")
	}

	file, err := aoc_lib.GrabLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	wires, err := parseWires(file)
	if err != nil {
		log.Fatal(err)
	}

	for wire_index, wire := range wires {
		fmt.Print("wire ", wire_index, " solveable: ", canBeSolved(wire, wires), " ")
		fmt.Println(wire)
	}

	for {
		solved_one := trySolveNext(&wires)
		if !solved_one {
			break
		}
	}

	onlyWiresStartingWithZ(&wires)
	slices.SortFunc(wires, func(a Wire, b Wire) int { return strings.Compare(string(a.key), string(b.key)) })

	for _, wire := range wires {
		fmt.Println(wire)
	}

	dec, err := wireToDecimal(wires)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dec)
}
