# Go Labs(Ngoprek): Data type Part 1 - String and Boolean

In this session, we will explore the essential concepts of Strings and Booleans. This docs is designed to provide a comprehensive understanding of these fundamental data types, which are crucial for text manipulation and logical operations in Go.

We will begin by examining string handling, including various methods for manipulating and formatting text. Following this, we will delve into Boolean logic, covering the principles of true and false values and their practical applications in conditional statements and logical operations.

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

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E6.png" width="560" height="315">](https://www.youtube.com/embed/smpPdrmOslE?si=ARGj9zcPVjdZC1co)
