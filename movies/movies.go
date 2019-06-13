package movies

import "fmt"

type PriceCode int

const (
	CHILDRENS PriceCode = 2
	REGULAR PriceCode = 0
	NEW_RELEASE PriceCode = 1
)

type Movie struct {
	Title string
	PriceCode PriceCode
}

func NewMovie(title string, priceCode PriceCode) *Movie {
	return &Movie{Title: title, PriceCode: priceCode}
}

func (m *Movie) ToString() string {
	return fmt.Sprintf(
		"title=%s (code=%v)", m.Title, m.PriceCode)
}

type Rental struct {
	Movie *Movie
	DaysRented int
}

func NewRental(movie *Movie, daysRented int) *Rental {
	return &Rental{Movie: movie, DaysRented: daysRented}
}
