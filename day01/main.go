package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInputFile(path string) []int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, 0, len(lines))
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			// Ignore empty lines
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func calcFuel(mass int) int {
	result := mass / 3
	result -= 2
	return result
}

// Calculate how much fuel the inputet fuel needs
func calcFuelFuel(fuel int) int {
	result := 0

	for {
		fuel = calcFuel(fuel)
		if fuel <= 0{
			break
		}
		result += fuel
	}

	return result
}

func part1(masses []int){
	fuel := 0
	for _, mass := range masses {
		fuel += calcFuel(mass)
	}

	fmt.Println("The fuel needed is: ", fuel)
}

func part2(masses []int){
	fuel := 0
	for _, mass := range masses {
		// Calculate how much fuel the mass needs
		massFuel := calcFuel(mass)

		// Calculate how much fuel the fuel for that mass needs
		fuel += massFuel + calcFuelFuel(massFuel)
	}

	fmt.Println("The fuel needed is: ", fuel)
}

func main() {
	masses := readInputFile("input.txt")

	fmt.Println(" --- Part 1 ---")
	part1(masses)

	fmt.Println("\n --- Part 2 ---")
	part2(masses)
}
