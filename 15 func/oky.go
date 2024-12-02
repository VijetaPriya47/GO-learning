package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang ");

	grreeter()

	// func GreeterTwo() {
	// 	fmt.Println("Another moethod")
	// }
	////YOu can nor declare functions insideA FUNCTION 

	result := adder(3,5)

	fmt.Println("Result is: ", result)

	proRes, proMessage := proAdder(2,5,8,7) 
	fmt.Println("Result is: ", proRes)
	fmt.Println("ProMessage is: ", proMessage)

}
//function signatures (call and return intializer)

func adder(valOne int, valTwo int) int{
	return valOne+valTwo
}

func proAdder(values ...int) (int, string) {
	total:= 0

	for _, val := range values{
		total += val
	}

	return total, "Hi Pro result function";
}


func grreeter() {
	fmt.Println("Namastey from golang")
}

