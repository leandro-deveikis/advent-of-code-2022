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
	input, err := os.Open("Day03/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	var sum int32 = 0
	for s.Scan() {
		t := s.Text()
		first, second := splitString(t)
		sum += findDuplicatedItem(first, second)
	}
	fmt.Printf("Challenge 1: %d \n", sum)
}

func challenge02() {
	input, err := os.Open("Day03/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	var sum int32 = 0
	for s.Scan() {
		t1 := s.Text()
		s.Scan()
		t2 := s.Text()
		s.Scan()
		t3 := s.Text()

		m1 := countLetters(t1, make(map[int32]int, 0))
		m2 := countLetters(t2, make(map[int32]int, 0))
		m3 := countLetters(t3, make(map[int32]int, 0))

		for index, _ := range m1 {
			if m2[index] > 0 && m3[index] > 0 {
				sum += convertToPriority(index)
			}
		}
	}

	fmt.Printf("Challenge 2: %d \n", sum)
}

func findDuplicatedItem(first string, second string) int32 {
	m := countLetters(first, make(map[int32]int, 0))

	// now we go to the second
	for _, c := range second {
		if m[c] > 0 {
			// and we only want the one that is missing
			return convertToPriority(c)
		}
	}
	fmt.Printf("ERROR duplicated not found.\n")
	return 0
}

func countLetters(srt string, list map[int32]int) map[int32]int {
	for _, c := range srt {
		list[c] += 1
	}
	return list
}

func convertToPriority(c int32) int32 {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	if c >= 'a' {
		return c - 96
	} else {
		return c - 64 + 26
	}
}

func splitString(t string) (string, string) {
	l := len(t)
	return t[0 : l/2], t[l/2 : l]
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
