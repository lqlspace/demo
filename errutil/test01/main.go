package main

import (
	"errors"
	"fmt"
)

var (
	ErrDataNotFound = errors.New("data not found")
)

func main() {
	err := fmt.Errorf("%w: yes", ErrDataNotFound)

	fmt.Println(errors.Is(err, ErrDataNotFound))
}