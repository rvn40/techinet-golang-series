package main

import "fmt"

func main() {

	isTrue := true
	isFalse := false

	fmt.Println(isTrue)
	fmt.Println(isFalse)

	resultAnd := isTrue && isFalse // false
	resultOr := isTrue || isFalse // true
	resultNe := !isTrue // false
	resultNe2 := !isFalse // true

	fmt.Println(resultAnd)
	fmt.Println(resultOr)
	fmt.Println(resultNe)
	fmt.Println(resultNe2)
}