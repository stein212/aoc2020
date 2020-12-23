package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type void struct{}

var (
	member                 void
	passportFieldChecklist map[string]int
	colorRegex             *regexp.Regexp
	eyeColors              map[string]void
	passportIDRegex        *regexp.Regexp
)

func main() {
	passportFieldChecklist = map[string]int{
		"byr": 0,
		"iyr": 0,
		"eyr": 0,
		"hgt": 0,
		"hcl": 0,
		"ecl": 0,
		"pid": 0,
	}

	colorRegex, _ = regexp.Compile(`^#[0-9a-f]{6}$`)

	eyeColors = map[string]void{
		"amb": member,
		"blu": member,
		"brn": member,
		"gry": member,
		"grn": member,
		"hzl": member,
		"oth": member,
	}

	passportIDRegex, _ = regexp.Compile(`^\d{9}$`)

	reader := bufio.NewReader(os.Stdin)

	numValid := 0
	passportLines := make([]string, 7)
	line := 0

	for {
		// _, err = fmt.Fscanf(reader, "%s\n", &passportLines[line])
		bytes, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		passportLines[line] = string(bytes)

		if passportLines[line] == "" {
			// check if valid
			if isValidPassport2(passportLines, line) {
				numValid++
			}

			line = 0
		} else {
			line++
		}
	}

	if isValidPassport2(passportLines, line) {
		numValid++
	}

	fmt.Println(numValid)
}

func isValidPassport(passportLines []string, noOfLines int) bool {
	resetPassportChecklist()

	for i := 0; i < noOfLines; i++ {
		kvStrs := strings.Split(passportLines[i], " ")

		for _, kvStr := range kvStrs {
			field := kvStr[:3]
			if _, exists := passportFieldChecklist[field]; exists {
				passportFieldChecklist[field]++
			}
		}
	}

	return getChecklistResult()
}

func resetPassportChecklist() {
	passportFieldChecklist["byr"] = 0
	passportFieldChecklist["iyr"] = 0
	passportFieldChecklist["eyr"] = 0
	passportFieldChecklist["hgt"] = 0
	passportFieldChecklist["hcl"] = 0
	passportFieldChecklist["ecl"] = 0
	passportFieldChecklist["pid"] = 0
}

func getChecklistResult() bool {
	p := passportFieldChecklist
	return p["byr"] == 1 && p["iyr"] == 1 && p["eyr"] == 1 && p["hgt"] == 1 && p["hcl"] == 1 && p["ecl"] == 1 && p["pid"] == 1
}

func isValidPassport2(passportLines []string, noOfLines int) bool {
	resetPassportChecklist()
	p := passportFieldChecklist

	for i := 0; i < noOfLines; i++ {
		kvStrs := strings.Split(passportLines[i], " ")

		for _, kvStr := range kvStrs {
			k, v := kvStr[:3], kvStr[4:]

			if k == "byr" {
				byr, err := strconv.Atoi(v)

				if err != nil {
					return false
				}

				if byr < 1920 || byr > 2002 {
					return false
				}

				p[k]++
				continue
			}

			if k == "iyr" {
				iyr, err := strconv.Atoi(v)

				if err != nil {
					return false
				}

				if iyr < 2010 || iyr > 2020 {
					return false
				}

				p[k]++
				continue
			}

			if k == "eyr" {
				eyr, err := strconv.Atoi(v)

				if err != nil {
					return false
				}

				if eyr < 2020 || eyr > 2030 {
					return false
				}

				p[k]++
				continue
			}

			if k == "hgt" {
				unit := v[len(v)-2:]
				hgt, err := strconv.Atoi(v[:len(v)-2])

				if err != nil {
					return false
				}

				if unit == "cm" {
					if hgt < 150 || hgt > 193 {
						return false
					}
				} else if unit == "in" {
					if hgt < 59 || hgt > 76 {
						return false
					}
				} else {
					return false
				}

				p[k]++
				continue
			}

			if k == "hcl" {
				if !colorRegex.Match([]byte(v)) {
					return false
				}

				p[k]++
				continue
			}

			if k == "ecl" {
				if _, exists := eyeColors[v]; !exists {
					return false
				}

				p[k]++
				continue
			}

			if k == "pid" {
				if !passportIDRegex.Match([]byte(v)) {
					return false
				}

				p[k]++
				continue
			}
		}
	}

	return getChecklistResult()
}
