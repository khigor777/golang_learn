package main

import "testing"

func TestGetSum(t *testing.T)  {
	result := getSum(10, 20)
	var actualInt int = 5
	if actualInt != result{
		t.Fatalf("Error expected %v actual %v", result, actualInt)
	}
}

func BenchmarkGetSum(b *testing.B)  {
	for i:=0; i<=1000000; i++{
		go getSum(i + 10, i+50)
	}
}