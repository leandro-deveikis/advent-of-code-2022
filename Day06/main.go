package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Challenge 1: %d \n", challenge(4))
	fmt.Printf("Challenge 2: %d \n", challenge(14))
}

func challenge(length int) int {
	input, err := os.Open("Day06/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	s.Scan()
	t := s.Text()

	for i := 0; i < len(t)-length; i++ {
		subT := t[i : i+length]
		if areAllCharDifferent(subT) {
			return i + length
		}
	}
	panic("Result not found")
}

func areAllCharDifferent(t string) bool {
	tmp := make(map[int32]int, 0)
	for _, s := range t {
		if tmp[s] == 1 {
			return false
		} else {
			tmp[s] = 1
		}
	}
	return true
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
