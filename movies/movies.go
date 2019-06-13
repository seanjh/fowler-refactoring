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

func (r *Rental) GetCharge() float64 {
	result := 0.0
	switch (r.Movie.PriceCode) {
	case REGULAR:
		result += 2
		if r.DaysRented > 2 {
			result += (float64)(r.DaysRented - 2) * 1.5
		}
	case NEW_RELEASE:
		result += (float64)(r.DaysRented) * 3
	case CHILDRENS:
		result += 1.5
		if r.DaysRented > 3 {
			result += (float64)(r.DaysRented - 3) * 1.5
		}
	}

	return result
}
