package figuras

type Square struct {
	side float64
}

//Constructor
func (square *Square) Construct(pSide float64) {
	square.side = pSide
}

//Getter y Setter
func (square *Square) GetSide() float64 {
	return square.side
}

func (square *Square) SetSide(pSide float64) {
	square.side = pSide
}

func (square *Square) getArea() float64 {
	return square.side * square.side
}

func (square *Square) getPerimeter() float64 {
	return 4 * square.side
}
