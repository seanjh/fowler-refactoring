package customers

import (
	"testing"

	"../movies"
)

func TestStatement(t *testing.T) {
	cases := []struct {
		customer Customer
		statement string
	}{
		{
			Customer{"Foo", make([]movies.Rental, 0)},
			"Rental record for Foo\n" +
			"Amount owed is 0.000000\n" +
			"You earned 0 frequent renter points",
		},
		{
			Customer{
				"Bar",
				[]movies.Rental{
					movies.Rental{&movies.Movie{"The Thing", movies.NEW_RELEASE}, 1},
				},
			},
			"Rental record for Bar\n" +
			"\tThe Thing\t3.000000\n" +
			"Amount owed is 3.000000\n" +
			"You earned 1 frequent renter points",
		},
		{
			Customer{
				"Baz",
				[]movies.Rental{
					movies.Rental{&movies.Movie{"Paddington 2", movies.CHILDRENS}, 1},
					movies.Rental{&movies.Movie{"Frozen", movies.CHILDRENS}, 30},
				},
			},
			"Rental record for Baz\n" +
			"\tPaddington 2\t1.500000\n" +
			"\tFrozen\t42.000000\n" +
			"Amount owed is 43.500000\n" +
			"You earned 2 frequent renter points",
		},
	}

	for _, c := range cases {
		if actual := c.customer.Statement(); actual != c.statement {
			t.Errorf("Wrong Customer statement\n---actual---\n%s\n---expected---\n%s\n---\n", actual, c.statement)
		}
	}
}

func TestHTMLStatement(t *testing.T) {
	cases := []struct {
		customer Customer
		statement string
	}{
		{
			Customer{"Foo", make([]movies.Rental, 0)},
			"<h1>Rentals for <em>Foo</em></h1><p>\n" +
			"<p>You owe <em>0.000000</em><p>\n" +
			"On this rental you earned <em>0</em> frequent renter points<p>",
		},
		{
			Customer{
				"Bar",
				[]movies.Rental{
					movies.Rental{&movies.Movie{"The Thing", movies.NEW_RELEASE}, 1},
				},
			},
			"<h1>Rentals for <em>Bar</em></h1><p>\n" +
			"The Thing: 3.000000<br>\n" +
			"<p>You owe <em>3.000000</em><p>\n" +
			"On this rental you earned <em>1</em> frequent renter points<p>",
		},
		{
			Customer{
				"Baz",
				[]movies.Rental{
					movies.Rental{&movies.Movie{"Paddington 2", movies.CHILDRENS}, 1},
					movies.Rental{&movies.Movie{"Frozen", movies.CHILDRENS}, 30},
				},
			},
			"<h1>Rentals for <em>Baz</em></h1><p>\n" +
			"Paddington 2: 1.500000<br>\n" +
			"Frozen: 42.000000<br>\n" +
			"<p>You owe <em>43.500000</em><p>\n" +
			"On this rental you earned <em>2</em> frequent renter points<p>",
		},
	}

	for _, c := range cases {
		if actual := c.customer.HTMLStatement(); actual != c.statement {
			t.Errorf("Wrong Customer statement\n---actual---\n%s\n---expected---\n%s\n---\n", actual, c.statement)
		}
	}
}
