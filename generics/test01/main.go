package main


type Container[T any] struct {

}


func Do(c Container[Student]) error {

	return nil 
}


type Student struct {

}

var student Container[Student] 
