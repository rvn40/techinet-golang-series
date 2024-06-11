# Operators in Go

Operators are fundamental components of any programming language, and Go is no exception. They enable developers to perform a variety of operations on data types, ranging from arithmetic calculations to logical comparisons and beyond. Understanding how to effectively utilize these operators is crucial for writing efficient and readable Go code. In this docs, we will delve into the different types of operators available in Go, including arithmetic, comparison, and logical operators, providing clear explanations and practical examples to illustrate their usage. Whether you're a beginner or an experienced programmer, this guide will enhance your understanding of how to leverage operators to manipulate data and control the flow of your Go programs.

## Arithmetic Operators

Arithmetic operators are used to perform basic mathematical operations. Go supports the following arithmetic operators:

- `+` : Addition
- `-` : Subtraction
- `*` : Multiplication
- `/` : Division
- `%` : Modulus

### Example

```go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 3

	sum := a + b
	difference := a - b
	multiplication := a * b
	division := a / b
	modulus := a % b

	fmt.Printf("Sum: %d\n", sum)  
	fmt.Printf("Difference: %d\n", difference) 
	fmt.Printf("Multiplication: %d\n", multiplication)   
	fmt.Printf("Division: %d\n", division) 
	fmt.Printf("Modulus: %d\n", modulus) 
}

```

## Logical Operators

Logical operators are used to combine multiple conditions or expressions and return a boolean value (`true` or `false`). Go supports the following logical operators:

- `&&` : Logical AND
- `||` : Logical OR
- `!` : Logical NOT

### Example

```go
package main

import (
	"fmt"
)

func main() {
	a := true
	b := false
	c := true
	d := false

	// Logical AND
	result := a && c
	fmt.Printf("a && c: %t\n", result) 

	// Logical OR
	result = b || d
	fmt.Printf("b || d: %t\n", result) 

	// Logical NOT
	result = !a
	fmt.Printf("!a: %t\n", result) 

	// Combining Logical AND and OR
	result = (a && c) || (b && d)
	fmt.Printf("(a && c) || (b && d): %t\n", result) 
}
```

## Comparison Operators

Comparison operators are used to compare two values. They return a boolean value (`true` or `false`) based on the comparison. Go supports the following comparison operators:

- `==` : Equal to
- `!=` : Not equal to
- `<` : Less than
- `>` : Greater than
- `<=` : Less than or equal to
- `>=` : Greater than or equal to

### Example

```go
package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	fmt.Printf("a == b: %t\n", a == b)
	fmt.Printf("a != b: %t\n", a != b)
	fmt.Printf("a < b: %t\n", a < b)   
	fmt.Printf("a > b: %t\n", a > b)  
	fmt.Printf("a <= b: %t\n", a <= b)
	fmt.Printf("a >= b: %t\n", a >= b)
}
```

## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E10.png" width="560" height="315">](https://youtu.be/eualxipISOI)