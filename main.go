package main

import (
	"fmt"
	"os"
)

func Read(file string) string {
	var val string

	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	val = string(dat)

	return val
}

func Day2() {
	fmt.Println("DAY 2")
	input := Read("day2_input.txt")
	partOne := Day2_PartOne(input)
	partTwo := Day2_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func Day3() {
	fmt.Println("DAY 3")
	input := Read("day3_input.txt")
	partOne := Day3_PartOne(input)
	partTwo := Day3_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func Day4() {
	fmt.Println("DAY 4")
	input := Read("day4_input.txt")
	partOne := Day4_PartOne(input)
	partTwo := Day4_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func main() {
	fmt.Println("ADVENT OF CODE")
	Day4()
}
