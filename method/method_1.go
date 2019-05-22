package main

import "fmt"

type Employee struct {
	Name  string
	Age   int
	Phone int
}

func (e *Employee) employeeDetails() {
	fmt.Printf("Name %s, age - %d contact details - %d.\n", e.Name, e.Age, e.Phone)
}

func main() {
	emp1 := Employee{"Sheetal", 23, 1234}
	emp2 := Employee{"sid", 32, 4561}
	fmt.Println("First employee details")
	emp1.employeeDetails()
	fmt.Println("Second employee details")
	emp2.employeeDetails()
}
