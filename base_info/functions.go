package base_info

import (
	"fmt"
	"os"

)
type MyDialer struct {
	Name string
	City string
}

func CreateFileAndWrite()  {
	file:= CreateFile("test.txt")
	WriteFile(file, "hello word")
}


func WriteFile(file *os.File, DialerName string)  {
	fmt.Fprintln(file, DialerName)
	defer file.Close()
}

func CreateFile(path string) *os.File{
	file, err := os.Create(path)
	if err != nil{
		panic(err)
	}
	return file
}

func getDialerInfo(s1 string, s2 string) string  {
	res := s1 + " " + s2
	return res
}

func GetDillerInfo() (float32, string){
	dilar:= MyDialer{"BMW", "Bryansk"}
	percent, sold := DialerSoldInfo(100, 175, dilar)
	return  percent, sold

}

func DialerSoldInfo(reaminig float32, start float32, dilaer MyDialer) (percentSold float32, info string){
	info = dilaer.Name + " sold"
	percentSold = 1 - reaminig/start
	return percentSold, info
}

func ManyParamsLikePython(dealer ...string)  {
	for _, item := range dealer{
		fmt.Println(item)
	}

}