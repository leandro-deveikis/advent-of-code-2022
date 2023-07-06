package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("Day08/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	s := bufio.NewScanner(input)
	treeMap := make([][]int, 0)
	// build the map
	for s.Scan() {
		line := make([]int, 0)
		for _, t := range s.Text() {
			line = append(line, int(t)-'0')
		}
		treeMap = append(treeMap, line)
	}
	// CHALLENGE 1 - now count the trees visibility
	totalWithFunc := countTreeVisibility(treeMap)
	fmt.Printf("Challenge 1: %d \n", totalWithFunc)

	// CHALLENGE 2 - get the highest scenic score
	highestScore := getHighestScenicScore(treeMap)
	fmt.Printf("Challenge 2: %d \n", highestScore)
}

func countTreeVisibility(treeMap [][]int) int {
	total := 0
	for idx, line := range treeMap {
		// first line and last line are always visible
		if idx == 0 || idx == (len(treeMap)-1) {
			total += len(line)
			continue
		}
		total += countVisibility(line, treeMap, idx)
	}
	return total
}

func countVisibility(line []int, treeMap [][]int, idx int) int {
	total := 0
	for idx2, tree := range line {
		// first and last tree are always visible
		if idx2 == 0 || idx2 == (len(line)-1) {
			total += 1
			continue
		}

		for _, dir := range allDirections {
			if visibleFrom(tree, treeMap, idx, idx2, dir) {
				total += 1
				break
			}
		}
	}
	return total
}

func getHighestScenicScore(treeMap [][]int) int {
	highestScore := -1
	for idx, line := range treeMap {
		for idx2, tree := range line {
			currentScore := 1
			for _, dir := range allDirections {
				// fmt.Printf("X: %d Y: %d  \n", idx2, idx)
				currentScore *= getScoreForDir(tree, treeMap, idx, idx2, dir)
			}
			if currentScore > highestScore {
				highestScore = currentScore
			}
		}
	}
	return highestScore
}

func visibleFrom(tree int, treeMap [][]int, y int, x int, direction func(int, int) (int, int)) bool {
	maxX := len(treeMap) - 1
	maxY := len(treeMap[0]) - 1

	for {
		x, y = direction(x, y)
		if x < 0 || x > maxX || y < 0 || y > maxY {
			break
		}
		if tree <= treeMap[y][x] {
			return false
		}
	}
	return true
}

func getScoreForDir(tree int, treeMap [][]int, y int, x int, direction func(x int, y int) (int, int)) int {
	maxX := len(treeMap) - 1
	maxY := len(treeMap[0]) - 1
	total := 0

	for {
		x, y = direction(x, y)
		if x < 0 || x > maxX || y < 0 || y > maxY {
			break
		}
		currentTree := treeMap[y][x]
		if tree > currentTree {
			total += 1
		} else if tree == currentTree {
			total += 1
			break
		} else {
			break
		}
	}
	return total
}

var allDirections = []func(x int, y int) (int, int){dirUp, dirDown, dirLeft, dirRight}

var dirUp = func(x int, y int) (int, int) {
	return x, y - 1
}
var dirDown = func(x int, y int) (int, int) {
	return x, y + 1
}
var dirLeft = func(x int, y int) (int, int) {
	return x - 1, y
}
var dirRight = func(x int, y int) (int, int) {
	return x + 1, y
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
