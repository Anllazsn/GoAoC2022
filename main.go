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
	input := Read("inputs/day2_input.txt")
	partOne := Day2_PartOne(input)
	partTwo := Day2_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func Day3() {
	fmt.Println("DAY 3")
	input := Read("inputs/day3_input.txt")
	partOne := Day3_PartOne(input)
	partTwo := Day3_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func Day4() {
	fmt.Println("DAY 4")
	input := Read("inputs/day4_input.txt")
	partOne := Day4_PartOne(input)
	partTwo := Day4_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func Day5() {
	fmt.Println("DAY 5")
	input := Read("inputs/day5_input.txt")
	partOne := Day5_PartOne(input)
	partTwo := Day5_PartTwo(input)

	fmt.Println("Part 1: ", partOne)
	fmt.Println("Part 2: ", partTwo)
}

func main() {
	fmt.Println("ADVENT OF CODE")
	//Day5()
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	res := Day5_PartOne(input)
	//ReadStacks(input)

	//fmt.Println(string(input[4:7][1]))

	fmt.Println(res)
}
