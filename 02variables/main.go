package main

import "fmt"

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
	var smallFloatVal float64 = 255.454667977897
	fmt.Println(smallFloatVal)
	fmt.Printf("variable of type %T \n", smallFloatVal)
}
