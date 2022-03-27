package pointer

import "fmt"

type S struct{ value string }

func (s *S) SetB(v string) {
	s.value = v
}

func base() {
	var s S
	// s.SetA("a")
	fmt.Println(s.value) // sはゼロ値のまま
	s.SetB("b")
	fmt.Println(s.value) // b
}

func Pointer() {
	base()
}
