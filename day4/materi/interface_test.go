package test

import (
	"fmt"
	"testing"
)

type Calculate interface {
	large() int
}

type Square struct {
	side int
}

func (s Square) large() int {
	return s.side * s.side
}

func (s Square) noInterfaceLarge() int {
	return s.side * s.side
}

func TestInterface(t *testing.T) {
	// buat kontrak dg interface calculate, maka tidak dapat menggunakan method diluar dari interface
	var s Calculate = Square{4}
	fmt.Println(s.large())

	// jika ingin menggunakan method diluar interface gunakan inisialisai origin struct
	sq := Square{4}
	fmt.Println(sq.noInterfaceLarge())

}

// empty interface, untuk dynamic value

func getVal(v interface{}) {
	fmt.Println(v)
}

// switch case dg interface kosong
func dynamicType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("this is string", v.(string))
	case int:
		fmt.Println("this is int", v.(int))
	default:
		fmt.Println("this is another type", v)
	}
}
func TestEmptyInterface(t *testing.T) {

	// empty interface
	getVal(1)
	getVal("string")
	getVal(true)

	// type assertion dg type interface untuk dynamic type
	var dynamic interface{}
	dynamic = 2
	num := dynamic.(int) * 10
	fmt.Println(num)

	// inisialisasi ulang lalu masukan type yang ingin digunakan
	dynamic = "lele ?"
	str := dynamic.(string) + "lele"
	fmt.Println(str)

	dynamic = []int{1, 2, 3, 4}
	slice := dynamic.([]int)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	dynamicType("string")
	dynamicType(123)
	dynamicType(true)
	dynamicType(1.2)

}
