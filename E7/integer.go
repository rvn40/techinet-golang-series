package main

import "fmt"

func main() {

	var myAge uint8 = 28
	var myBrotherAge uint8 = 20

	TotalAge := myAge + myBrotherAge

	fmt.Println("Result: ", TotalAge)

	
	var minusNum int32 = -28
	var positiveNum int = 20

	result := minusNum + int32(positiveNum)

	fmt.Println("Result: ", result)
}