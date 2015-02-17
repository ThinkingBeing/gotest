package main

import (
    "fmt"
    "github.com/gotest/stringutil"
    )

func main() {
    msg := "Hello, 世界!"
	fmt.Printf(msg + "\n")
    msg = stringutil.Reverse(msg)
    fmt.Printf(msg)
}
