package main

import "fmt"

const LoginToken = "asdf" // Public Variable

func main() {
	var username string = "Jatan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 69.696969696969
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVal int
	fmt.Println(anotherVal) // output: 0 
	fmt.Printf("variable is of type: %T \n", anotherVal) 

	var anotherValStr string
	fmt.Println(anotherValStr) // output: blank i.e. nothing will be printed.
	fmt.Printf("variable is of type: %T \n", anotherValStr)

	// implicit type
	var website = "jatanshah.com"
	fmt.Println(website) 

	// no var style
	numberOfUsers := 200000
	fmt.Println(numberOfUsers)

	// Using Public Variable
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
