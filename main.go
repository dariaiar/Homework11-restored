package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	var fileName string
	var expression string

	fileName = "listofnumbers.txt"
	expression = `(\d{10}|\(\d{3}\)[ ]?\d{3}[- ]?\d{4}|\d{3}[-. ]\d{3}[-. ]\d{4})`
	findExpression(fileName, expression)

	fileName = "text.txt"
	expression = `[А-Яа-яЁёЇїІіЄєҐґ']*і[А-Яа-яЁёЇїІіЄєҐґ']*і[А-Яа-яЁёЇїІіЄєҐґ']*`
	//Знайти всі слова в яких є дві букви "і"
	findExpression(fileName, expression)
}

func findExpression(fileName string, expression string) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("reading file: %v", err)
		return
	}

	pattern := regexp.MustCompile(expression)
	results := pattern.FindAllString(string(fileContent), -1)
	for i, v := range results {
		fmt.Printf("Match %v: %v\n", i, v)
	}
}
