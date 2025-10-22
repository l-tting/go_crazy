package main

import (
	"fmt"
	
)

func demoUserInput(){

	var name string
	fmt.Print("Enter text: ")
	fmt.Scanln(&name)
	fmt.Println("Brian just input:",name)


	var text,text2 string
	fmt.Print("Enter yet anotherr...: ")
	fmt.Scan(&text,&text2)
	fmt.Println(text,text2)
}
