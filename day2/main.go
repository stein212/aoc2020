package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// change function used in checkPasswords to answer different problem
	checkPasswords()
}

func checkPassword1(min, max int, letter rune, password string) bool {
	counter := 0
	for _, c := range password {
		if c == letter {
			counter++
		}
	}

	return min <= counter && counter <= max
}

func checkPassword2(min, max int, letter rune, password string) bool {
	bletter := byte(letter)
	return (password[min-1] == bletter || password[max-1] == bletter) && (password[min-1] != password[max-1])
}

func checkPasswords() {
	reader := bufio.NewReader(os.Stdin)

	numValid := 0

	for {
		var num1, num2 int
		var letter rune
		var password string
		_, err := fmt.Fscanf(reader, "%d-%d %c: %s\n", &num1, &num2, &letter, &password)

		if err != nil {
			break
		}

		// if checkPassword1(num1, num2, letter, password) {
		// 	numValid++
		// }

		if checkPassword2(num1, num2, letter, password) {
			numValid++
		}
	}

	fmt.Println(numValid)
}
