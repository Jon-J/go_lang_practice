package main

import (
	"fmt"
	"time"
	)

func test(input string) {
	for i := 0; i < 5; i++{
		fmt.Println("Printing value -",input)
		time.Sleep(time.Second * 1)
	}
}
func main() {

	go test("sheep")
	test("fish")
}

