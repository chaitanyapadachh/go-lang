package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}
	if 4%2 == 0 {
		fmt.Println("4 is Even")
	} else {
		fmt.Println("4 is Odd")
	}
	num := 5
	var xyz = 5
	fmt.Println(xyz)
	if num%2 == 0 {
		fmt.Println(num, "Number is Even")
	} else if !(num%2 == 0) {
		fmt.Println(num, "Number is Not Divisible by 2")
	}
}
