package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan int, 1)

	go func() {
		count := 0
		for {
			select {
			case ch <- count:
			default:
			}

			if count >= 30 {
				close(ch)
				break 
			}
			count++

			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(3 * time.Second)
		cnt, ok  := <- ch
		if ok {
			fmt.Println("cnt = ", cnt)
		} else {
			break 
		}
	}
}