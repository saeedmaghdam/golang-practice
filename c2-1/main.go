package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
