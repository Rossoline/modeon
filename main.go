package main

import "fmt"

type Person struct {
	Name        string
	PhoneNumber string
	Task        []string
}

//Task Manager App

func main() {
	person1 := Person{"New person", "101", []string{}}
	info(person1)
}

func info(p Person) {
	fmt.Println(p)
}
