package decoration

import "testing"

func TestCustomerBug(t *testing.T) {
	cus := Customer{}
	cus.Buy()
}
