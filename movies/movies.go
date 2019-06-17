package movies

type PriceCode int

const (
	CHILDRENS   PriceCode = 2
	REGULAR     PriceCode = 0
	NEW_RELEASE PriceCode = 1
)

type Price interface {
	Code() PriceCode
	Charge(daysRented int) float64
	FrequentRenterPoints(daysRented int) int
}

type childrensPrice struct{}

func (p *childrensPrice) Code() PriceCode { return CHILDRENS }

func (p *childrensPrice) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += (float64)(daysRented-3) * 1.5
	}
	return result
}

func (p *childrensPrice) FrequentRenterPoints(daysRented int) int { return 1 }

type newReleasePrice struct{}

func (p *newReleasePrice) Code() PriceCode { return NEW_RELEASE }

func (p *newReleasePrice) Charge(daysRented int) float64 {
	return float64(daysRented) * 3.0
}

func (p *newReleasePrice) FrequentRenterPoints(daysRented int) int {
	// add bonus for a two day new release rental
	if daysRented > 1 {
		return 2
	}
	return 1
}

type regularPrice struct{}

func (p *regularPrice) Code() PriceCode { return REGULAR }

func (p *regularPrice) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += (float64)(daysRented-2) * 1.5
	}
	return result
}

func (p *regularPrice) FrequentRenterPoints(daysRented int) int { return 1 }

type Movie struct {
	Title string
	price Price
}

func (m *Movie) PriceCode() PriceCode {
	return m.price.Code()
}

func NewMovie(title string, priceCode PriceCode) *Movie {
	m := &Movie{Title: title}
	m.setPriceCode(priceCode)
	return m
}

func (m *Movie) setPriceCode(priceCode PriceCode) {
	switch priceCode {
	case REGULAR:
		m.price = new(regularPrice)
	case CHILDRENS:
		m.price = new(childrensPrice)
	case NEW_RELEASE:
		m.price = new(newReleasePrice)
	}
}

func (m *Movie) charge(daysRented int) float64 {
	return m.price.Charge(daysRented)
}

func (m *Movie) frequentRenterPoints(daysRented int) int {
	return m.price.FrequentRenterPoints(daysRented)
}

type Rental struct {
	Movie      *Movie
	DaysRented int
}

func (r *Rental) Charge() float64 {
	return r.Movie.charge(r.DaysRented)
}

func (r *Rental) FrequentRenterPoints() int {
	return r.Movie.frequentRenterPoints(r.DaysRented)
}
