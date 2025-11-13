package main

func demoClosure() func() int{
	count := 0
	return func() int {
		count++
		return count
	}
}