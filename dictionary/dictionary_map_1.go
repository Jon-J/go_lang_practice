package main

import "fmt"

func main() {
    test := make(map[string]string)
    test["sheetal"] = "vada pav"
    test["sid"] = "bhel"
    test["ravi kumar"] = "paratha"
    test["ravi kiran"] = "beene dosa"

    for name, dish := range(test){
        fmt.Println("name -", name, ", fav dish -",dish)
    }

   n := map[string]int{"foo": 1, "bar": 2}
   fmt.Println(n)
}
