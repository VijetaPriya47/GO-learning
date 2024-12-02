package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pizza rating:=")

	// // comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!=nil {
		fmt.Println(err)
		// fmt.Println(numrating)

	} else{
		fmt.Println("W",numrating+1)
	}

	presentTime:=time.Now();
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"));

	createDate:=time.Date(2020, time.August,10, 23, 23, 0, 0,time.UTC);
	fmt.Println(createDate);
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"));

}
