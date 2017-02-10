package main

import "fmt"

func main(){
	MapRun()
	//arrays()
}




func MapRun()  {

	carDialer := make(map[string]int)
	carDialer["A1 Auto"] = 76
	fmt.Println(carDialer["A1 Auto"])
	fmt.Println(len(carDialer))

	carDialer["new Car"] = 20
	fmt.Println(carDialer["new Car"])

	delete(carDialer, "new Car")
	fmt.Println(carDialer)

	carDaelerStr := make(map[string]string)
	carDaelerStr["A1 Bryansk"] = "Bryansk"
	carDaelerStr["A1 Bryansk2"] = "Bryansk2"

	fmt.Println(carDaelerStr["A1 Bryansk2"])

}