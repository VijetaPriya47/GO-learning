package main

import "fmt"

///function into classes = methods
//structs => classes

func main() {
	fmt.Println("Methods in golang")
	//no inheritance in golang: No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details: %+v \n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus();
	hitesh.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int  ///not importable bacuse th first letter is small
}


func (u User) GetStatus(){
	fmt. Println("Is User active: ",u.Status)
}

//Call by copy is going in functions, 
//We should use POINTERS to change the properties of the user
func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

