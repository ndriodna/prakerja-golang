package test

import (
	"fmt"
	"testing"
)

// pointer memiliki 2 operator * dan &
// * digunakan untuk mendeklarasikan pointer dan mengakses nilai pointer * bisa merubah value dari address
// & mengakses address pointer dan mengembalikan address
func TestPointer(t *testing.T) {
	name := "jhon"
	nameAddress := &name
	fmt.Println("value", name) //value jhon
	fmt.Println(&name)         //address jhon
	fmt.Println(*nameAddress)  //value pointer jhon
	fmt.Println(&nameAddress)  //value address jhon

	// ubah reference pointer akan merubah semua
	name = "lele"

	// merubah pointer name lewat address menggunakan *
	*nameAddress = "lulu"
	fmt.Println("value", name) //value lulu
	fmt.Println(&name)         //address lulu
	fmt.Println(*nameAddress)  //value pointer lulu
	fmt.Println(&nameAddress)  //value address lulu

	// zero value pointer
	num_a := 25
	var num_a_pointer *int
	if num_a_pointer == nil {
		fmt.Println(num_a_pointer)
		num_a_pointer = &num_a
		fmt.Println(num_a_pointer)  // akan mengembailkan address
		fmt.Println(*num_a_pointer) // akan mengembailkan value

	}

	// kita bisa mendeklarasikan pointer dengan fungsi bawaan new()
	var size = new(int)
	// cara ini hanya bisa dilakukan dg printf seperti %d, %T dll
	fmt.Printf("size value is %d \n", *size)  //mengecek nilai dari pointer
	fmt.Printf("size type is %T \n", size)    //mengecek type dari pointer
	fmt.Printf("size address is %v \n", size) //mengecek address dari pointer
	*size = 90
	fmt.Println("new size", *size)
}
