package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	direction string
	steps     int
}

func main() {
	input, err := os.Open("Day09/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	motions := make([]line, 0)
	// build the map
	for s.Scan() {
		t := strings.Split(s.Text(), " ")
		s, err := strconv.Atoi(t[1])
		check(err)
		motions = append(motions, line{direction: t[0], steps: s})
	}

	// fmt.Printf("total lines: %d \n", len(motions))

	// CHALLENGE 1 - now count the trees visibility
	fmt.Printf("Challenge 1: %d \n", 0)

	// CHALLENGE 2 - get the highest scenic score
	fmt.Printf("Challenge 2: %d \n", 0)
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
