package main

import "fmt"

func main() {
    // var num1 complex64 = 9 + 8i      
    // var num2 complex64 = 5.5 + 2.1i

		num1 := complex(9, 8)
    num2 := complex(5.5, 2.1)

    // Performing arithmetic operations
    sum := num1 + num2
    difference := num1 - num2
    multiply := num1 * num2

    // Outputting the results
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
    fmt.Println("Multiplier:", multiply)
}
