package observer

import "fmt"

type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) GetID() string {
	return c.id
}
