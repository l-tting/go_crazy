package main

// import "fmt"

func Recursor(n int) int{
	if n==0 || n == 1{
		return 1
	}
	return n * Recursor(n -1)
}