package main

import "fmt"

func demoMaps() {
	//nil map
	var m map[string]int
	fmt.Println(m)

	//inbuilt make
	x := make(map[string]int)

	x["t1"] = 1
	x["t2"] = 2

	fmt.Println("map: ", x)

	v1 := x["t1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(x))

	//remove key-vals
	delete(x, "t2")
	fmt.Println("map: ", x)

	//rm all key-vals
	clear(x)
	fmt.Println("map: ", x)

	_, prs := x["t2"]
	fmt.Println("prs:", prs)

	y := make(map[string]int)
	y["item1"] = 123
	y["item2"] = 456
	y["item3"] = 789

	for key, value := range y {
		fmt.Println(key, value)
	}

	for _, value := range y {
		fmt.Println(value)
	}

	for key := range y {
		fmt.Println(key)
	}

}
