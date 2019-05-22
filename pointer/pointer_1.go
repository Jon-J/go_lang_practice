package main

import "fmt"

type Person struct {
	name string;
	age int;
}

func main() {
	var ps1 *Person = new(Person);
	ps2 := new(Person);
	fmt.Println(ps1)
	fmt.Println(&ps1)
	fmt.Println(ps2)
	fmt.Println(&ps2)

	ps1.name = "Shilpa"
	ps1.age = 28

	ps2.name = "Sheetal"
	ps2.age = 32

	fmt.Println(ps1, ps1.name)
	fmt.Println(ps2, ps2.name)
}
