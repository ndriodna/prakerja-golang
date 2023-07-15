package test

import (
	"errors"
	"fmt"
	"testing"
)

// dg package errors
func ErrorHandle(i int) (int, error) {
	if i <= 0 {
		return -1, errors.New("harus lebih besar dari 0")
	}
	return i, nil
}

// jadi disini panic akan trigger jika tanpa recover program akan berhenti, recover akan menangkap value panic jika terjadi error tanpa perlu membuat program berhenti/crash
func ErrorPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error message", err)
		}
	}()

	if true {
		panic("panic")
	}
}
func TestError(t *testing.T) {
	res, err := ErrorHandle(-1)
	fmt.Println(res, err)

	ErrorPanic()
}
