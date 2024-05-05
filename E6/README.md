# Go Labs(Ngoprek): Data type Part 1 - String and Boolean

As many people know Golang has a Go compiler that is initially written in C. So, go adopts two types of variable writing, namely with and without data types. Both methods are valid and have the same purpose, the difference is only in the way of writing.

## Data Type String

Strings are like the most fundamental data type in any programming language, and Go is no exception. They are your go-to for handling text operation, whether it is a simple greeting or a very long Novel like text. Let's keep it simple for now.

```
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
```

In this example, we're declaring a string variable message with greeting to techis. We then print it out and even find out its length.


## Data Type Boolean

Boolean is data type which consists only with 2 values, true and false. They are like the yes/no, true/false, on/off switches that guide our program's logic. Let's check them out in action. 

```
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

```

In this snippet, we have declare a couple of boolean variables, isTrue and isFalse. Then, we are perform some logical operations using these booleans. We do some AND, OR, and even NOT operations.


## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E6.png" width="560" height="315">](https://www.youtube.com/embed/zu1hgHPo3H4?si=XwpKc0RaWApFNEi3)
