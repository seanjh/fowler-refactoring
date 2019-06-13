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
		thisAmount := c.amountFor(rental)

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
	result += fmt.Sprintf("You earned %d frequent renter points", frequentRenterPoints)
	return result
}

func (c *Customer) amountFor(rental movies.Rental) float64 {
	result := 0.0
	switch (rental.Movie.PriceCode) {
	case movies.REGULAR:
		result += 2
		if rental.DaysRented > 2 {
			result += (float64)(rental.DaysRented - 2) * 1.5
		}
	case movies.NEW_RELEASE:
		result += (float64)(rental.DaysRented) * 3
	case movies.CHILDRENS:
		result += 1.5
		if rental.DaysRented > 3 {
			result += (float64)(rental.DaysRented - 3) * 1.5
		}
	}

	return result
}
