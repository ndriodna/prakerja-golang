package test

import (
	"fmt"
	"testing"
)

func Mapping(slice []string) map[string]int {
	res := map[string]int{}
	for _, v := range slice {
		res[v]++
	}
	return res
}

func TestCountSlice(t *testing.T) {
	fmt.Println(Mapping([]string{"asd", "qwe", "lele", "lele", "qwe", "asd", "asd", "asd"}))
	fmt.Println(Mapping([]string{}))
}
