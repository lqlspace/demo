package main

import "fmt"


func main() {
	i := 0
	for ; i < 10; i++ {
		if i >= 0 {
			if i == 5 {
				break 
			} 
		}
	}

	fmt.Println("i = ", i)
}