package test

import (
	"fmt"
	"testing"
)

func TestKelipatanTujuh(t *testing.T) {
	limit := 100 //masukan batas kelipatan
	total := 0
	if limit > 7 {
		for i := 7; i <= limit; i++ {
			// jika data habis dibagi 7 = 0 termasuk kelipatan ke 7
			if i%7 == 0 {
				total++
				fmt.Println(i)
			}
		}
	}
	fmt.Println(limit, "terdapat", total, "kelipatan 7")
}
