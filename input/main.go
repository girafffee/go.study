package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	// for Windows's cmd
	. "github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

/**
For correct colorable use 'log' package, NOT 'fmt'
*/
func init() {
	log.SetOutput(colorable.NewColorableStdout()) // or NewColorableStderr()
}

type stringSlice []string
type indexesChar map[int]int
type limitingChar rune

func main() {
	// Saving opening and closing char
	open, close := inputBorderChars()

	// Read test string from console
	line := inputString(open, close)

	// for chunks of text separated by bounding characters
	var couples stringSlice
	var err error

	// founding all the chuncks or get the error
	couples, err = couples.getBrackets(line, open, close)

	couples.printResult(open, close, err)
}

func inputBorderChars() (limitingChar, limitingChar) {
	open := readRune(fmt.Sprintf("Input opening char: "))
	close := readRune(fmt.Sprintf("Input closing char: "))

	return limitingChar(open), limitingChar(close)
}

func readRune(info string) rune {
	fmt.Print(info)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return char
}

func inputString(open, close limitingChar) string {
	fmt.Printf("Enter the string to be checked with %c and %c: ", rune(open), rune(close))

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()

	return line
}

func (coupleBrakets stringSlice) printResult(open, close limitingChar, err error) {
	if err == nil && len(coupleBrakets) > 0 {
		fmt.Printf("couples: " + getSliceUl(coupleBrakets))
		log.Println(Green(Bold("String is correct")))
	} else if err != nil {
		log.Println(Red(err))
	} else {
		log.Println(Bold(Cyan("\n\tThese characters were not found in the text: " + string(open) + " " + string(close) +
			"\n\tEverything is great.")))
	}
}

/*
A recursive method that looks for a string
with a substring from one(!) pair of bounding characters,
and returns the remainder to re-check for bounding characters
*/
func (coupleBrakets stringSlice) getBrackets(str string, open limitingChar, close limitingChar) (stringSlice, error) {

	var mainErr error
	if strings.ContainsAny(str, string(open)+string(close)) {
		couple, excessString, err := openBracket(str, open, close)
		if err == nil {
			coupleBrakets = append(coupleBrakets, couple)
		} else {
			return coupleBrakets, err
		}

		if err == nil && strings.ContainsAny(excessString, string(open)+string(close)) {
			coupleBrakets, err = coupleBrakets.getBrackets(excessString, open, close)
		}
		mainErr = err
	}

	return coupleBrakets, mainErr
}

func openBracket(str string, open limitingChar, close limitingChar) (string, string, error) {

	indexLastOpen := strings.LastIndex(str, string(open))
	if indexLastOpen == -1 {
		errorInfo := fmt.Sprintf("Not founded the open bracket[%c]\nIn this line: %v", open, str)
		return "", "", errors.New(errorInfo)
	}

	var stringBeforeLastBracket string
	var stringAfterLastBracket string
	for position, char := range str {

		if position < indexLastOpen {
			stringBeforeLastBracket += string(char)
		} else {
			stringAfterLastBracket += string(char)
		}
	}

	coupleBrakets, excessString, err := closeBracket(stringAfterLastBracket, close)
	stringBeforeLastBracket += excessString

	return coupleBrakets, stringBeforeLastBracket, err
}

func closeBracket(str string, close limitingChar) (string, string, error) {

	indexFirstClose := strings.Index(str, string(close))
	if indexFirstClose == -1 {
		errorInfo := fmt.Sprintf("Not founded the close bracket[%c]\nIn this line: %v", close, str)
		return "", "", errors.New(errorInfo)
	}

	var stringBeforeFirstBracket string
	var stringAfterFirstBracket string
	for position, char := range str {

		if position <= indexFirstClose {
			stringBeforeFirstBracket += string(char)
		} else {
			stringAfterFirstBracket += string(char)
		}
	}

	return stringBeforeFirstBracket, stringAfterFirstBracket, nil
}

func colorBrackets(line string, open limitingChar, close limitingChar) string {

	indexOpen := strings.LastIndex(line, string(open))
	indexClose := indexClose(line, indexOpen, close)

	newLine := colorSymbol(line, indexOpen)
	newLine = colorSymbol(line, indexClose)

	return newLine
}

func indexClose(line string, indexOpen int, close limitingChar) int {
	for index, char := range line {
		if index > indexOpen && char == rune(close) {
			return index
		}
	}
	return -1
}

func colorSymbol(line string, indexChar int) string {
	var newLine string
	for index, char := range line {
		if index == indexChar {
			newLine += fmt.Sprint(Cyan(string(char)))
		} else {
			newLine += string(char)
		}
	}

	return newLine
}

func getSliceUl(slice stringSlice) string {

	str := "\n"
	for i, item := range slice {
		str += fmt.Sprintf("%v: %v\n", i, item)
	}
	return str + "\n"
}
