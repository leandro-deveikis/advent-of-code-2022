package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Win: 6 - Draw: 3 - Lost: 0
	// Rock (A - X): 1 - Paper (B - Y): 2 - Scissors (C - Z): 3
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

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
