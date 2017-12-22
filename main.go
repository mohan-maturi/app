package main

import "fmt"
import "github.com/mohan-maturi/hutil/h"
import mh "github.com/MaturiWorld/hutil/h"

func main() {
    fmt.Println("hello world")
    fmt.Println("\n\nAccessing mohan maturi")
    fmt.Printf("%s\n", h.PrintHello("Anika"))
    fmt.Println("\n\nAccessing maturi world")
    fmt.Printf("%s\n", mh.PrintHello("Trisha"))
}
