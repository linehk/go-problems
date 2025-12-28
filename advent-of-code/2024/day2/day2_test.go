package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		t.Fatal("can't open input file")
	}
	defer file.Close()

	var safeCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineStrSlice := strings.Split(line, " ")

		lineSlice := make([]int, 0, len(lineStrSlice))
		for _, s := range lineStrSlice {
			v, err := strconv.Atoi(s)
			if err != nil {
				t.Fatal(err)
			}
			lineSlice = append(lineSlice, v)
		}

		// false -> increasing
		// true -> decreasing
		order := false
		if lineSlice[0] > lineSlice[1] {
			order = true
		}

		safe := true
		for i := range lineSlice {
			if i == len(lineSlice)-1 {
				break
			}

			level := lineSlice[i] - lineSlice[i+1]

			if level == 0 {
				safe = false
			}
			if level > 0 && !order {
				safe = false
			}
			if level < 0 && order {
				safe = false
			}
			if level < 0 {
				level = -level
			}
			if level > 3 {
				safe = false
			}

			if !safe {
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	if scanner.Err() != nil {
		t.Fatal("scanner err")
	}

	fmt.Println(safeCount)
}

func TestPart2(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		t.Fatal("can't open input file")
	}
	defer file.Close()

	var safeCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineStrSlice := strings.Split(line, " ")

		lineSlice := make([]int, 0, len(lineStrSlice))
		for _, s := range lineStrSlice {
			v, err := strconv.Atoi(s)
			if err != nil {
				t.Fatal(err)
			}
			lineSlice = append(lineSlice, v)
		}

		safe := isSafe(lineSlice)

		// use dampener
		if !safe {
			for i := range lineSlice {
				newLineSlice := make([]int, 0, len(lineSlice)-1)
				newLineSlice = append(newLineSlice, lineSlice[:i]...)
				newLineSlice = append(newLineSlice, lineSlice[i+1:]...)

				if isSafe(newLineSlice) {
					safe = true
					break
				}
			}
		}

		if safe {
			safeCount++
		}
	}

	if scanner.Err() != nil {
		t.Fatal("scanner err")
	}

	fmt.Println(safeCount)
}

func isSafe(lineSlice []int) bool {
	// false -> increasing
	// true -> decreasing
	order := false
	if lineSlice[0] > lineSlice[1] {
		order = true
	}

	safe := true
	for i := range lineSlice {
		if i == len(lineSlice)-1 {
			break
		}

		level := lineSlice[i] - lineSlice[i+1]

		if level == 0 {
			safe = false
		}
		if level > 0 && !order {
			safe = false
		}
		if level < 0 && order {
			safe = false
		}
		if level < 0 {
			level = -level
		}
		if level > 3 {
			safe = false
		}

		if !safe {
			break
		}
	}

	return safe
}
