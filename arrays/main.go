package main

import "fmt"

func main() {
	planets := []string{
		"Merquery",
		"Earth",
		"Mars",
		"Pluton",
		"Jupiter",
		"Neptyn",
		"Venera",
	}

	printSlice(planets)
	changeSlice(planets)

	printSlice(planets)

}

func printSlice(array []string) {

	for i, item := range array {
		fmt.Printf("%v: %v\n", i, item)
	}
	fmt.Println()
}

func changeSlice(slice []string) {

	for i, item := range slice {
		slice[i] = reverse(item)
	}

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
