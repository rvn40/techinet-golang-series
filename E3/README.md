# Go Labs(Ngoprek): Initiate Golang Project

In this chapter we will learn how to initialize a project using Go Modules. In this case we can also interprete modules as a golang project.

When creating a project in golang, the go mod command is a method of project or module dependency management that is officially used by Go. Modules are used to initialize a project, as well as manage 3rd party or other libraries used.

## Project Initialization

Here is an example of how to use the go mod init command to initialize a new project.
```
mkdir myproject
cd myproject
go mod init myproject
```

It can be seen in the command above that there is a myproject directory created by mkdir linux command. After entering that directory, the go mod init myproject command is executed. As the result, we have initialized the myproject directory as a Go project.

For the project name, generally it is should named as the directory name, but it would be fine actually to use a different name as well.

Executing the go mod init command would produces a new file named go.mod. This file is would be used by the Go toolchain to mark that the directory where the file is located on the directory folder. So do not delete this file.

## Demo

[<img src="https://storage.googleapis.com/techinet-public/youtube/thumbnails/GolangSeries/E3.png" width="560" height="315">](https://www.youtube.com/embed/Kuo9o4dd1tE)