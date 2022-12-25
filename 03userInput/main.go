package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcom to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate for our pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
