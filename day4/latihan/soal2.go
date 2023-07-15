package latihan

import "fmt"

type Student struct {
	Name  []string
	Score []int
}

func (s Student) Average() int {
	total := 0
	for _, v := range s.Score {
		total += v
	}
	avg := total / len(s.Score)
	return avg
}

func (s Student) Min() string {
	min := s.Score[0]
	minIndex := 0
	for i, v := range s.Score {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return fmt.Sprint("Nilai minimum ", s.Name[minIndex], ":", min)
}
func (s Student) Max() string {
	max := s.Score[0]
	maxIndex := 0
	for i, v := range s.Score {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	return fmt.Sprint("Nilai maximum ", s.Name[maxIndex], ":", max)
}
