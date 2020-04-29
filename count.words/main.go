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

type textSeeder interface {
	getText() string
	countWords() map[string]int
}

// Data ...
type Data struct {
	text  string
	words map[string]int
}

// DataAPI ...
type DataAPI struct {
	Data
	urlAPI string
}

func main() {
	apidata := newDataAPI("https://baconipsum.com/api/?type=all-meat&paras=3")

	fmt.Println(text(apidata))

	fmt.Println(words(apidata))
}

/*
Reads a file
*/
func readDataFromFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}
	// Если чтение данных прошло успено
	// выводим их в консоль
	return string(data)
}

func (d *DataAPI) countWords() map[string]int {
	exploedBy := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(d.text, exploedBy)
	mapWords := make(map[string]int)

	for _, word := range words {
		mapWords[word]++
	}
	d.words = mapWords
	return d.words
}

func text(s textSeeder) string {
	return s.getText()
}

func words(s textSeeder) map[string]int {
	return s.countWords()
}

func (d *DataAPI) getText() string {
	// Создаем объект реквеста
	request, err := http.NewRequest("GET", d.urlAPI, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//timeout := time.Duration(1 * time.Second)
	client := http.Client{
		//Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result []string
	json.NewDecoder(response.Body).Decode(&result)

	d.text = strings.ToLower(strings.Join(result, " "))
	return d.text
}

func newDataAPI(url string) *DataAPI {
	d := DataAPI{urlAPI: url}
	return &d
}
