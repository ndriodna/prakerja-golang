package main

import (
	"day_four/latihan"
	"fmt"
)

func main() {
	// soal 1
	// latihan.HitungJarak()

	// soal 2
	soal2 := latihan.Student{}
	for i := 0; i < 5; i++ {
		// masukan nilai kosong untuk mendapatkan index name
		soal2.Name = append(soal2.Name, "")
		fmt.Println("Masukan nama :")
		fmt.Scanln(&soal2.Name[i])

		// masukan nilai kosong untuk mendapatkan index score
		soal2.Score = append(soal2.Score, 0)
		fmt.Println("Masukan score :")
		fmt.Scanln(&soal2.Score[i])
	}
	fmt.Println(soal2.Average())
	fmt.Println(soal2.Min())
	fmt.Println(soal2.Max())

	// soal 3
	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)
	// min, max = latihan.GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println(min, "is min number")
	// fmt.Println(max, "is max number")
}
