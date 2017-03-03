package main

import (
	"fmt"
	"runtime"
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

	runtime.GOMAXPROCS(5)
	go abcGen()
	//time.Sleep(10 * time.Millisecond)
	obj := new(mynewStruct)
	obj.name = "My Name"
	obj.number = 105
	obj.addressdetails = "line1"

	somenew := obj
	obj.name = "New Name"
	fmt.Println(*obj)

	fmt.Println("SOMENEW **")
	fmt.Println(somenew)
	x := make(map[string]string)
	x["first"] = "First Item"
	fmt.Println(x)
	println(x["first"])

	p := New(100, "Another Name", "Some Address")
	fmt.Println(p)

	m := make(map[int]string)

	m[1] = "Uno"
	m[2] = "Duos"
	m[3] = "Threz"

	for idx, val := range m {
		println(idx, val)
	}

	n := []string{"this", "is", "great"}

	for _, value := range n { // foreach (var x in n)
		println(value)
	}

	t := make(map[int]mynewStruct)

	t[1] = *New(44, "Forty Four", "My Address")
	t[2] = *New(55, "Fifty Five", "My new Address")
	t[3] = *New(66, "sdfsdf", "sfdsf")
	cc := New(77, "sdfsdf", "qweqwe")

	fmt.Println(t)
	fmt.Println(*cc)

	AnotherFunction(100)

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

//AnotherFunction just prints an integer
func AnotherFunction(number int) *mynewStruct {
	fmt.Print(number)
	return nil
}
