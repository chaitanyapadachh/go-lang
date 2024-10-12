package main

import "fmt"

func main() {
	i := 1
	//Single Condition Basic For Loop
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	//  Intialise Condition After For Loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// Range over an Integer
	for i := range 3 {
		fmt.Println("range", i)
	}
	// Unlimited Loop --> Exit on Break or Return From the Function
	for {
		fmt.Println("Loop!")
		break
	}
	// Continue Functionality
	for n := range 6 {
		if n%2 == 0 {
			fmt.Println("Divisible by 2", n)
			continue
		}
		fmt.Println("Not Divisible by 2", n)

	}
}
