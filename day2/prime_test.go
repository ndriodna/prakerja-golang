package test

import (
	"fmt"
	"testing"
)

func TestPrima(t *testing.T) {
	num := 11 //masukan nilai untuk di cek

	total := 0
	if num < 2 {
		fmt.Println("harus lebih dari 1")
	} else {
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				total++
			}
		}
	}
	if total < 2 {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "adalah bukan bilangan prima")
	}
}

// optimasi dg chatgpt
func TestPrime(t *testing.T) {
	num := 110 //masukan nilai untuk di cek

	isPrime := true
	if num < 2 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "adalah bukan bilangan prima")
	}
}
