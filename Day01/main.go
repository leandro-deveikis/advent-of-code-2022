package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Day 01 - https://adventofcode.com/2022/day/1
func main() {
	challenge1()
	challenge2()
}

func challenge1() {
	input, err := os.Open("input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	max, current := 0, 0
	for s.Scan() {
		t := s.Text()
		if t == "" {
			if max < current {
				max = current
			}
			current = 0
		} else {
			i, err := strconv.Atoi(t)
			check(err)
			current += i
		}
	}

	fmt.Printf("Challenge 1: %d \n", max)
}

func challenge2() {
	input, err := os.Open("input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	current := 0
	results := make([]int, 0)
	for s.Scan() {
		t := s.Text()
		if t == "" {
			results = append(results, current)
			current = 0
		} else {
			i, err := strconv.Atoi(t)
			check(err)
			current += i
		}
	}

	sort.Slice(results, func(i int, j int) bool {
		return results[i] > results[j]
	})

	fmt.Printf("Challenge 1 (new method): %d \n", results[0])
	fmt.Printf("Challenge 2: %d \n", results[0]+results[1]+results[2])
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
