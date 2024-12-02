package main

import "fmt"

func main() {

	fmt.Println("Welcome")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist", fruitList)
	fmt.Println("Fruitlist", len(fruitList))

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println("Veg list",vegList);

	

}
