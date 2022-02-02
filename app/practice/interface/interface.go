package interfaces

import (
	"fmt"
	"strings"
)

func printIf(src interface{}) {
	value, ok := src.(int)
	if ok {
		fmt.Printf("parameter is integer. [value: %d]\n", value)
		return
	}

	if value, ok := src.(string); ok {
		value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
		fmt.Printf("parameter is string. [value: %s]\n", value)
		return
	}

	if value, ok := src.([]string); ok {
		value = append(value, "unknown") // 対象がsliceなのでAppendができる
		fmt.Printf("parameter is slice string. [value: %s]\n", value)
		return
	}

	if values, ok := src.([3]int); ok {
		value = values[0]
		fmt.Printf("parameter is array int. [value: %d]\n", value)
		return
	}

	fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
}

func Base() {
	// var x interface{}
	// x = 1
	// x = "aaa"
	// x = [...]int{0, 1, 2}
	// value, ok := x.(int)
	// hoge := value + 100
	// fmt.Println(ok, hoge+2)
	// s := "100"
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(ok, i)
	// }
	/* ------------------------------------ */
	printIf(12)
	printIf("hello")
	printIf([]string{"cat", "dog"})
	printIf([2]string{"hello", "world"})
	printIf([...]int{1, 2, 3})
}

/* ----------------- */
type Person struct {
	Common string
	Name   string
	Age    int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%vm, Age=%v, CustomCommon=%v", p.Name, p.Age, p.Common+"ですよーーー")
}
func (p *Person) Commons() string {
	return fmt.Sprintf("p.Common=%vm", p.Common)
}

type Car struct {
	Common string
	Number string
	Model  string
}

func (p *Car) ToString() string {
	return fmt.Sprintf("Number=%vm, Model=%v, CustomCommon=%v", p.Number, p.Model, p.Common+"ですよーーー")
}
func (p *Car) Commons() string {
	return fmt.Sprintf("p.Common=%vm", p.Common)
}

func commonFunc() {

	type Stringfy interface {
		Commons() string
		ToString() string
	}

	vs := []Stringfy{
		&Person{Name: "taro", Age: 33, Common: "PersonのCommon"},
		&Car{Number: "taro", Model: "model", Common: "CarのCommon"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
		fmt.Println(v.Commons())
	}
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生した", ErrCode: 1234}
}

func customError() {
	err := RaiseError()
	fmt.Println(err.Error())
	fmt.Println("errは", err)
	e, ok := err.(*MyError)
	if ok {
		fmt.Println("ErrCode", e.ErrCode)
	}
}

type Point struct {
	A int
	B string
}

func iStringer() {
	p := &Point{A: 1000, B: "BBBBBB"}
	fmt.Println(p)
}

func (p *Point) String() string {
	return fmt.Sprintf("%v, %v", p.A, p.B)
}

func InterfacePractice() {
	Base()
	// interface 異なる型に共通の性質を付与する
	// commonFunc()
	// カスタムエラー
	// customError()
	// interface Stringer
	// iStringer()
}
