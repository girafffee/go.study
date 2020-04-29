package main

import (
	"fmt"
	"math/rand"
	"time"
)

var data [100]int

const checkInterval = 10
const minSpeed = 10
const maxSpeed = 100

func main() {
	// initializing random generator with current time
	rand.Seed(time.Now().UnixNano())

	// initializing data array with random speeds
	for i := range data {
		data[i] = rand.Intn(maxSpeed-minSpeed+1) + minSpeed
		// uncomment following to test first checkInterval elements
		// if i < 10 {
		//  data[i] = maxSpeed
		// }

		// uncomment following to test last checkInterval elements
		// if i >= 90 {
		//  data[i] = maxSpeed
		// }
	}

	// here will be your code
	maxDist := 0
	maxSecond := 0
	for thisSecond := 0; thisSecond <= len(data)-checkInterval; thisSecond++ {
		thisDist := 0
		for _, currentDist := range data[thisSecond : thisSecond+checkInterval] {
			thisDist += currentDist
		}
		if thisDist > maxDist {
			maxDist, maxSecond = thisDist, thisSecond
		}
	}

	//fmt.Println(data) // this print is only for deemo purpose
	fmt.Printf("Fastest %d seconds start form %d second with average speed %d\n", checkInterval, maxSecond+1, maxDist/checkInterval)

}
