package methods

import (
	"errors"
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.Name, u.Age)
}

func Hello(name string) (string, error) {
	if name == "invalid" {
		return "", errors.New("invalid name")
	}
	return fmt.Sprintf("Hello %s", name), nil
}

func ReflectMethods() {
	// 型情報の取得
	rt := reflect.TypeOf(1.)

	// 型名
	fmt.Println(rt.Name()) // == float64

	// 型の種別
	fmt.Println(rt.Kind()) // == float64

	// 型の割当サイズ
	fmt.Println(rt.Size()) // == 8

	fmt.Println("-- スライス型独自の型情報")
	rt = reflect.TypeOf([]int{})

	// 要素の型情報
	fmt.Printf("%#v\n", rt.Elem()) // == &reflect.rtype{size:0x8, ptrdata:0x0, hash:0xf75371fa, tflag:0xf, align:0x8, fieldAlign:0x8, kind:0x2, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x10032e0), gcdata:(*uint8)(0x10e85c0), str:951, ptrToThis:27104}

	fmt.Println("-- 配列型独自の型情報")
	rt = reflect.TypeOf([1]int{0})

	// 要素の型情報
	fmt.Printf("%#v\n", rt.Elem()) // == &reflect.rtype{size:0x8, ptrdata:0x0, hash:0xf75371fa, tflag:0xf, align:0x8, fieldAlign:0x8, kind:0x2, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x10032e0), gcdata:(*uint8)(0x10e85c0), str:951, ptrToThis:27104}

	// 配列の長さ
	fmt.Println(rt.Len()) // == 1

	fmt.Println("-- マップ型独自の型情報")
	rt = reflect.TypeOf(map[int]string{})

	// 値の要素の型情報
	fmt.Printf("%#v\n", rt.Elem()) // == &reflect.rtype{size:0x10, ptrdata:0x8, hash:0xe0ff5cb4, tflag:0x7, align:0x8, fieldAlign:0x8, kind:0x18, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x1003440), gcdata:(*uint8)(0x10d1618), str:4914, ptrToThis:35040}

	// キーの要素の型情報
	fmt.Printf("%#v\n", rt.Key()) // == &reflect.rtype{size:0x8, ptrdata:0x0, hash:0xf75371fa, tflag:0xf, align:0x8, fieldAlign:0x8, kind:0x2, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x10032e0), gcdata:(*uint8)(0x10e8d20), str:951, ptrToThis:27168}

	fmt.Println("-- 構造体型独自の型情報")
	rt = reflect.TypeOf(User{})

	// フィールド数
	fmt.Println(rt.NumField()) // == 2

	// フィールドの型情報
	fmt.Printf("%#v\n", rt.Field(0)) // == reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x10d6b60), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}

	// メソッド数
	fmt.Println(rt.NumMethod()) // == 1

	// メソッドの型情報
	fmt.Printf("%#v\n", rt.Method(0)) // == reflect.Method{Name:"String", PkgPath:"", Type:(*reflect.rtype)(0xc00004e180), Func:reflect.Value{typ:(*reflect.rtype)(0xc00004e180), ptr:(unsafe.Pointer)(0xc00000e048), flag:0x13}, Index:0}

	fmt.Println("-- 関数型独自の型情報")
	rt = reflect.TypeOf(Hello)

	// 引数の数
	fmt.Println(rt.NumIn()) // == 1

	// 引数の型情報
	fmt.Printf("%#v\n", rt.In(0)) // == &reflect.rtype{size:0x10, ptrdata:0x8, hash:0xe0ff5cb4, tflag:0x7, align:0x8, fieldAlign:0x8, kind:0x18, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x1003440), gcdata:(*uint8)(0x10fc9e0), str:5178, ptrToThis:39616}

	// 返り値の数
	fmt.Println(rt.NumOut()) // == 2

	// 返り値の型情報
	fmt.Printf("%#v\n", rt.Out(0)) // == &reflect.rtype{size:0x10, ptrdata:0x8, hash:0xe0ff5cb4, tflag:0x7, align:0x8, fieldAlign:0x8, kind:0x18, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x1003440), gcdata:(*uint8)(0x10fc9e0), str:5178, ptrToThis:39616}
}
