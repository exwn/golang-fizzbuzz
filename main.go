package main

import "fmt"

func main() {

	input := 15
	for i := 1; i <= input; i++ {
		fmt.Println(i)
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		}
	}
}
