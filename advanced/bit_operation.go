package main

import "fmt"

var i int = 10
var b int = 5

func main()  {
	i = i << 1
	b = b >> 1
	z := 10
	y :=10
	x := z | y | 10
	fmt.Println(i, b, x)
	println(factorial(10))
}
