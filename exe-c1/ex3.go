package main

import "fmt"

func getStars(){

	var rows int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&rows)

	for i:=1;i< rows;i++{
		for j:=1 ;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}