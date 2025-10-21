package main

import "fmt"

func demoSwitch(){

	day:= "Monday"

	switch day{
	case "Monday":
		fmt.Println("Monday is here")
	case "Tuesday":
		fmt.Println("Tuesday is here")
	case "Wednesday":
		fmt.Println("Wednesday is here")
	}
}