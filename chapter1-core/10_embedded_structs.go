package main

import "fmt"


type Employee struct{
	name string
	employee_id int
}

type Manager struct{
	Employee 
	level string
}

func (x Employee) Work(){
	fmt.Println(x.name ,"works 80hrs a week")
}


func demoEmbeddedStructs(){
	e1 := Employee{name:"Jane Doe",employee_id: 23418}

	m1 := Manager{
		Employee:e1,
		level: "SeniorSE-II",
	}

	m2 := Manager{
		Employee: Employee{name:"Alice",employee_id: 43120},
		level: "Principal-E-I",

	}
	m2.Work()
	fmt.Println(m1)
	fmt.Println(m2)
}