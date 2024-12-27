package main

import "fmt"

func main() {
	var fName string
	fmt.Println("Enter First name")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("Enter Last name")
	fmt.Scanln(&lName)
	fmt.Println("Full name is: ", fName+" "+lName)
}
