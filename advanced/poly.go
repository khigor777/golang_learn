package main

import "fmt"

type Colors string

const (
	RED   Colors = "RED"
	GREEN Colors = "GREEN"
	BLUE  Colors = "BLUE"
)
type vichele struct {
	Mpg int
	NumberOfDors int
}

type myCar struct {
	vichele
	color Colors
	curName string
}

func (c *myCar) GetVin() string{
	return c.curName
}

type truck struct {
	vichele
	color Colors
	curName string
}

func (tr *truck) GetVin() string {
	return tr.curName
}


type manifactures interface {
	GetVin() string
}


func main()  {
// polimorphic
	aTruc := truck{vichele:vichele{15, 2},curName:"Truck"}
	aCar := myCar{vichele:vichele{20, 30}, curName:"myCar"}
	mans:= [...]manifactures{&aTruc, &aCar}
	for i, _ := range mans{
		fmt.Println(mans[i].GetVin())
	}

}