package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type dataJson struct {
	Rows []RowsP
}
type RowsP struct {
	temp     float64
	pressure int
	humidity int
}

func main() {
	var city string = inputString("Write the name of city: ")

	queryURL := "http://api.openweathermap.org/data/2.5/weather?q={city name}&appid={api key}"

	params := map[string]string{
		"city name": city,
		"api key":   "689c7cf23d4b6f2fed6b6f1f69c441f6",
	}

	// Создаем объект реквеста
	request, err := http.NewRequest("GET", getQueryURLParams(queryURL, params), nil)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]map[string]float64

	json.NewDecoder(response.Body).Decode(&result)

	if _, ok := result["main"]["pressure"]; ok {
		fmt.Printf("Temperature: %.2f ℃\n", toCels(result["main"]["temp"]))
		fmt.Printf("Pressure: %v\n", result["main"]["pressure"])
		fmt.Printf("Humidity: %v%%\n", result["main"]["humidity"])
	} else {
		fmt.Printf("City not founded.\n")
	}

}

func inputString(str string) string {
	fmt.Print(str)

	var line string
	fmt.Scan(&line)

	return line
}

func getQueryURLParams(url string, params map[string]string) string {
	exploedBy := func(c rune) bool {
		return c == '{' || c == '}'
	}
	urlParams := strings.FieldsFunc(url, exploedBy)
	for i, param := range urlParams {
		if val, ok := params[param]; !strings.ContainsAny(param, "./&") && ok {
			urlParams[i] = val
		}
	}
	return strings.Join(urlParams, "")
}

func toCels(kelv float64) float64 {
	return kelv - 273.15
}
