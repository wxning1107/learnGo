package main

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	//局部变量s逃逸到堆
	s := new(Student)
	s.Name = name
	s.Age = age

	return s
}

//  go build -gcflags=-m
func main() {
	StudentRegister("wxning", 18)
}
