package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	citizen bool
}

// struct methods with receiver
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.name)
}

// method -> copy receiver -> param
func (p Person) cook(meal string) {
	fmt.Println(p.name, "can cook ", meal)

}

//pointer receiver
func (p *Person) ChangeName(){
	p.name = "Ethan"
}

func demoStruct() {

	//struct variable - p is of type Person
	var x Person
	x.name = "Brian"
	// p.age = 26
	// p.address = "123 Kimathi St"

	// short initialization
	s := Person{name: "Brian", age: 21, address: "Nairobi", citizen: true}

	// fmt.Println(p)
	fmt.Println(s)

	x.ChangeName()
	x.Greet()
	x.cook("pasta")


}
