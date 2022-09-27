package main

import (
	"fmt"
	"time"
)


func main() {
	ticker := time.NewTicker(1 * time.Second)

	for e := range ticker.C {
		fmt.Println(e)
		time.Sleep(3 * time.Second)
	}

}