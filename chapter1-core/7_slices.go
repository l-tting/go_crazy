package main

import (
	"fmt"
	"reflect"
)

func demoSlices(){
	//create a slice method 1
	arr := [5]int{10,20,30,40,50}
	slice := arr[1:4]
	fmt.Println(reflect.TypeOf(slice))

	//create a slice method 2
	s := []int{1,2,3}
	fmt.Println(reflect.TypeOf(s))

	//using make() -method 3
	s = make([]int, 12,)
	fmt.Println(s)
	s = append(s,10,20,30 )
	fmt.Println(s)

}