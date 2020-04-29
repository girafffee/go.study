package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// MinSpeed is minimal speed fo random generate
const MinSpeed = 10

// MaxSpeed is maximal speed fo random generate
const MaxSpeed = 100

// Amount ...
const Amount = 10

// Size of array
const Size = 100

func main() {
	arrayFloat, sumArray := generateRandSlice(Size)
	fmt.Printf("Average speed for %v seconds is: %.2f\n", Size, average(sumArray, len(arrayFloat)))

	maxAverageSpeed, startPosition := findAverageAtDistance(arrayFloat)

	//writeToFile("D:\\Projects\\Go\\src\\speed\\speed.data", convertFloatArray(arrayFloat))

	fmt.Printf("Max average speed for %v seconds is: %.2f\n", Amount, maxAverageSpeed)
	fmt.Printf("From %v to %v\n", startPosition+1, startPosition+Amount)
}

/*
finds the largest arithmetic mean
in a slice with a specified length (Amount)
*/
func findAverageAtDistance(array []float64) (float64, int) {
	var maxAverageSpeed float64
	var startPosition int
	for pos := range array {
		if end := len(array) - Amount; pos <= end {
			var totalSpeed float64
			for i := pos; i < (pos + Amount); i++ {
				totalSpeed += array[i]
			}

			averageSpeed := average(totalSpeed, Amount)
			if averageSpeed > maxAverageSpeed {
				maxAverageSpeed = averageSpeed
				startPosition = pos
			}
		}
	}

	return maxAverageSpeed, startPosition
}

func average(sum float64, amount int) float64 {
	return sum / float64(amount)
}

/*
Get slice of random numbers
in float64
*/
func generateRandSlice(size int) ([]float64, float64) {
	// initializing random generator with current time
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, size)
	// initializing data array with random speeds
	var totalSum float64
	for i := range data {
		data[i] = MinSpeed + rand.Float64()*(MaxSpeed-MinSpeed)
		totalSum += data[i]
	}
	return data, totalSum
}

/*
Reads a file line by line
*/
func readDataFromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}
	dataSlice := strings.Split(string(data), "\r\n")
	// Если чтение данных прошло успено
	// выводим их в консоль
	return dataSlice
}

/*
Write string in file (overwrites)
*/
func writeToFile(filename, text string) {
	err := ioutil.WriteFile(filename, []byte(text), 0777)
	// Обработка ошибки
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

/*
Turns an array with numbers
into a string separated by "\r\n"
*/
func convertFloatArray(arr []float64) string {
	var floatlstrings []string
	for _, item := range arr {
		// Numeration
		// fmt.Sprintf("%v: ", pos+1)
		b64 := []byte("")
		b64 = strconv.AppendFloat(b64, item, 'E', -1, 64)
		floatlstrings = append(floatlstrings, string(b64))
	}

	return strings.Join(floatlstrings, "\r\n")
}
