package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	line1 address = "AddressLine1"
	line2 address = "AddressLine2"
	line3 address = "Zipcode"
)

type mynewStruct struct {
	number         int
	name           string
	addressdetails string
}

func main() {

	runtime.GOMAXPROCS(1)
	//go abcGen()
	time.Sleep(10 * time.Millisecond)
	obj := new(mynewStruct)
	obj.name = "My Name"
	obj.number = 105
	obj.addressdetails = "line1"

	fmt.Println(obj)

	//	a := make([]int, 10, 11)
	// fmt.Println(a)
	// println(len(a))
	x := make(map[string]string)
	x["first"] = "First Item"
	fmt.Println(x)
	println(x["first"])

	p := New(100, "Another Name", "Some Address")
	fmt.Println(p)
}

func abcGen() {
	for l := byte('A'); l <= byte('Z'); l++ {
		go println(string(l))
	}
}

type address string

// func (mf *mynewStruct) anewfunction() {

// }

// New function is a constructor
func New(numberparam int, nameparam, addressparam string) *mynewStruct {
	// obj := new(mynewStruct)
	// obj.name = nameparam
	// obj.number = numberparam
	// obj.addressdetails = addressdetailsparam
	// return obj

	return &mynewStruct{number: numberparam, name: nameparam, addressdetails: addressparam}
}
