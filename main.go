package main

import "fmt"

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
	p.Task = append(p.Task, s)
}
