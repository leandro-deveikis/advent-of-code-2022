package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	challenge01()
	challenge02()
}

func challenge01() {
	// Win: 6 - Draw: 3 - Lost: 0
	// Rock (A - X): 1 - Paper (B - Y): 2 - Scissors (C - Z): 3
	// As this has only 9 combinations, we just precalculate them and store it into a matrix.
	// we could do something even better, intead of having this map, having a function that calculates them
	// according to the rules, so is not that hardcoded
	cheatsheet := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,

		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,

		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	input, err := os.Open("input")
	check(err)
	defer input.Close()

	s := bufio.NewScanner(input)
	t := 0
	for s.Scan() {
		t += cheatsheet[s.Text()]
	}

	fmt.Printf("Challenge 1: %d \n", t)
}

func challenge02() {
	// Win: 6 - Draw: 3 - Lost: 0
	// Rock (A - X): 1 - Paper (B - Y): 2 - Scissors (C - Z): 3
	cheatsheet_old := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,

		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,

		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	// X = loose, Y = draw, Z = win
	// The idea is that we transform into the old map, to reuse what I did on challenge 01
	cheatsheet_transformation := map[string]string{
		"A X": "A Z",
		"A Y": "A X",
		"A Z": "A Y",

		"B X": "B X",
		"B Y": "B Y",
		"B Z": "B Z",

		"C X": "C Y",
		"C Y": "C Z",
		"C Z": "C X",
	}

	input, err := os.Open("input")
	check(err)
	defer input.Close()

	s := bufio.NewScanner(input)
	t := 0
	for s.Scan() {
		t += cheatsheet_old[cheatsheet_transformation[s.Text()]]
	}

	fmt.Printf("Challenge 2: %d \n", t)
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
