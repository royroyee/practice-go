package main

import "fmt"

type Stringer interface { // java 의 toString() 메서드 느낌이랄까? 물론 이건 인터페이스라고 한다만..
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {

	student := Student{"철수", 12}

	var stringer Stringer = student

	fmt.Printf("%s\n", student.String())

	fmt.Printf("%s\n", stringer.String())

}
