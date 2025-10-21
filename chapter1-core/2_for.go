package main

import "fmt"

func demoFor(){
	//standard for
	for i:=0 ;i<5 ;i++{
		fmt.Println("Count is :",i)
	}

	//while style 
	j:=0
	for j < 5{
		fmt.Println("Count 2 is: ",j)
		j++
	}

	//infinite
	for {
		fmt.Println("Infinitely inifinte....")
		break // break manually
	}

}

