package main

import "fmt"

func main()  {
	p :=  new(person) //same thing &person{}
	fmt.Println(p.talk())
	fmt.Println(p.walk())

	o := new(PolliceOficer)
	fmt.Println(o.walk(), o.talk())
}

type person struct {
	firstName string
	lastName string
}

type PolliceOficer struct {
	begadeNumber int
	name string

}

type behavior interface {
	talk() string
	walk() int
	run() bool
}

func (p *person)talk()string  {
	return "Hello"
}

func (p *person) walk() int {
	return 10
}

func (o * PolliceOficer) talk()string  {
	return "PolliceOficer talks"
}

func (o *PolliceOficer) walk() int  {
	return 20
}

func (o *PolliceOficer) run() bool {
	return true
}
