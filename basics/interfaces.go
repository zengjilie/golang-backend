package basics

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	length float64
}

func (s square) getArea() float64 {
	return s.length * s.length
}

// Caveat: exported function must start with capital letters
func PrintShape() {
	shapes := []shape{
		square{2.3},
		square{3.3},
		square{4.3},
	}
	for _, v := range shapes {
		fmt.Println(v)
	}
}
