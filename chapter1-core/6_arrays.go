package main

import "fmt"

func demoArrays() {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2

	fmt.Println(numbers)
	//uninit. elements defult to default values

	var values = [5]int{10, 20, 30, 40, 50}

	//stf for
	for i := 0; i < len(values); i++ {
		fmt.Printf("Value here at index %v is %v: \n", i, values[i])
	}

	//cleaner for -for range
	for index, value := range values {
		fmt.Println("Index is: ", index, "Value is; ", value)
	}
	
	//ignore index just output val
	for _, value := range numbers {
		fmt.Println(value)
	}

	//ignore values
	for _, value := range numbers {
		fmt.Println(value)
	}

}
