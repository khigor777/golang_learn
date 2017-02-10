package base_info

import (
	"strconv"
)

func ConvertStringToInt() int {
	var myStringInt string = "10"
	var myInt int = 15
	number, _ := strconv.Atoi(myStringInt)
	return number + myInt
}

func ConvertIntToString() string {
	var myInt int = 55
	var myString string  = "Igor"
	str := strconv.Itoa(myInt)
	return str + " " + myString
}