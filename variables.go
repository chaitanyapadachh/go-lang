package main

import "fmt"

func main() {
	var a = "initial" //go supports type inference so when assigning values no need to specify type
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int //But here it is necessary since no assignment of values
	fmt.Println(e)

	f := "apple" //this works only inside functions. This doesnt need a variable type
	fmt.Println(f)

	var (
		x int
		y float64
		z string
		w bool
	)
	fmt.Println(x, y, z, w)
	//When no value is assign default value of 0,"",false are assigned to the variables
}
