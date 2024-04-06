package main

import (
	"fmt"
)

func main() {

	fmt.Println("loop and condition")
	loginCount := 12
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 loged in"
	}
	fmt.Println(result)

	// important for api
	if num := 3; num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// loop on slice
	days := []int{4, 1, 2, 7, 4, 9}
	for i := 0; i < len(days); i++ {
		fmt.Print(days[i], " ")
	}
	fmt.Println()
	for d := range days {
		fmt.Print(days[d], " ")
	}
	fmt.Println()
	// for each comma ok syntax
	for _, day := range days {
		fmt.Print(day)
	}
	fmt.Println()
	number := 1

	// while loop
	for number < 4 {
		if number == 1 {
			goto lco
		}
		if number == 2 {
			break
		}
		fmt.Println(number)
		number++
	}

	// goto statement
    
	lco:
	fmt.Println("golang is best language")
	fmt.Println("i will become best in both golang and js")


}
