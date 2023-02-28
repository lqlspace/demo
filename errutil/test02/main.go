package main

import (
	"strings"
	"fmt"
)


func main() {
	var a []string 

	b := strings.Join(a, ",")

	fmt.Println(b)

	a = []string{}

	b = strings.Join(a, ",")

	fmt.Println(b)

	a = []string{"abc"}

	b = strings.Join(a, ",")

	fmt.Println(b)

	a = []string{"a", "b"}

	b = strings.Join(a, ",")

	fmt.Println(b)
}
