package test

import (
	"fmt"
	"testing"
)

func ArrayMerge(a, b []string) []string {
	combine := append(a, b...)

	// masukan array ke key map karena key map hanya bisa menerima unique key
	unique := map[string]bool{}
	for _, v := range combine {
		unique[v] = true
	}

	// masukan map tadi ke array
	res := []string{}
	for i := range unique {
		res = append(res, i)
	}
	return res
}

func TestCombineArray(t *testing.T) {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin", "akuma"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
