package customers

import (
	"fmt"

	"../movies"
)

type Customer struct {
	Name string
	Rentals []movies.Rental
}

func NewCustomer(name string) *Customer {
	return &Customer{name, make([]movies.Rental, 5)}
}

func (c *Customer) Statement() string {
	var totalAmount float64 = 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental record for %s\n", c.Name)
	for _, rental := range c.Rentals {
		var thisAmount float64 = 0
		switch (rental.Movie.PriceCode) {
		case movies.REGULAR:
			thisAmount += 2
			if rental.DaysRented > 2 {
				thisAmount += (float64)(rental.DaysRented - 2) * 1.5
			}
		case movies.NEW_RELEASE:
			thisAmount += (float64)(rental.DaysRented) * 3
		case movies.CHILDRENS:
			thisAmount += 1.5
			if rental.DaysRented > 3 {
				thisAmount += (float64)(rental.DaysRented - 3) * 1.5
			}
		default:
		}

		// add frequent renter points
		frequentRenterPoints++
		// add bonus for a two dat new release rental
		if (rental.Movie.PriceCode == movies.NEW_RELEASE &&
			rental.DaysRented > 1) {
			frequentRenterPoints++
		}

		// show figures for this rental
		result += fmt.Sprintf("\t%s\t%f\n", rental.Movie.Title, thisAmount)
		totalAmount += thisAmount
	}
	// add footer lines
	result += fmt.Sprintf("Amount owed is %f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points")
	return result
}