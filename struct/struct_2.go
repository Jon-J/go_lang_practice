package main
import "fmt"

func main() {
	sid := struct{
		firstName ,lastName string;
		age int;
		alive bool;
	}{
		firstName: "sheetal",
		lastName: "kumar",
		age: 12,
		alive: true,
	}
	fmt.Println("Hi -",sid)
}

