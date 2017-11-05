package main

import "fmt"

type userprofile interface {
	getname() string
	getaddress() string
	getcontact() map[string]string
}

func getUserInfo(u userprofile) {
	fmt.Println("Name: ", u.getname())
	fmt.Println("Address: ", u.getaddress())
	fmt.Println("Contact Info: ", u.getcontact())
}

func interfacedemo() {
	s := profile{
		name:       "Piyush Santwani",
		address:    "Mumbai",
		email:      "psantwani@gmail.com",
		contactnum: "555555555",
	}

	getUserInfo(s)
}
