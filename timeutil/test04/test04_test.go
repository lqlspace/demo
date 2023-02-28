package test04

import (
	"fmt"
	"testing"
	"time"
)

func Test04(t *testing.T) {
	tt := time.Now().Add(time.Duration(24)* time.Hour)

	fmt.Println("tt = ", tt)
}