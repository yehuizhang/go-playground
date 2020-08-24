package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func main() {
	alex := person{
		"Yehui",
		"Zhang",
		contactInfo{
			"yehuizhang@outlook.com",
			95035,
		},
	}

	alex.updateName("Kyle")
	alex.print()
}
