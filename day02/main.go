package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	strings "strings"
)

func readInput(path string) []int {
	// Read the files content as a byte array
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Unable to read input file")
	}

	// Convert the byte array into a slice of strings that hold numbers
	colums := strings.Split(string(content), ",")

	// convert the slice of strings into a slice of numbers
	intCode := make([]int, 0, len(colums))
	for _, colum := range colums {
		code, err := strconv.Atoi(colum)
		if err != nil {
			// ignore errors
			continue
		}

		intCode = append(intCode, code)
	}

	return intCode
}

func calcIntcode(originalIntCode []int, noun int, verb int) int {
	// Copy the original code into a temporary one so we don't change the original
	intCode := make([]int, len(originalIntCode))
	copy(intCode, originalIntCode)

	// Change the input
	intCode[1] = noun
	intCode[2] = verb

	// Setup the machine
	instructionPointer := 0
	hasExitOk := false

	// Calculate
	loop:
	for {
		opcode := intCode[instructionPointer]

		switch opcode {
		// Addition
		case 1:
			arg1 := intCode[instructionPointer + 1]
			arg2 := intCode[instructionPointer + 2]
			arg3 := intCode[instructionPointer + 3]

			intCode[arg3] = intCode[arg1] + intCode[arg2]
			instructionPointer += 4

		// Multiplication
		case 2:
			arg1 := intCode[instructionPointer + 1]
			arg2 := intCode[instructionPointer + 2]
			arg3 := intCode[instructionPointer + 3]

			intCode[arg3] = intCode[arg1] * intCode[arg2]
			instructionPointer += 4

		// Exit
		case 99:
			hasExitOk = true
			break loop

		// Bad instruction
		default:
			hasExitOk = false
			break loop
		}
	}

	// check if everything was ok
	if hasExitOk == false {
		log.Fatal("Error occoured :(")
	}

	return intCode[0]
}

// Part one of the daily challenge
func part1(intCode []int) {
	result := calcIntcode(intCode, 12, 2)
	fmt.Println("Result is: ", result)
}

// Part 2 of the daily challenge
func part2(intCode []int) {
	// Try all noun and verb combinations from 0 to 99
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			result := calcIntcode(intCode, noun, verb)

			// If the result matches the desired number we are done
			if result == 19690720 {
				fmt.Println("Result is: ", 100 * noun + verb)
				return
			}
		}
	}

	fmt.Println("Error: No verb & noun pair found")
}

func main() {
	intCode := readInput("input.txt")

	fmt.Println("--- Part 1 ---")
	part1(intCode)

	fmt.Println("\n--- Part 2 ---")
	part2(intCode)
}