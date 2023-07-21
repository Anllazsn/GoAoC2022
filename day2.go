package main

import (
	"strings"
)

var rock, paper, scissor int = 1, 2, 3
var win, draw, lose int = 6, 3, 0

func Day2_PartOne(input string) int {
	var points int = 0
	for _, s := range strings.Split(input, "\n") {
		if s == "A X" {
			points += rock + draw
		} else if s == "A Y" {
			points += paper + win
		} else if s == "A Z" {
			points += scissor + lose
		} else if s == "B X" {
			points += rock + lose
		} else if s == "B Y" {
			points += paper + draw
		} else if s == "B Z" {
			points += scissor + win
		} else if s == "C X" {
			points += rock + win
		} else if s == "C Y" {
			points += paper + lose
		} else if s == "C Z" {
			points += scissor + draw
		}
	}

	return points
}

func Day2_PartTwo(input string) int {
	// X Lose
	// Y Draw
	// Z Win
	var points int = 0
	for _, s := range strings.Split(input, "\n") {
		if s == "A X" {
			points += scissor + lose
		} else if s == "A Y" {
			points += rock + draw
		} else if s == "A Z" {
			points += paper + win
		} else if s == "B X" {
			points += rock + lose
		} else if s == "B Y" {
			points += paper + draw
		} else if s == "B Z" {
			points += scissor + win
		} else if s == "C X" {
			points += paper + lose
		} else if s == "C Y" {
			points += scissor + draw
		} else if s == "C Z" {
			points += rock + win
		}
	}

	return points
}
