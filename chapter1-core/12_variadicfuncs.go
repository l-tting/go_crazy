package main



func demoVariadics(nums ...int) int{
	sum := 0
	for _ ,num := range nums{
		sum+= num
	}
	return sum
	
}

func demoVariad(chars ...string) string{
	message := "Hello " 
	for _,extra_m := range chars{
		message += extra_m  + " "
	}
	return message
}

