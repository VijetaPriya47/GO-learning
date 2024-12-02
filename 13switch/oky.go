package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(3) + 1
	fmt.Println("Value", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("2")
		fallthrough  // for case to go for this and the next case also
	default:
		fmt.Println("What you love")
	}
	// days := [string]

}
