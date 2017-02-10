package main

import (
	"fmt"
)

var ShippingDays = 20
var ShippingDaysPtr = new(int)

func main()  {
	*ShippingDaysPtr = 15
	ShippinDaysAdjastmentsPtr(ShippingDaysPtr)
	fmt.Println(ShippingDays)
	fmt.Println(*ShippingDaysPtr)

	shipper := Shipper{}
	shipper.id = 50
	shipper.name = "TryIt"
	ShipperUpdate(&shipper)
	fmt.Println(shipper.name, shipper.id)

}

func ShippinDaysAdjastments(shippingDays int)  {
	shippingDays += 10
}

func ShippinDaysAdjastmentsPtr(shippingDays *int)  {
	*shippingDays += 10
}

func ShipperUpdate(shipper *Shipper)  {
	shipper.id+=10
	shipper.name +=", hello"
}

type Shipper struct {
	name string
	id int
}


func getSum(x int, y int) int  {
	return (x + y)*2/35*5
}