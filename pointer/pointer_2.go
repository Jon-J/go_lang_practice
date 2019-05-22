package main

import "fmt"

type Person struct {
	name string;
	age int;
}

func main() {
	a := Person{"Ravi", 32};
	ps1 := &a
	ps2 := &Person{"sid", 32};
	fmt.Println(ps1)
	fmt.Println(&ps1)
	fmt.Println(ps2)
	fmt.Println(&ps2)

	fmt.Println(ps1, ps1.name)
	fmt.Println(ps2, ps2.name)
}
