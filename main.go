package main

import "fmt"

type Person struct {
	Name        string
	PhoneNumber string
}

//Task Manager App

func main() {
	person1 := Person{"New person", "101"}
	info(person1)
}

func info(p Person) {
	fmt.Println(p)
}
