package main

import "fmt"

func main()  {
	fmt.Println("If else in Golang")
	
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regulate user"
	} else if loginCount>10 {
		result = "Watch Out"
	} else {
		result = "equals to 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println(("Number is odd"))
	}

	if num:=3; num<10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println(("Number is NOT less than 10"))
	}

	// if err != nil {

	// }


	
}