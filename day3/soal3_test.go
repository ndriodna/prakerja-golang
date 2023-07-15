package test

import (
	"fmt"
	"strconv"
	"testing"
)

func munculSekali(v string) []int {
	var combine []int

	for _, z := range v {
		strToInt, _ := strconv.Atoi(string(z))
		combine = append(combine, strToInt)
	}
	// masukan array ke key map karena key map hanya bisa menerima unique key
	unique := map[int]bool{}
	for _, v := range combine {
		unique[v] = true
	}

	res := []int{}

	for v := range unique {
		res = append(res, v)
	}
	return res
}

func TestCountUnique(t *testing.T) {
	fmt.Println(munculSekali("123454455567676"))
}
