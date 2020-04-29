package main

import "fmt"

var (
	Black   = "\033[1;30m-\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m*\033[0m"
	Yellow  = "\033[1;33mIIII\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

func main() {
	var n int = 15

	for h := 0; h < n; h = h + 1 {
		fmt.Printf("%v", h)

		for k := 20; k > h; k = k - 4 {
			fmt.Print(" ")
		}
		for i := 0; i < h; {
			if i%5 == 0 {
				fmt.Print(Black)
			} else {
				fmt.Print(Green)
			}

			i = i + 1

		}
	}
}
