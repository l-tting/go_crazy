package main

import "fmt"

func getArea(){

	 var base int
	 var height int

	fmt.Print("Enter base: ")
	fmt.Scanln(&base)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	area := 0.5 * float64(base) * float64(height)
	fmt.Printf("Area of this t is %v m^2\n",area)



}
