package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var column1 []int
	var column2 []int

	// 1. scan data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ex: 31594   93577
		line := scanner.Text()

		lineStrSlice := strings.Split(line, "   ")
		lineIntVar1, err := strconv.Atoi(lineStrSlice[0])
		if err != nil {
			t.Fatal(err)
		}

		lineIntVar2, err := strconv.Atoi(lineStrSlice[1])
		if err != nil {
			t.Fatal(err)
		}

		column1 = append(column1, lineIntVar1)
		column2 = append(column2, lineIntVar2)
	}

	if scanner.Err() != nil {
		t.Fatal("scanner err")
	}

	// 2. sort column1 and column2
	sort.Ints(column1)
	sort.Ints(column2)

	// 3. get abs value and sum
	var sum int
	for i := range column1 {
		distance := column1[i] - column2[i]
		if distance < 0 {
			distance = -distance
		}
		sum += distance
	}

	fmt.Println(sum)
}

func TestPart2(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		t.Fatal("can't open input file")
	}
	defer file.Close()

	var column1 []int
	var column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ex: 31594   93577
		line := scanner.Text()

		lineStrSlice := strings.Split(line, "   ")
		lineIntVar1, err := strconv.Atoi(lineStrSlice[0])
		if err != nil {
			t.Fatal(err)
		}

		lineIntVar2, err := strconv.Atoi(lineStrSlice[1])
		if err != nil {
			t.Fatal(err)
		}

		column1 = append(column1, lineIntVar1)
		column2 = append(column2, lineIntVar2)
	}

	if scanner.Err() != nil {
		t.Fatal("scanner err")
	}

	var score int
	for _, v1 := range column1 {
		var sum int
		for _, v2 := range column2 {
			if v1 == v2 {
				sum += v1
			}
		}
		score += sum
	}

	fmt.Println(score)
}
