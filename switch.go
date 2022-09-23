package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("I switch ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one\n")
	case 2:
		fmt.Println("two\n")
	case 3:
		fmt.Println("three\n")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend\n")
	default:
		fmt.Println("Its a weekday\n")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon \n")
	default:
		fmt.Println("Its after noon \n")
	}

	whatAmIType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Its bool\n")

		case int:
			fmt.Println("Its int\n")

		case float64:
			fmt.Println("Its float\n")

		case string:
			fmt.Println("Its string\n")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmIType(true)
	whatAmIType(1)
	whatAmIType("hola")
	whatAmIType(false)
	whatAmIType(1.02)
	whatAmIType("thanks")

}
