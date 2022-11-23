package main

import "testing"

func Testcalculator(t *testing.T){
	if 	Calculator(2) != 4{
		t.Error("Expect 4 check the code")
	}
}