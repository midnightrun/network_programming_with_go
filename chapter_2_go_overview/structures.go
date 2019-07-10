package main

import "fmt"

type Name struct {
	family   string
	personal string
}

type Email struct {
	kind    string
	address string
}

type Person struct {
	name  Name
	email []Email
}

func main() {
	person := Person{
		name: Name{
			family:   "Mustermann",
			personal: "Max",
		},
		email: []Email{Email{
			kind: "Work", address: "max@work.de",
		}, Email{
			kind: "Private", address: "max@private.de",
		}},
	}

	fmt.Printf("Hello %s %s.\n", person.name.family, person.name.personal)

	fmt.Println("E-Mail(s)")

	for _, v := range person.email {
		fmt.Printf("%s => %s\n", v.kind, v.address)
	}
}
