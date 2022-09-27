package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
}


func main() {
	stu := &Student{
		Name: "allen",
	}

	bs, _ := json.Marshal(&stu)


	var req interface{} = bs 
	req = bs 

	m := make(map[string]interface{})
	json.Unmarshal(req.([]byte), &m)

	fmt.Println(m)
}