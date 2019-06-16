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
	result := fmt.Sprintf("Rental record for %s\n", c.Name)
	for _, rental := range c.Rentals {
		// show figures for this rental
		result += fmt.Sprintf("\t%s\t%f\n", rental.Movie.Title, rental.Charge())
	}
	// add footer lines
	result += fmt.Sprintf("Amount owed is %f\n", c.totalCharge())
	result += fmt.Sprintf("You earned %d frequent renter points", c.totalFrequentRenterPoints())
	return result
}

func (c *Customer) HTMLStatement() string {
	result := fmt.Sprintf("<h1>Rentals for <em>%s</em></h1><p>\n", c.Name)
	for _, rental := range c.Rentals {
		result += fmt.Sprintf("%s: %f<br>\n", rental.Movie.Title, rental.Charge())
	}
	// add footer lines
	result += fmt.Sprintf("<p>You owe <em>%f</em><p>\n", c.totalCharge())
	result += fmt.Sprintf("On this rental you earned <em>%d</em> frequent renter points<p>", c.totalFrequentRenterPoints())
	return result
}

func (c *Customer) totalCharge() float64 {
	result := 0.0
	for _, rental := range c.Rentals {
		// show figures for this rental
		result += rental.Charge()
	}
	return result
}

func (c *Customer) totalFrequentRenterPoints() int {
	result := 0
	for _, rental := range c.Rentals {
		// show figures for this rental
		result += rental.FrequentRenterPoints()
	}
	return result
}
