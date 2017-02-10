package main

//import "fmt"

var sum = 10

func addName(sum *int)  {
	*sum+=10
}
/*
func main()  {
	addName(&sum)
	fmt.Println(sum)
}
*/