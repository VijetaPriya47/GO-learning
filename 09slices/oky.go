package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of Fruits %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	//ranges are always non-inclusive

	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 344
	highScore[2] = 674
	highScore[3] = 984

	// highScore[4]=349

	highScore = append(highScore, 555, 666, 3210)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	var courses = []string{"reactjs", "swift", "javascript", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
