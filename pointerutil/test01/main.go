package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string `json:"name"`
}

func (self Student) Print() {
	fmt.Print("stu")
	sort.IntSlice()
}

func (self *Student) PrintPointer() {
	fmt.Print("stu pointer")
}


func main() {
	Student.Print(Student{})
	var s *Student
	s.PrintPointer()
}