package latihan

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) HitungJarak() float64 {
	fuel := 1.5
	distance := c.FuelIn * fuel
	return distance
}
func HitungJarak() {
	car := Car{
		Type:   "sedan",
		FuelIn: 2,
	}

	fmt.Println("Tipe", car.Type)
	fmt.Println("Bensin (liter)", car.FuelIn)
	fmt.Println("Jarak yang dapat ditempuh", car.HitungJarak())
}
