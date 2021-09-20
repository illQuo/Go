package main

import "fmt"

type person struct {
	name    string
	age     int
	contact contactInfo
}

type contactInfo struct {
	email    string
	postCode int
}

func main() {
	bob := person{
		name: "Bob",
		age:  40,
		contact: contactInfo{
			email:    "yes@no.com",
			postCode: 4019,
		},
	}

	bob.updateName("Josh")
	bob.print()
}

func (p *person) updateName(name string) {
	p.name = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
