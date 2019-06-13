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
	}

	for _, c := range cases {
		if actual := c.customer.Statement(); actual != c.statement {
			t.Errorf("Wrong Customer statement\n---\n%s\n---\n%s\n---\n", actual, c.statement)
		}
	}
}
