# LOOP in Go with For Statement

Hey techis! Today, we’re diving into one of the most essential and versatile tools in the Go programming language: the for loop. Whether you’re iterating over slices, maps, or even running infinite loops, the for loop is your go-to structure for all your looping needs in Go. Unlike other languages that have multiple looping constructs like while and do-while, Go keeps it simple and powerful with just one looping keyword: for. This simplicity, however, doesn't mean we're limited; in fact, it opens up a wide array of possibilities and functionalities. So, let’s break down the various ways you can harness the power of the for loop in Go and see some practical examples to get you coding more efficiently and effectively. Here are some examples showcasing the usage of the `for` loop in Go, highlighting its different forms and functionalities.

### 1. Basic `for` Loop
A basic `for` loop is similar to the traditional `for` loop in other languages like C or Java.

```go
package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    fmt.Println("Iteration", i)
  }
	fmt.Println("Finish")
}
```

### 2. `for` Loop as a While Loop
In Go, a `for` loop can also be used as a `while` loop by omitting the initialization and post statements.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println("Iteration", i)
        i++
    }
}
```

### 3. Infinite `for` Loop
An infinite loop can be created by omitting all three components of the `for` loop.

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        fmt.Println("Iteration", i)
        i++
        if i >= 5 {
            break
        }
    }
}
```

### 4. Range `for` Loop with Slices
The `range` keyword is used to iterate over elements of a slice, array, or other data structures.

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### 5. Range `for` Loop with Maps
Using `range` to iterate over key-value pairs in a map.

```go
package main

import "fmt"

func main() {
    myMap := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

### 6. Nested `for` Loop
A nested `for` loop is useful for multi-dimensional data structures like matrices.

```go
package main

import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("Element [%d][%d] = %d\n", i, j, matrix[i][j])
        }
    }
}
```

### 7. `for` Loop with Labels
Labels can be used to control flow in nested loops.

```go
package main

import "fmt"

func main() {
  outerLoop:
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if i == 1 && j == 1 {
        break outerLoop
      }
      fmt.Printf("i: %d, j: %d\n", i, j)
    }
  }
}
```

### 8. Skipping Iterations with `continue`
Using `continue` to skip certain iterations in a loop.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println("Odd number:", i)
    }
}
```

### 9. Looping Over Strings
Using `range` to iterate over the characters of a string.

```go
package main

import "fmt"

func main() {
    str := "Hello, Go!"
    for index, char := range str {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }
}
```

### 10. Using `defer` in a Loop
Deferring function calls within a loop.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        defer fmt.Println("Deferred iteration:", i)
    }
}
```

These examples should give you a broad overview of how versatile the `for` loop is in Go. Each one demonstrates a different aspect of looping and iteration, making it a powerful tool in your Go programming toolkit.

## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E11.png" width="560" height="315">](https://youtu.be/eualxipISOI)