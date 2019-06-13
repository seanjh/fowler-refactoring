package movies

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

type Rental struct {
	Movie *Movie
	DaysRented int
}
