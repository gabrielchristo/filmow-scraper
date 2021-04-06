package main

import "fmt"

func main() {
    fmt.Print("Enter username: ")

    var input string

    fmt.Scanln(&input)

    fmt.Println(input)
}