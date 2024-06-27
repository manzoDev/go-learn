package main

import "fmt"

type Human struct {
	Name string;
	Age int;
	Genre string;
	Work bool;
}

func main() {
	humanA := Human {
		Name: "Roger",
		Age: 27,
		Genre: "Male",
		Work: true,
	}

	humanA.sayHi(3)
}

func (r Human) sayHi(a int) {
	fmt.Println("Hi", r.Name, a)
}