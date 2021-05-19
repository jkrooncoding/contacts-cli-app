package main

type contact struct {
	firstName    string
	lastName     string
	emailAddress string
	phoneNumber  string
	socials      social
}

type social struct {
	fb string
	tw string
	ig string
}
