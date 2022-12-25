package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in Golang")

	// var myint int = 2
	// var another float64 = 65

	// fmt.Println("sum is :", myint+another)

	// Random Numbers
	// using rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// using crypto

	input, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}
}
