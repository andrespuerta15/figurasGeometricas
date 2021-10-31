package figuras

import "fmt"

type iFigure interface {
	getArea() float64
	getPerimeter() float64
}

func Calculate(pFigure iFigure) {
	fmt.Println("Size: ", pFigure)
	fmt.Println("Area:", pFigure.getArea())
	fmt.Println("Perimeter:", pFigure.getPerimeter())
}
