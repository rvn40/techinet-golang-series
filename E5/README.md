# Go Labs(Ngoprek): Variables

As many people know Golang has a Go compiler that is initially written in C. So, go adopts two types of variable writing, namely with and without data types. Both methods are valid and have the same purpose, the difference is only in the way of writing.

## Variable with Data Types

Go has pretty strict concept when it comes to writing variables, This concept is called manifest typing. Manifest typing refers to the practice of explicitly declaring the type of a variable or expression in your code. When you declare a variable, you provide its type explicitly. This approach ensures that the type of a variable is known at compile time and helps catch type-related errors early in the development process.

Here is the example of how to create a variable whose data type must be written:
```
package main

import "fmt"

func main() {
  var firstName string = "James"
  
  fmt.Println(firstname)

  var b, c int = 1, 2
  
  fmt.Println(b, c)
}
```

At the string example The "var" keyword above is used for variable declarations, followed by variable name (firstname) and the data type(string) and equals operator with value inside the double-quotes. For the "int" example I want to show you how multiple values could be defined at the same line as long as the values share the same data types(int).

## Variable without Data Types
Instead of manifest typing, GO has another concept called type inference. It means we can declare variable without explicitly write the data types. Go will define the data types based on the value. Type inference is a feature that allows the compiler to determine the type of a variable or expression based on its value. This means that you don't always need to explicitly declare the type of a variable; the compiler can infer it from the value assigned to it.

```
package main

import "fmt"

func main() {
  firstName := "James"
  
  fmt.Println(firstname)
}
```
The variable firstname is declared using the type inference method. The data type is not written at the time of the declaration. So by using this method, the operand = must be replaced with := and the var keyword is omitted.

The firstname data type will automatically be determined according to its value. If the value is a string, the variable's data type is string. In the example above, the value is "James" which is a string.

The operand := only use at the first time varable initialized. If there is an update of the value in the next line, use the operand = for update the values. 
```
firstName := "James"
firstName = "Danny"
firstName = "Lucy"
```

## Blank Identifier

In Go, the blank identifier (often represented as _) is a special identifier used as a placeholder for a variable or value that you wish to ignore or discard. It's commonly used in situations where you want to call a function or perform an operation that returns a value, but you're not interested in using that value in your code.

There are commonly three scenarios where the blank identifier is useful:
  1. Ignoring Function Return Values

    ```
    package main

    import (
        "fmt"
    )

    func getValues() (int, int) {
        return 10, 20
    }

    func main() {
        a, _ := getValues()
        fmt.Println("Value of a:", a)
    }

    ```

  2. Importing Packages for Side Effects

    ```
    package main

    import (
        _ "github.com/example/somepackage"
    )

    func main() {
        // This program imports "somepackage" just for its side effects (initialization, registration, etc.)
    }

    ```
  3. Looping Through Arrays/Slices Without Using Index or Element

    ```
    package main

    func main() {
        numbers := []int{1, 2, 3, 4, 5}
        for _, _ = range numbers {
            // Do something without using index or element
        }
    }
    
    ```

Using the blank identifier can improve code readability by making it clear that you're intentionally ignoring certain values. It also helps prevent unused variable warnings from the Go compiler. However, excessive use of the blank identifier should be avoided, as it can make code harder to understand for other developers.

It's important to note that Go requires that all declared variables must be used. If you declare a variable but don't use it, the code will not compile. The blank identifier is a way to acknowledge that a variable is intentionally not used.

At least there are 2 more variants regarding writing variables in go, namely writing variables for pointers and writing variables for make.

However, judging from the complexity of the discussion of these two variables, it seems that it would be better discussed later when discussing the topic.

## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E5.png" width="560" height="315">](https://www.youtube.com/embed/zu1hgHPo3H4?si=XwpKc0RaWApFNEi3)