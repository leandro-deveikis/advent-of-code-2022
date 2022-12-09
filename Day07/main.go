package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	challenge()
}

type dir struct {
	name string
	size int

	parent *dir
	dirs   []*dir
	files  []*file
}

type file struct {
	name string
	size int
}

func challenge() {
	input, err := os.Open("Day07/input")
	check(err)
	defer func() {
		err := input.Close()
		check(err)
	}()

	rootDir := &dir{name: "/"}
	currentDir := rootDir

	s := bufio.NewScanner(input)
	s.Scan()
	if t := s.Text(); t != "$ cd /" {
		panic("Error, first line not expected: " + t)
	}

	// this is used to parse the file line
	re := regexp.MustCompile(`^([0-9]+) ([a-zA-Z.]+)$`)

	for s.Scan() {
		t := s.Text()

		if t == "$ cd .." {
			currentDir = currentDir.parent
			continue
		} else if t == "$ cd /" {
			currentDir = rootDir
			continue
		} else if t == "$ ls" {
			// we just omit this one for now
			continue
		} else if strings.HasPrefix(t, "$ cd ") {
			// change directory
			fName := strings.Replace(t, "$ cd ", "", 1)
			found := false
			for _, f := range currentDir.dirs {
				if fName == f.name {
					currentDir = f
					found = true
					break
				}
			}
			// just to validate
			if !found {
				panic("CD did not work for input: " + fName)
			}
			continue
		} else if strings.HasPrefix(t, "dir ") {
			//create a directory
			currentDir.dirs = append(currentDir.dirs, &dir{
				name:   strings.Replace(t, "dir ", "", 1),
				parent: currentDir,
			})
		} else if re.Match([]byte(t)) {
			substr := re.FindStringSubmatch(t)
			size, err := strconv.Atoi(substr[1])
			check(err)
			currentDir.files = append(currentDir.files, &file{
				name: substr[2],
				size: size,
			})
		}
	}

	// now we complete the sizes for the complete tree, from the root
	fillSize(rootDir)

	total := findAnswerChallenge1(rootDir)
	fmt.Printf("Challenge 1: %d \n", total)

	// --- PART 2
	totalDisk := 70000000
	neededSpace := 30000000
	alreadyFreeSpace := totalDisk - rootDir.size
	neededToDeleteAtLeast := neededSpace - alreadyFreeSpace
	// now find the smallest directory which size is at least 'alreadyFreeSpace'
	challenge2 := totalDisk
	challenge2 = findAnswerChallenge2(rootDir, challenge2, neededToDeleteAtLeast)

	fmt.Printf("Challenge 2: %d \n", challenge2)

}

func findAnswerChallenge2(dir *dir, currentAnswer int, neededToDeleteAtLeast int) int {
	if isBetterAnswer(dir.size, currentAnswer, neededToDeleteAtLeast) {
		currentAnswer = dir.size
	}
	for _, d := range dir.dirs {
		size := findAnswerChallenge2(d, currentAnswer, neededToDeleteAtLeast)
		if isBetterAnswer(size, currentAnswer, neededToDeleteAtLeast) {
			currentAnswer = size
		}
	}
	return currentAnswer
}

func isBetterAnswer(size int, currentAnswer int, neededToDeleteAtLeast int) bool {
	return currentAnswer > size && size > neededToDeleteAtLeast
}

func findAnswerChallenge1(rootDir *dir) int {
	// find all the directories with a total size of at most 100000,
	// then calculate the sum of their total sizes
	sum := 0
	// we do this to skip the root dir
	for _, d := range rootDir.dirs {
		sum += _findAnswer(d)
	}
	return sum
}

func _findAnswer(dir *dir) int {
	sum := 0
	for _, d := range dir.dirs {
		sum += _findAnswer(d)
	}
	// now the current dir
	if dir.size <= 100000 {
		sum += dir.size
	}
	return sum
}

func fillSize(dir *dir) {
	// first fill the size for all the directories
	for _, d := range dir.dirs {
		fillSize(d)
	}

	// now complete this one
	for _, d := range dir.dirs {
		dir.size += d.size
	}
	for _, f := range dir.files {
		dir.size += f.size
	}
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
