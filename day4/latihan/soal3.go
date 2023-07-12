package latihan

func GetMinMax(num ...*int) (min, max int) {
	min = 0
	max = 0
	for _, v := range num {
		if *v > max {
			max = *v
		} else {
			min = *v
		}
	}
	return min, max
}
