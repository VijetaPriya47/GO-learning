package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang: No super or parent

	hitesh := User{"Hitesh", "hitesh@gmail.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details: %+v \n", hitesh)
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
