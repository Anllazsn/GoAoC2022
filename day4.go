package main

import (
	"strconv"
	"strings"
)

func Day4_PartOne(input string) int {
	var acc int = 0
	for _, row := range strings.Fields(input) {
		pairs := strings.Split(row, ",")
		assign_a := pairs[0]
		assign_b := pairs[1]

		a := strings.Split(assign_a, "-")
		b := strings.Split(assign_b, "-")

		min_a, _ := strconv.Atoi(a[0])
		max_a, _ := strconv.Atoi(a[1])
		min_b, _ := strconv.Atoi(b[0])
		max_b, _ := strconv.Atoi(b[1])

		if min_a <= min_b && max_a >= max_b {
			acc++
			continue
		} else if min_b <= min_a && max_b >= max_a {
			acc++
			continue
		}
	}

	return acc
}

func Day4_PartTwo(input string) int {
	var acc int = 0
	for _, row := range strings.Fields(input) {
		pairs := strings.Split(row, ",")
		assign_a := pairs[0]
		assign_b := pairs[1]

		a := strings.Split(assign_a, "-")
		b := strings.Split(assign_b, "-")

		min_a, _ := strconv.Atoi(a[0])
		//max_a, _ := strconv.Atoi(a[1])
		//min_b, _ := strconv.Atoi(b[0])
		max_b, _ := strconv.Atoi(b[1])

		if min_a <= max_b {
			acc++
			continue
		} else if min_a >= max_b {
			acc++
			continue
		}
	}

	return acc
}
