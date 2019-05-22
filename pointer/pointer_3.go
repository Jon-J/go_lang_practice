package main
import "fmt"

type Student struct {
	Name string;
}

func Bob(s Student) {
	s.Name = "Bob" // This function will change name of local variable.
}

func Cool(ps *Student) {
	ps.Name = "Cool"
}

func main() {
	s := Student{"Alice"}
	fmt.Println("Name -", s)

	Bob(s)
	fmt.Println("Name -", s) //Name will be "Alice", because data is copied to another memory location of that variable.

	Cool(&s)
	fmt.Println("Name -", s) //Name will be "Cool", hence we are sending pointer so name will be modified from same location.
}
