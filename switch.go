package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("The Weekend!")
	default:
		fmt.Println("Weekdays")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before Noooooon")
	default:
		fmt.Println("Post Nooooooooon!")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Booleans")
		case int:
			fmt.Println("Integers")
		default:
			fmt.Println("Don't know type for %T\n", t)
		}

	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
