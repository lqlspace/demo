package main

import "fmt"

type Student struct {
	Name string `json:"name"`
}

var stu Student

const c = 1 

type Inter interface{
	
}

func greeter() {
	fmt.Println("hello, world")
}