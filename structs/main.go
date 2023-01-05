package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// var pallab person
	// pallab.firstName = "Pallab"
	// pallab.lastName = "Kumar"
	// pallab.contact.email = "pallab@gmail.com"
	// pallab.contact.zipCode = 12345

	pallab := person{
		firstName: "Pallab",
		lastName:  "Kumar",
		contact: contactInfo{
			email:   "pallab@gmail.com",
			zipCode: 12345,
		},
	}
	pallab.updateName("Pallabito")
	pallab.print()
}

func (p person) updateName(updatedName string) {
	p.firstName = updatedName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
