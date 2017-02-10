package main

import (
	"fmt"
)

func main(){
	arrays()
}

func arrays()  {
	var arrStr [3]string
	arrStr[0] = "Toyota"
	arrStr[1] = "Ford"
	arrStr[2] = "Nissan"

	fmt.Println(arrStr[1])

	arrStr2 := [3]string{"Ford", "Toyota", "Nissan"}
	fmt.Println(arrStr2[2])

	arrSlice := []string{"Ford", "Toyota", "Nissan"}
	fmt.Println(arrSlice[len(arrSlice)-1])

	arrSlice = append(arrSlice, "Tesla")
	fmt.Println(arrSlice)

	copy(arrSlice, []string{"Mitsubishi"})
	fmt.Println(arrSlice)

	arrStrMake := make([]string, 3)
	arrStrMake[0] = "Toyota"
	arrStrMake[1] = "Ford"
	arrStrMake[2] = "Nissan"

	newSlice := arrStrMake[1:2]
	fmt.Println(newSlice[0])

	myNewArr := []string{}
	myNewArr = append(myNewArr, "igor")
	fmt.Println(myNewArr)

}

func cond()  {
	carLotA := 30
	carLotB := 50

	if carLotA > carLotB{
		fmt.Println("ok")
	}else if carLotA > 29{
		fmt.Printf("grater %v", carLotA)
	}else {
		fmt.Println("No")
	}

	switch {
	case carLotA > 30:
		fmt.Println("carLotA > 20")
	case carLotA < 50:
		fmt.Println("carLotA < 50")
	default:
		fmt.Println("default")
	}

	switch carLotB{
	case 30:
		fmt.Println("30")
	case 51, 52:
		fmt.Println("51,52")
	default:
		fmt.Println("default")
	}
}