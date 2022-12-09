package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	challenge01()
	challenge02()

	fmt.Printf("Now with higher functions...\n")

	// It is kind of overkill to still use the methods within a function but it is easier
	fullyOverlaps := func(s1 section, s2 section) bool {
		return s1.fullyOverlaps(s2) || s2.fullyOverlaps(s1)
	}
	hasSomeOverlap := func(s1 section, s2 section) bool {
		return s1.hasSomeOverlap(s2) || s2.hasSomeOverlap(s1)
	}
	fmt.Printf("Challenge 1: %d \n", challengeHigherFunctions(fullyOverlaps))
	fmt.Printf("Challenge 2: %d \n", challengeHigherFunctions(hasSomeOverlap))
}

type section struct {
	start int
	end   int
}

func (s section) fullyOverlaps(s2 section) bool {
	return s.start <= s2.start && s.end >= s2.end
}

func (s section) hasSomeOverlap(s2 section) bool {
	return (s.start <= s2.start && s.end >= s2.start) ||
		(s.start <= s2.end && s.end >= s2.end)
}

func challenge01() {
	input, err := os.Open("Day04/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	count := 0
	for s.Scan() {
		sec1, sec2 := getSectionsFromLine(s.Text())
		if sec1.fullyOverlaps(sec2) || sec2.fullyOverlaps(sec1) {
			count++
		}
	}
	fmt.Printf("Challenge 1: %d \n", count)
}

func challengeHigherFunctions(f func(s section, s2 section) bool) int {
	input, err := os.Open("Day04/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	count := 0
	for s.Scan() {
		sec1, sec2 := getSectionsFromLine(s.Text())
		if f(sec1, sec2) {
			count++
		}
	}
	return count
}

func challenge02() {
	input, err := os.Open("Day04/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	count := 0
	for s.Scan() {
		sec1, sec2 := getSectionsFromLine(s.Text())
		if sec1.hasSomeOverlap(sec2) || sec2.hasSomeOverlap(sec1) {
			count++
		}
	}
	fmt.Printf("Challenge 1: %d \n", count)
}

func getSectionsFromLine(t string) (section, section) {
	strs := strings.Split(t, ",")
	return buildSection(strs[0]), buildSection(strs[1])
}

func buildSection(s string) section {
	strs := strings.Split(s, "-")
	start, err := strconv.Atoi(strs[0])
	check(err)
	end, err := strconv.Atoi(strs[1])
	check(err)
	return section{
		start: start,
		end:   end,
	}
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
