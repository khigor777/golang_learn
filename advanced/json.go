package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main()  {
	url:= "http://localhost:8181"
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		panic(err)
	}
	var p Payload
	err = json.Unmarshal(body, &p)

	if err != nil{
		panic(err)
	}

	fmt.Println(p.Envelop)
}


type Payload struct {
	Envelop Data
}

type Data struct {
	CarsByDor Cars
	TrucksByTon Trucks
}

type Cars map[string]int
type Trucks map[string]int
