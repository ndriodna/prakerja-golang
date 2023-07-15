package test

import (
	"fmt"
	"testing"
)

func TestTrapesium(t *testing.T) {
	a := 5
	b := 10
	h := 6
	z := 0.5
	luas := z * float64((a+b)*h) //konversi int ke float karena operator matematika hanya bisa digunakan dg variable tipe yg sama
	fmt.Println(luas)
}
