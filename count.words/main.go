package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"
)

// TopPositions ...
const TopPositions = 5

type textSeeder interface {
	getText() string
	countWords() map[string]int
	getLongestSentence() string
	getTotalWords() int
	sortWordsMap(int) []string
}

// Data ...
type Data struct {
	text       string
	words      map[string]int
	topWords   []string
	sentences  []string
	totalWords int
}

/*
	You can add your custom Data[Type]
	to get text in your own way
*/

// DataAPI ...
type DataAPI struct {
	Data
	urlAPI string
}

// DataFile ...
type DataFile struct {
	Data
	filename string
}

func main() {
	/*
		You can change a `paras` parameter.
		He is responsible for the number of paragraphs in the text.
	*/
	apidata := newDataAPI("https://baconipsum.com/api/?type=all-meat&paras=3")
	fmt.Println(apidata)

	/*
		A file with this name should be in the same folder
		with the program and run from the directory of the program itself
		(ex. - `go run main.go`)
	*/
	filedata := newDataFile("text.txt")
	fmt.Println(filedata)
}

func (d *DataAPI) String() string {
	str := "\n"
	str += fmt.Sprintf("\tRandom text from API\n\n")
	//str += fmt.Sprintf("Source text: %s\n\n", text(d))
	str += fmt.Sprintf("Word count: %d\n\n", totalWords(d))
	str += fmt.Sprintf("Longest sentence: %s\n\n", longestSentence(d))
	str += fmt.Sprintf("Top words:\n")
	for pos, word := range topWords(d, TopPositions) {
		str += fmt.Sprintf("%d - %s\n", pos+1, word)
	}
	return str
}

func (d *DataFile) String() string {
	str := "\n"
	str += fmt.Sprintf("\tText from FILE\n\n")
	//str += fmt.Sprintf("Source text: %s\n\n", text(d))
	str += fmt.Sprintf("Word count: %d\n\n", totalWords(d))
	str += fmt.Sprintf("Longest sentence: %s\n\n", longestSentence(d))
	str += fmt.Sprintf("Top words:\n")
	for pos, word := range topWords(d, TopPositions) {
		str += fmt.Sprintf("%d - %s\n", pos+1, word)
	}
	return str
}

func (d *Data) sortWordsMap(positions int) []string {
	for i := 0; i < positions; i++ {
		d.MaxIntMap()
	}
	return d.topWords
}

func topWords(s textSeeder, positions int) []string {
	return s.sortWordsMap(positions)
}

// MaxIntMap ...
func (d *Data) MaxIntMap() {
	max := 0
	var maxWord string
	for word, count := range d.words {
		if count > max {
			max = count
			maxWord = word
		}
	}
	delete(d.words, maxWord)
	d.topWords = append(d.topWords, maxWord)
}

func (d *Data) countWords() map[string]int {
	if len(d.words) > 0 {
		return d.words
	}

	exploedWords := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(strings.ToLower(d.text), exploedWords)

	mapWords := make(map[string]int)
	for _, word := range words {
		d.totalWords++
		mapWords[word]++
	}
	d.words = mapWords
	return d.words
}

func (d *Data) countSentences() []string {
	if len(d.sentences) > 0 {
		return d.sentences
	}
	exploedSentences := func(c rune) bool { return c == '.' || c == '!' || c == '?' }
	for _, sent := range strings.FieldsFunc(d.text, exploedSentences) {
		d.sentences = append(d.sentences, strings.TrimSpace(sent))
	}
	return d.sentences
}

func (d *Data) getLongestSentence() string {
	if len(d.sentences) == 0 {
		d.countSentences()
	}

	maxLength := 0
	maxPos := 0
	for pos, sent := range d.sentences {
		if length := len(sent); length > maxLength {
			maxLength = length
			maxPos = pos
		}
	}
	return d.sentences[maxPos]
}

func (d *Data) getTotalWords() int {
	d.countWords()
	return d.totalWords
}

func text(s textSeeder) string {
	return s.getText()
}

func words(s textSeeder) map[string]int {
	return s.countWords()
}

func longestSentence(s textSeeder) string {
	return s.getLongestSentence()
}
func totalWords(s textSeeder) int {
	return s.getTotalWords()
}

/*
Reads a file
*/
func (d *DataFile) getText() string {
	data, err := ioutil.ReadFile(d.filename)
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}
	d.text = string(data)
	return d.text
}

func (d *DataAPI) getText() string {
	if len(d.text) > 0 {
		return d.text
	}
	// Создаем объект реквеста
	request, err := http.NewRequest("GET", d.urlAPI, nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result []string
	json.NewDecoder(response.Body).Decode(&result)

	d.text = strings.Join(result, " ")
	return d.text
}

func newDataFile(file string) *DataFile {
	d := DataFile{filename: file}
	d.getText()
	return &d
}

func newDataAPI(url string) *DataAPI {
	d := DataAPI{urlAPI: url}
	d.getText()
	return &d
}
