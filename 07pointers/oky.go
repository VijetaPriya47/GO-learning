package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	// var ptr *int
	// fmt.Println("value of pointer",ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("val", ptr)
	fmt.Println("val", *ptr)

	*ptr = *ptr *2;
	fmt.Print(myNumber);
	
}
