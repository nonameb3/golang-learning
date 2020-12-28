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
	bank := person{
		firstName: "Waraphon",
		lastName:  "Roonnapai",
		contact: contactInfo{
			email:   "roonnapai.dev@gmail.com",
			zipCode: 10220,
		},
	}

	// pointer origin
	// pointer := &bank
	// pointer.updateFristName("Bank'Waraphon")

	bank.updateFristName("Bank'Waraphon")
	bank.print()
}

// func origin origin
// func (pPointer *person) updateFristName(name string) {
// 	(*pPointer).firstName = name
// }

func (p *person) updateFristName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
