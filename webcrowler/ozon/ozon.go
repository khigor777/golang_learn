package ozon

import (
	_ "fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"

)

type Book struct {
	Title string
	Price string

}

type Tags string



func ParseBook(url string)  *Book {
	testUrl := "https://www.ozon.ru/context/detail/id/137634276/"
	resp, err := http.Get(testUrl)
	checkError(err)
	defer resp.Body.Close()
	ctype:= resp.Header.Get("Content-Type")
	body, err:= charset.NewReader(resp.Body, ctype)
	checkError(err)
	htmlDoc, err := goquery.NewDocumentFromReader(body)
	title := htmlDoc.Find("h1").Text()
	price:= htmlDoc.Find("div.bOzonPrice").Text()
	return &Book{Title:title, Price:price}

}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}

