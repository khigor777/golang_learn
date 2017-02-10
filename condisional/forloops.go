package main

import (
	"fmt"
)

func main()  {
	testMapLoop()
	testForLoop()
}

func testMapLoop() {
	MyMap := make(map[string]int)
	MyMap["igor"] = 31
	MyMap["ivan"] = 32

	//it dosen't work
	for m := 0; m < len(MyMap); m++ {

	}

	for key, val := range MyMap {
		fmt.Println(key, val)
	}

	for _, val := range MyMap {
		fmt.Println(val)
	}

	for key := range MyMap {
		fmt.Println(key)
	}
}


func testForLoop()  {
	cars := [3]string{"toyota", "ford", "tesla"}

	i :=0
	for i < len(cars){
		fmt.Println(cars[i])
		i++
	}

	for j:=0; j < len(cars); j++{
		fmt.Println(cars[j])
	}

	for k, v := range cars{
		fmt.Println(k, v)
	}

}