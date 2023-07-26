package main

import (
	"fmt"
	"strings"
)

const (
	Crate   = 1
	Command = 2
)

type Stack = struct {
	crates []string
	no     int
}

type Lexer = struct {
	start, end, row int
	t               int
	value           string
}

func Day5_PartOne(input string) int {
	// move [qty] from [column] to [column]
	split := strings.Split(input, "\n")
	// var stacks []Stack
	var lexers []Lexer

	for i := 0; i < len(split); i++ {
		row := split[i]
		for j := 0; j < len(row); j++ {
			col := row[j]
			if col == '[' {
				lexer := Lexer{start: j, row: i, t: Crate}
				k := j + 1
				for ; k < len(row); k++ {
					if row[k] == ']' {
						lexer.end = k
						break
					}

					lexer.value += string(row[k])
				}
				j = k
				lexers = append(lexers, lexer)
			}
		}

		if strings.TrimSpace(row) == "" {
			break
		}
	}

	printLexers(lexers)

	return 0
}

func Day5_PartTwo(input string) int {
	return 0
	// var acc int = 0
	// for _, row := range strings.Fields(input) {
	// 	pairs := strings.Split(row, ",")
	// 	assign_a := pairs[0]
	// 	assign_b := pairs[1]

	// 	a := strings.Split(assign_a, "-")
	// 	b := strings.Split(assign_b, "-")

	// 	min_a, _ := strconv.Atoi(a[0])
	// 	max_a, _ := strconv.Atoi(a[1])
	// 	min_b, _ := strconv.Atoi(b[0])
	// 	max_b, _ := strconv.Atoi(b[1])

	// 	if (min_a >= min_b && min_a <= max_b) || (max_a >= min_b && max_a <= max_b) {
	// 		acc++
	// 		continue
	// 	} else if (min_b >= min_a && min_b <= max_a) || (max_b >= min_a && max_b <= max_a) {
	// 		acc++
	// 		continue
	// 	}
	// }

	// return acc
}

func printLexers(lexers []Lexer) {
	for _, l := range lexers {
		fmt.Printf("LEXER R:%d C:[%d..%d] > %v \n", l.row, l.start, l.end, l.value)
	}
}
