package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	// return fmt.Printf("Student Age : %d", s.Age)  Prinft()는 리턴 값이 int 타입
	return fmt.Sprintf("Student Age:%d", s.Age) // Sprintf() 는 문자열을 리턴한다.
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)

}

func main() {

	s := &Student{15}

	PrintAge(s)
}
