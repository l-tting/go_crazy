package main

import "fmt"

func demoVariables() {

    var a = "initial"
    fmt.Println(a)

	var test string = "Testing Go strings"
	fmt.Println(test)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

    x := "we're back"
    fmt.Printf("Testing %v",x)
}