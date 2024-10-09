package main

import "fmt"

func sumOfLetters(input string) int {
	letterValues := buildLetterValues()
	sum := 0
	for _, c := range input {
		sum += findValue(letterValues, fmt.Sprintf("%c", c))
	}

	return sum
}

func buildLetterValues() []string {
	letterValues := make([]string, 0)
	letterValues = append(letterValues, "")
	for pos := 'a'; pos <= 'z'; pos++ {
		letterValues = append(letterValues, fmt.Sprintf("%c", pos))
	}
	return letterValues
}

func findValue(letterValues []string, s string) int {
	for index, c := range letterValues {
		if s == c {
			return index
		}
	}
	return 0
}
