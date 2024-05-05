package main

import "fmt"

func main() {
	message := "Hello para techi!!"

	fmt.Println("Message: ", message)

	length := len(message)

	fmt.Println("Length: ", length)

	firstChar := message[0]

	fmt.Println("First Character: ", string(firstChar))

}