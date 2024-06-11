package main

import (
	"fmt"
	"slices"
)

type Person struct {
	Name        string
	PhoneNumber string
	Task        []Task
}

type Task struct {
	Date        string
	Description string
}

//Task Manager App

func main() {
	p := Person{}
	info(p)
}

func info(p Person) {
	fmt.Println(p)
}

func AddTask(p *Person, t Task) {
	if slices.Contains(p.Task, t) {
		return
	}
	p.Task = append(p.Task, t)
}
