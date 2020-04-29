package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := 30

	//easyTree(size)
	prettyTree(size)
}

func prettyTree(size int) {
	rand.Seed(time.Now().UnixNano())

	for row := size; row > 0; row-- {
		var rowStr string

		// Рисуем половину строки
		for column := 0; column <= size; column++ {
			if column+2 > size && row+(size/2-3) < size {
				rowStr += "|"

			} else if row < column-2 && row+(size/2+5) > size {
				symbolNum := rand.Intn(5)

				switch symbolNum {
				case 0:
					rowStr += "$"
					break
				default:
					rowStr += "/"
				}
			} else if row < column && row+(size/2+5) > size {
				rowStr += "*"

			} else {
				rowStr += " "
			}
		}
		// Отзеркаливаем
		rowStr += reverse(rowStr)
		fmt.Println(rowStr)
	}
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func easyTree(size int) {

	for row := 0; row < size; row++ {
		for column := 0; column <= size*2; column++ {

			if column == size || column+row == size || column-row == size {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}

		}

		fmt.Println()
	}
}
