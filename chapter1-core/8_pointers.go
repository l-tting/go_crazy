package main
import "fmt"

func demoPointers(){
	x := 5
	ptr := &x
	fmt.Println("Memory address is -- ",ptr)
	fmt.Println("Value from x is -- ",*ptr)
}