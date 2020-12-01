package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() []int {
	f, _ := ioutil.ReadFile("day1.txt")

	t := strings.TrimSpace(string(f))
	list := strings.Split(t, "\n")

	ints := []int{}

	for _, s := range list {
		number, _ := strconv.Atoi(s)
		ints = append(ints, number)
	}

	log.Print(ints)
	return ints
}

func getResultOfTwoEntries(list []int) {
	perm := [][]int{}
	for _, x := range list {
		for _, y := range list {
			s := []int{x, y}
			perm = append(perm, s)
		}
	}

	for i := range perm {
		x := perm[i][0]
		y := perm[i][1]
		sum := x + y
		if sum == 2020 {
			result := x * y
			s := fmt.Sprintf("%v * %v = %v", x, y, result)
			fmt.Println(s)
			break
		}
	}
}

func getResultOfThreeEntries(list []int) {
	perm := [][]int{}
	for _, x := range list {
		for _, y := range list {
			for _, z := range list {
				s := []int{x, y, z}
				perm = append(perm, s)
			}
		}
	}

	for i := range perm {
		x := perm[i][0]
		y := perm[i][1]
		z := perm[i][2]
		sum := x + y + z
		if sum == 2020 {
			result := x * y * z
			s := fmt.Sprintf("%v * %v * %v = %v", x, y, z, result)
			fmt.Println(s)
			break
		}
	}
}

func main() {
	input := getInput()

	getResultOfTwoEntries(input)
	getResultOfThreeEntries(input)
}
