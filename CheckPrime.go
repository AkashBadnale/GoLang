package main

import "fmt"

func main() {

    var input int
    var result bool // default is false

    fmt.Print("Enter the Number to find the Prime Numbers = ")
    fmt.Scan(&input)

    for i := 2; i < input; i++ {
        if (input%i == 0) {
            result = true
            break
        }
    }

    if (!result && input != 1) {
        fmt.Println(input, " is a Prime Number")
    } else {
        fmt.Println(input, " is Not a Prime Number")
    }
}
