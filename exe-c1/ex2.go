package main 

import "fmt"

func checkEven(){
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number%2==0{
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}
}