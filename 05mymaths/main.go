package main

import (
	// "math/rand"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
