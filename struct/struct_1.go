package main

import "fmt"

type Person struct {
	firstName, lastName string;
	contactDetails int;
	age int;
	alive bool;
}

func main() {
	sid := Person{
		firstName: "sheetal",
		lastName: "kumar",
		contactDetails: 123,
		age: 12,
		alive: true,
		}
	fmt.Println("sid -", sid)

	var test Person;
	fmt.Println("test -",test)

}
