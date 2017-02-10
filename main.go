package main

import (
	"Udemy/base_info"
	"fmt"
)



func main()  {
	name:= base_info.Name
	base_info.Change(&name)
	fmt.Println(name)
	base_info.AnonimStruct()
}

func variatic(){
	base_info.CreateFileAndWrite()
	base_info.ManyParamsLikePython("Car", "Phonto", "Player")
	params:=[]string{"Name", "Family", "CerName"}
	base_info.ManyParamsLikePython(params...)
}

func printFunct()  {
	percent, sold:= base_info.GetDillerInfo()
	fmt.Println(percent, sold)
}

func getSite() string{
	fmt.Println(base_info.TestStruct())
	return base_info.GetSite("http://alib.ru")
}

func curStruct(){
	stru:= person{name:"Igor", address:"Br"}
	fmt.Println(stru)
}

func myStructs()  {
	test:= base_info.DealerShip{Name:"Igor", City:"Bryansk"}
	fmt.Println(test.Name, test.City)
}

func myTypes(){
	fmt.Println(
		base_info.ConvertStringToInt(),
		base_info.ConvertIntToString(),
	)
}

func myConstants()  {
	fmt.Println(
		base_info.MYCONSTANT,
		base_info.MyNewCar,
		base_info.MyNewPlane,
		base_info.MyStringCONSTANT,
	)
}

func variables(){
	fmt.Println(
		base_info.CarType,
		base_info.PlaneType,
		base_info.C_Type,
		base_info.P_Type,
	)
}

type person struct {
	name string
	city string
	country string
	address string
	money int
}
