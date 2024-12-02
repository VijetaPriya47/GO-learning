package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["js"] = "JAVA"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list", languages)
	fmt.Println("Js", languages["js"])

	delete(languages, "rb")
	fmt.Println("List", languages)

	//list are intresting in golang
	//key, value

	for _, value := range languages {
		fmt.Printf("For key v value is  %v \n", value)
	}

}
