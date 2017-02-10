package main

import (
	"fmt"
)

func main()  {
	carA :=car{}
	fmt.Println(carA)
	carA.mpg = 55
	carA.numberOfDors = 2
	fmt.Println(carA.numberOfDors)
	carA.getMpg()
	bCar:= car{vichcle{25, 3},Red}
	fmt.Println(bCar)
	bCar.getMpg()
	cCar := car{}
	cCar.color = Red
	cCar.mpg = 29
	cCar.getMpg()
	defaultCar :=newCar()
	fmt.Println(defaultCar)
}

type vichcle struct {
	mpg int
	numberOfDors int
}

type car struct {
	vichcle
	color Color
}

//Create Constructor

func newCar() *car  {
	result:= &car{
		color:Blue,
	}
	result.numberOfDors = 25
	result.mpg = 25
	return result
}

func (v *vichcle)getMpg()  {

	fmt.Println("vichel get mpg", v.mpg)
}

type Color string

//custom types
const (
	Blue Color = "blue"
	Red Color = "red"
	Black Color = "Black"
)