package main

import "fmt"

func demoStruct(){
	type Person struct{
		name string
		age int
		address string
		citizen bool
	}
	//struct variable - p is of type Person
	var p Person
	p.name = "Brian"
	p.age = 26
	p.address = "123 Kimathi St"

	fmt.Println(p)
}