package base_info

import (
	"github.com/headzoo/surf"
	"fmt"
)

type DealerShip struct {
	Name string
	City string
}

type User struct {
	Id string
	Dealer *DealerShip
}

func TestStruct() string {
	var city string
	dealer := DealerShip{Name:"Igor", City:"Bryansk"}
	city = dealer.Name
	return city

}


func GetSite(url string) string  {
	br := surf.NewBrowser()
	err := br.Open(url)
	if err != nil{
		panic(err)
	}

	return br.Body()
}

func AnonimStruct()  {
	var config  struct{
		Name string
	}
	config.Name = "Vasya"
	
	data := struct {
		Id int
		Name string
	}{
		Id:55, Name: "Igor",
	}
	fmt.Println(data.Id, config.Name)
}

