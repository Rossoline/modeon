package main

import (
	"fmt"
	"slices"
)

type Person struct {
	Name        string
	PhoneNumber string
	Task        []string
}

//Task Manager App

func main() {

}

func info(p Person) {
	fmt.Println(p)
}

func AddTask(p *Person, s string) {
	if slices.Contains(p.Task, s) {
		return
	}
	p.Task = append(p.Task, s)
}
