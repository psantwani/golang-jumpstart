package main

import (
	"fmt"
)

type profile struct {
	name       string
	address    string
	email      string
	contactnum string
}

func (p profile) getname() string {
	return p.name
}

func (p profile) getaddress() string {
	return p.address
}

func (p profile) getcontact() map[string]string {
	m := map[string]string{"email": p.email, "contactnum": p.contactnum}
	return m
}

func methodsdemo() {
	s := profile{
		name:       "Piyush Santwani",
		address:    "Mumbai",
		email:      "psantwani@gmail.com",
		contactnum: "555555555",
	}

	fmt.Println("Name: ", s.getname())
	fmt.Println("Address: ", s.getaddress())
	fmt.Println("Contact Info: ", s.getcontact())
}
