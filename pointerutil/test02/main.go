package main

import (
	"fmt"
)

type Stu struct {
	Name string 
}

func main() {
	a := Stu{
		Name : "allen",
	}

	b := &a 

	fmt.Printf("%p, %p", &a, b)
}