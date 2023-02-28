package student

type Student interface {
	Name() string 

	age() int 
}


type stu struct {
	Name string 
	Age int 
}

func (s *stu) Name() string {
	return 
}