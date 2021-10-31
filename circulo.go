package figuras

import "math"

type Circle struct {
	radio float64
}

//constructor
func (circle *Circle) Construct(pRadio float64) {
	circle.radio = pRadio
}

//getter y setter
func (circle *Circle) GetRadio() float64 {
	return circle.radio
}

func (circle *Circle) SetRadio(pRadio float64) {
	circle.radio = pRadio
}

func (circle *Circle) getArea() float64 {
	return math.Pi * (circle.radio * circle.radio)
}

func (circle *Circle) getPerimeter() float64 {
	return 2 * (math.Pi * circle.radio)
}
