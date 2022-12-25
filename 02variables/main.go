package main

import "fmt"

// jwtToken := 30000 // walrus opertaor is not allowed outside a method
// only val is allowed

const LoginToken string = "dfdsfdsfds789799sdfds987" // Public as first value id caps

func main() {
	var username string = "Pravar"
	fmt.Println(username)
	fmt.Printf("variable of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable of type %T \n", smallVal)

	var smallFloatVal float64 = 255.454667977897
	fmt.Println(smallFloatVal)
	fmt.Printf("variable of type %T \n", smallFloatVal)

	// default values and aliases
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("variable of type %T \n", anotherVariables)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("variable of type %T \n", anotherString)

	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println("LoginToken: ", LoginToken)
}
