package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	b = 7

	upperAlphabetStart = rune('A')
	lowerAlphabetStart = rune('a')
	alphabetLength     = 26

	phraseMarryMe   = "Will you marry me?"
	phraseOk        = "Yeah, okay!"
	phraseGreat     = "Great!"
	phraseBeFriends = "Let's be friends."
	phrasePity      = "What a pity!"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	g, p, err := parseInputValues(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing g, p values")
		return
	}

	fmt.Println("OK")

	scanner.Scan()
	_, A, err := parseInputValue(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing A value")
		return
	}

	B := calculatePow(g, b, p)
	fmt.Println(fmt.Sprintf("B is %d", B))

	s := calculatePow(A, b, p)

	fmt.Println(encryptMessage(phraseMarryMe, s))

	scanner.Scan()
	answer := decryptMessage(scanner.Text(), s)

	if answer == phraseOk {
		fmt.Println(encryptMessage(phraseGreat, s))
	}

	if answer == phraseBeFriends {
		fmt.Println(encryptMessage(phrasePity, s))
	}
}

func parseInputValues(input string) (int, int, error) {
	var g int
	var p int

	parts := strings.Split(input, "and")
	for _, part := range parts {
		varName, varValue, err := parseInputValue(strings.TrimSpace(part))
		if err != nil {
			return 0, 0, err
		}

		if varName == "g" {
			g = varValue
		} else if varName == "p" {
			p = varValue
		}
	}

	return g, p, nil
}

func parseInputValue(input string) (string, int, error) {
	vars := strings.Split(strings.TrimSpace(input), "is")

	value, err := strconv.Atoi(strings.TrimSpace(vars[1]))
	if err != nil {
		return "", 0, err
	}

	return strings.TrimSpace(vars[0]), value, nil
}

func calculatePow(g, pow, p int) int {
	result := 1
	for i := 0; i < pow; i++ {
		result = (result * g) % p
	}
	return result
}

func encryptMessage(message string, key int) string {
	var result string
	for _, c := range message {
		result += string(encryptRune(c, key))
	}

	return result
}

func encryptRune(runeValue rune, key int) rune {
	if unicode.IsSpace(runeValue) || unicode.IsPunct(runeValue) {
		return runeValue
	}

	alphabetStart := lowerAlphabetStart
	if unicode.IsUpper(runeValue) {
		alphabetStart = upperAlphabetStart
	}

	letterIdx := int(runeValue - alphabetStart)
	offset := key % alphabetLength
	newIdx := (letterIdx + offset) % alphabetLength
	newLetterRune := rune(newIdx) + alphabetStart

	return newLetterRune
}

func decryptMessage(message string, key int) string {
	var result string
	for _, c := range message {
		result += string(decryptRune(c, key))
	}

	return result
}

func decryptRune(runeValue rune, key int) rune {
	if unicode.IsSpace(runeValue) || unicode.IsPunct(runeValue) {
		return runeValue
	}

	alphabetStart := lowerAlphabetStart
	if unicode.IsUpper(runeValue) {
		alphabetStart = upperAlphabetStart
	}

	letterIdx := int(runeValue - alphabetStart)
	offset := key % alphabetLength
	guessIdx := letterIdx - offset
	if guessIdx == 0 {
		return alphabetStart
	}

	if guessIdx < 0 {
		newIdx := alphabetLength - (-guessIdx % alphabetLength)
		newLetterRune := rune(newIdx) + alphabetStart
		return newLetterRune
	}

	return rune(guessIdx) + alphabetStart
}
