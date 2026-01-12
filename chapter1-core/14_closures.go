package main



func demoClosure() func() int{
	count := 0
	return func() int {
		count++
		return count
	}
}

func Outer() func() int{ // higher order 
	x := 8

	return func() int{
		return x
	}


}