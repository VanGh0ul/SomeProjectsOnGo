package main

import "fmt"

func main() {

	if 17%3 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 25%5 == 0 {
		fmt.Println("25 is divisible by 5")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "1 digit")
	} else {
		fmt.Println(num, "has multi digit")
	}

}
