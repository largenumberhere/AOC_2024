package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type CPU struct {
	registerA           int
	registerB           int
	registerC           int
	halted              bool
	instruction_pointer int
	ticks               int
}

func makeCpu() CPU {
	cpu := CPU{
		registerA:           0,
		registerB:           0,
		registerC:           0,
		halted:              false,
		instruction_pointer: 0,
		ticks:               0,
	}

	return cpu
}

type Instructions struct {
	values []int
}

/*
Parses a file like this using heuristics
```
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
```
*/
func parseFile(path string) (CPU, Instructions, error) {

	lines, err := aoc_lib.GrabLines(path)
	if err != nil {
		return CPU{}, Instructions{}, err
	}

	register_a := 0
	register_b := 0
	register_c := 0

	rega_number := lines[0][12:]
	regb_number := lines[1][12:]
	regc_number := lines[2][12:]

	register_a, err = strconv.Atoi(rega_number)
	if err != nil {
		return CPU{}, Instructions{}, err
	}
	register_b, err = strconv.Atoi(regb_number)
	if err != nil {
		return CPU{}, Instructions{}, err
	}
	register_c, err = strconv.Atoi(regc_number)
	if err != nil {
		return CPU{}, Instructions{}, err
	}

	cpu := makeCpu()
	cpu.registerA = register_a
	cpu.registerB = register_b
	cpu.registerC = register_c

	instructions := Instructions{values: make([]int, 0, 32)}

	instruction_tribits := lines[4][9:]
	for _, n := range strings.Split(instruction_tribits, ",") {
		bit, err := strconv.Atoi(n)
		if err != nil {
			return cpu, instructions, err
		}

		instructions.values = append(instructions.values, bit)
	}

	return cpu, instructions, nil
}

func (cpu *CPU) isHalted() bool {
	return cpu.halted
}

const (
	instructionADV = 0
	instructionBXL = 1
	instructionBST = 2
	instructionJNZ = 3
	instructionBXC = 4
	instructionOUT = 5
	instructionBDV = 6
	instructionCDV = 7
)

func evaluateoperandLiteral(tribit int) int {
	return tribit
}

func evaluateOperandCombination(tribit int, rega int, regb int, regc int) int {
	switch tribit {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return rega
	case 5:
		return regb
	case 6:
		return regc
	case 7:
		log.Fatal("unsupported combination value 7")
		return -1
	default:
		log.Fatal("invalid combination value")
		return -1
	}
}

func div(operand int, rega int, regb int, regc int) int {
	arg_value := evaluateOperandCombination(operand, rega, regb, regc)
	return rega / (int(math.Pow(2, float64(arg_value))))
}

func xor(a int, b int) int {
	res := a ^ b
	println("xor of ", a, "and", b, " = ", res)

	return res
}

func (cpu *CPU) tickTock(instructions Instructions) (int, bool) {
	out_value := -1
	jumped := false

	if cpu.instruction_pointer > len(instructions.values) {
		cpu.halted = true
	}

	if cpu.halted {
		return out_value, (out_value != -1)
	}

	rega_in := cpu.registerA
	regb_in := cpu.registerB
	regc_in := cpu.registerC

	// if not branching
	instruction_bits := instructions.values
	bit_pos := cpu.instruction_pointer
	instruction := instruction_bits[bit_pos : bit_pos+2]

	operand := instruction[1]

	switch instruction[0] {
	case instructionADV:
		cpu.registerA = div(operand, rega_in, regb_in, regc_in)
	case instructionBXL:
		arg_value := evaluateoperandLiteral(operand)
		cpu.registerB = xor(regb_in, arg_value)
	case instructionBST:
		arg_value := evaluateOperandCombination(operand, rega_in, regb_in, regc_in)
		cpu.registerB = arg_value % 8
	case instructionJNZ:
		if rega_in != 0 {
			cpu.instruction_pointer = evaluateoperandLiteral(operand)
			jumped = true
		}
	case instructionBXC:
		// arg_value := evaluateOperandCombination(operand, rega_in, regb_in, regc_in)
		cpu.registerB = xor(regb_in, regc_in)
	case instructionOUT:
		arg_value := evaluateOperandCombination(operand, rega_in, regb_in, regc_in)
		out_value = arg_value % 8
	case instructionBDV:
		cpu.registerB = div(operand, rega_in, regb_in, regc_in)
	case instructionCDV:
		cpu.registerC = div(operand, rega_in, regb_in, regc_in)
	}

	// if not jumped
	if !jumped {
		cpu.instruction_pointer += 2
	}
	cpu.ticks += 1
	return out_value, (out_value != -1)
}

func main() {
	cpu, instructions, err := parseFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output_buffer := make([]int, 0)

	fmt.Println(cpu, instructions)
	for !cpu.isHalted() {
		output, is_output := cpu.tickTock(instructions)
		fmt.Println("ticks:", cpu.ticks)
		if is_output {
			fmt.Println("output ", output)
			output_buffer = append(output_buffer, output)
		} else {
			fmt.Println("output nil")
		}
		fmt.Println(cpu)
		fmt.Println()
	}

	for i, v := range output_buffer {
		fmt.Print(v)
		if i < len(output_buffer)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()

}
