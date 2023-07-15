package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	// len, digunakan untuk mengetahui panjang string
	hello := "hello world"

	lenHello := len(hello)

	fmt.Println(lenHello)

	// contains, digunakan untuk mencari string
	string := "something"
	substring := "thing"

	contain := strings.Contains(string, substring)
	fmt.Println(contain)

	// substring
	val := "cat;dog"
	substr := val[4:]
	fmt.Println(substr)

	// replace
	s := "this [thing] i would like to remove"
	text := strings.Replace(s, "[", "", -1)
	fmt.Println(text)

	// Insert
	p := "green"
	index := 2
	// p index mengambil dari depan hingga index , q mengambil dari belakang hingga index
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)

}

// function
// variadic
func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

// closures

func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func TestFunction(t *testing.T) {
	// variadic, bisa memasukan lebih dari 1 data dapat berupa int/string dipisahkan dg koma atau slice
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)

	// anonymous function, function tanpa nama
	func() {
		fmt.Println("ini anonymous function")
	}()

	// anonymous function dg variable seperti di js
	anon := func() {
		fmt.Println("ini anonymous function variable")
	}
	anon()

	// anonymous function, passing parameter dibawah khusus function tanpa nama / variable
	func(params string) {
		fmt.Println("ini anonymous function dengan params", params)
	}("params")

	// closures
	// berguna  untuk kasus dimana fungsi ingin state tetap ada tanpa bantuan variable global meski fungsi dipanggil berkali2
	counter := NewCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// defer, akan dieksekusi ketika program selesai
	// disini pakai anonymous func sebagai contoh
	defer func() {
		fmt.Println("ini defer akan jalan terakhir")
	}()
	fmt.Println("ini jalan sebelum defer")
}
