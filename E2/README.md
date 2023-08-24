# Go Labs(Ngoprek): First Program and Values 

In this chapter we will describe the creation of the first program, a small application that displays Hello world text. Below is an example of a simple program code for bringing Hello world text to the output command prompt screen.

```
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

After the code is copied, open the terminal (or CMD for Windows users), then go to the project directory, then run the program using the command go run filename.go

As a result, it's supposed to appear hello world on the console screen. If hello world appears on your console or terminal then Congratulations! You have successfully created the Go program!

## Keyword package

Each program file must have a package. Each project must have at least one file with the name package main. The file that has a play package will be executed the first time the program is run.

How to define a package using keyword package can be seen in the code snippet above.

## Keyword import

The keyword "import" is used to do "import", or to insert another package into a program file, so that the contents of the package that is in " import" can be used.

In this case, the "fmt" package is one of the built-in packages provided by Go, containing many functions for input output purposes related to text.

## Function main() and fmt.Println()

In a project there must be a program file containing a function called main(). The function must be in a file whose package is called main.

The fmt.Println() function is used to generate text to the screen (terminal atau CMD). In the first program we've created, this function generates Hello world.

The keyword writing scheme fmt.Println() can also be seen in the example above.

The fmt.Println() function is in the fmt package, so to use it you need to import the package first.

## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E2.png" width="560" height="315">](https://www.youtube.com/embed/Kuo9o4dd1tE)