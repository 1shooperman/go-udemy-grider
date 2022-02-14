package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Foo"
	// alex.lastName = "Bar"

	var alexContact contactInfo
	alexContact.email = "foo@bar.com"
	alexContact.zipCode = 90210

	alex.contact = alexContact

	// fmt.Println(alex)
	alex.print()

	alexPointer := &alex
	alexPointer.updateName("Jim")

	alex.print()
}

// go is pass by value

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
