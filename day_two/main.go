package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getInput() [][]string {
	f, _ := ioutil.ReadFile("day2.txt")

	t := strings.TrimSpace(string(f))
	s := strings.Split(t, "\n")
	data := [][]string{}

	for _, v := range s {
		p := strings.Split(v, ":")

		rules := p[0]
		pw := p[1]

		d := []string{rules, pw}
		data = append(data, d)
	}

	return data
}

type Rule struct {
	min  int
	max  int
	char string
}

func parseRules(rule string) Rule {
	fields := strings.Fields(rule)
	minMax := strings.Split(fields[0], "-")

	char := fields[1]
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return Rule{char: char, min: min, max: max}
}

func validateMinMaxPasswords(data [][]string) {
	validPws := []string{}
	for _, v := range data {
		rules := parseRules(v[0])
		pw := v[1]

		reps := strings.Count(pw, rules.char)

		if reps <= rules.max && reps >= rules.min {
			validPws = append(validPws, pw)
		}
	}
	log.Print(len(validPws))
}

func validatePositionPassword(data [][]string) {
	validPws := []string{}
	for _, v := range data {
		rules := parseRules(v[0])
		pw := v[1]
		t := strings.TrimSpace(pw)
		s := strings.Split(t, "")

		if s[rules.min-1] == rules.char && s[rules.max-1] != rules.char {
			validPws = append(validPws, pw)
		} else if s[rules.min-1] != rules.char && s[rules.max-1] == rules.char {
			validPws = append(validPws, pw)
		}

	}
	log.Print(len(validPws))
}

func main() {
	d := getInput()
	validateMinMaxPasswords(d)
	validatePositionPassword(d)

}
