package main

import (
	"fmt"
	"strings"
)

func Day3_PartOne(input string) int {
	var acc int
	for _, it := range strings.Fields(input) {
		m := len(it) / 2
		first := it[:m]
		second := it[m:]
		// a = 97 z = 122
		// A = 65 Z = 90

	out:
		for _, f := range first {
			for _, s := range second {
				if f == s {
					acc += runeToPriorite(f)
					break out
				}
			}
		}
	}

	return acc
}

func Day3_PartTwo(input string) int {
	rucksacks := strings.Fields(input)
	var acc int
	for i := 0; i < len(rucksacks); i++ {
		if i+2 > len(rucksacks) {
			panic(fmt.Sprintf("INDEX: %d, Impossível agrupar em grupos de 3", i))
		}

		acc += runeToPriorite(itemInElfGroup(rucksacks[i], rucksacks[i+1], rucksacks[i+2]))

		i += 2
	}
	return acc
}

func runeToPriorite(v rune) int {
	val := int(v)
	if val >= 97 && val <= 122 {
		return val - 96
	} else if val >= 65 && val <= 90 {
		return val - 38
	}

	panic(fmt.Sprintf("Valor %d não é um char válido", v))
}

func itemInElfGroup(one, two, three string) rune {
	for _, a := range one {
		for _, b := range two {
			for _, c := range three {
				if a == b && a == c {
					return a
				}
			}
		}
	}

	panic("Impossível encontrar item igual nas 3 rucksacks")
}
