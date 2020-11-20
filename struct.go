package main

import "fmt"

type dog struct {
	name string
	age  int
	sex  string
};

type person struct {
	name string
	age  int
};

func main() {
	p := person{name: "John", age: 23}
	fmt.Println(p)
	d := dog{name: "Fido", age: 10, sex: "Male"}
	// uses '.' operator to just select the name field: will print "Fido"
	fmt.Println(d.name)
}
