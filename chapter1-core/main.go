package main

import (
	"fmt"
	"go_dev/tester"
)




func main(){
	demoVariables()
	tester.Tester1()

	sum := tester.Tester2()
	fmt.Println(sum)

	// demoFor()
	// demoifElse()
	// demoSwitch()
	// demoUserInput()
	// demoArrays()
	// demoSlices()
	// demoPointers()
	// demoStruct()
	// demoEmbeddedStructs()
	// demoMaps()
	// sum := demoVariadics(1,2,3,4,5)
	// fmt.Printf("Sum from variadic is %v\n",sum)

	// message := demoVariad("Mike","KImberly","Joe")
	// fmt.Printf("We get to say: %q\n",message)

	// increment := demoClosure()
	// fmt.Printf("Value of count is %d",increment())
	// fmt.Printf("Value of count is %d\n",increment())

	// factorial := Recursor(5)
	// fmt.Println(factorial)
	

}