package main

import "fmt"

func demoifElse(){
	//standard if
	num := 2
	if num%2 ==0{
		fmt.Println("Is even")
	}else{
		fmt.Println("Is odd")
	}


	//if with a short
	if x := 5;x>10{
		fmt.Println("Greater than 10")
	}else{
		fmt.Println("Defintely not")
	}

}