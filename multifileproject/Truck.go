package main

type Truck struct {
	numOfDor int
	bedZize int
	weightByTon weight
}

type weight string

const (
	oneTon weight = "One ton"
	twoTon weight = "Two ton "
)
