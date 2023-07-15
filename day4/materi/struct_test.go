package test

import (
	"fmt"
	"math"
	"testing"
)

type Person struct {
	firtstName, lastName string
	age                  int
}

func TestStruct(t *testing.T) {
	// deklarasi panjang
	var person0 Person
	person0.firtstName = "lele"
	person0.lastName = "llu"
	person0.age = 20

	fmt.Println(person0)

	// deklarasi langsung tanpa menyebutkan nama field
	person1 := Person{"lesi", "lusa", 24}
	fmt.Println(person1)

	// deklarasi langsung dengan nama field
	person2 := Person{
		firtstName: "jaja",
		lastName:   "jamal",
		age:        30,
	}

	fmt.Println(person2)

	// deklarasi pointer dengan new()
	person3 := new(Person)
	person3.firtstName = "pointer"
	person3.lastName = "ini"
	person3.age = 33

	fmt.Println(person3)  // hasilnya akan &{pointer ini 33}
	fmt.Println(*person3) //hasilnya akan pointer ini 33
}

// method
// method berbeda dari function, method terisolasi dg type seperti struct dg code dibawah ini kita membuat method dari struct Person

func (p *Person) GetFullName() string {
	return p.firtstName + " " + p.lastName
}

// method mengatasi nama fungsi yang duplikat
type Rect struct {
	widht, height float64
}

type Circle struct {
	radius float64
}

// disini menggunakan pointer sebagai receiver yang berarti method memiliki akses mengubah object struct, tetapi jika tidak menggunakan pointer nilai pada struct tidak akan berubah
// untuk proses perhitungan dibawah tidak perlu menggunakan pointer jika tidak ada proses assigment untuk merubah nilai struct
func (r *Rect) Area() float64 {
	r.widht = 2
	return r.widht * r.height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func TestMethodStruct(t *testing.T) {
	fullName := Person{
		firtstName: "rizal",
		lastName:   "adit",
		age:        24,
	}
	fmt.Println(fullName.GetFullName())

	r := Rect{3.4, 2.4}
	c := Circle{5.0}

	fmt.Println("sebelum mengubah nilai dg method", r)

	fmt.Println(r.Area())
	fmt.Println(c.Area())

	fmt.Println("sesudah mengubah nilai dg method", r)

}
