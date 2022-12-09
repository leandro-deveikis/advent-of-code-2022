package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("Challenge 1: %s \n", challenge(true))
	fmt.Printf("Challenge 1: %s \n", challenge(false))
}

func challenge(shouldReverse bool) string {
	var cargoCrane = []string{
		"ZPMHR",
		"PCJB",
		"SNHGLCD",
		"FTMDQSRL",
		"FSPQBTZM",
		"TFSZBG",
		"NRV",
		"PGLTDVCM",
		"WQNJFML",
	}
	//
	input, err := os.Open("Day05/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	for s.Scan() {
		re := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
		match := re.FindStringSubmatch(s.Text())
		from, err := strconv.Atoi(match[2])
		check(err)
		to, err := strconv.Atoi(match[3])
		check(err)
		amount, err := strconv.Atoi(match[1])
		check(err)
		cargoCrane = performOperation(cargoCrane, amount, from, to, shouldReverse)
	}
	return getLastItemFromEach(cargoCrane)
}

func getLastItemFromEach(crane []string) string {
	result := ""
	for _, s := range crane {
		result = result + s[len(s)-1:]
	}
	return result
}

func performOperation(crane []string, amount int, from int, to int, shouldReverse bool) []string {
	a := crane[from-1][len(crane[from-1])-amount:]
	if shouldReverse {
		a = reverseString(a)
	}
	crane[to-1] = crane[to-1] + a
	crane[from-1] = crane[from-1][:len(crane[from-1])-amount]
	return crane
}

// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

/**
            [L] [M]         [M]
        [D] [R] [Z]         [C] [L]
        [C] [S] [T] [G]     [V] [M]
[R]     [L] [Q] [B] [B]     [D] [F]
[H] [B] [G] [D] [Q] [Z]     [T] [J]
[M] [J] [H] [M] [P] [S] [V] [L] [N]
[P] [C] [N] [T] [S] [F] [R] [G] [Q]
[Z] [P] [S] [F] [F] [T] [N] [P] [W]
 1   2   3   4   5   6   7   8   9
*/

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
