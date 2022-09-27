package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
}

type People struct {
	stu []byte 
}
func main() {
	var p People
	var s *Student
	err := json.Unmarshal(p.stu, &s)
	if err != nil {
		fmt.Println("err = ", err)
	}

	ss := Student{Name: "allen"}
	bs, _ := json.Marshal(&ss)
	pp := People{
		stu: bs,
	}

	err = json.Unmarshal(pp.stu, &s)
	if err != nil {
		fmt.Println("err = ", err)
	}
}