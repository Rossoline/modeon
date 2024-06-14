package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	PhoneNumber string
	Task        []Task
}

type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
}

type TaskManager struct {
	Tasks []Task
}

//Task Manager App

func main() {
	fmt.Println("App started")
	p := Person{}
	info(p)
}

func info(p Person) {
	fmt.Println(p)
}

func (t Task) Equals(other Task) bool {
	return t.Title == other.Title && t.Description == other.Description && t.DueDate.Equal(other.DueDate) && t.Completed == other.Completed
}

// AddTask : Adds a task to the person's task list if it doesn't already exist
func AddTask(p *Person, t Task) {
	for _, task := range p.Task {
		if task.Equals(t) {
			return
		}
	}
	p.Task = append(p.Task, t)
}

// EditTask : Edits an existing task based on its title
func EditTask(p *Person, oldTitle string, newTask Task) bool {
	for i, task := range p.Task {
		if task.Title == oldTitle {
			p.Task[i] = newTask
			return true
		}
	}
	return false
}

// DeleteTask : Removes a task based on its title.
func DeleteTask(p *Person, title string) bool {
	for i, task := range p.Task {
		if task.Title == title {
			p.Task = append(p.Task[:i], p.Task[i+1:]...)
			return true
		}
	}
	return false
}

// MarkTaskAsCompleted : Marks a task as completed based on its title.
func MarkTaskAsCompleted(p *Person, title string) bool {
	for i, task := range p.Task {
		if task.Title == title {
			p.Task[i].Completed = true
			return true
		}
	}
	return false
}

// ListAllTasks : Lists all tasks of a person.
func ListAllTasks(p *Person) {
	for _, task := range p.Task {
		fmt.Printf("Title: %s, Description: %s, DueDate: %s, Completed: %t\n", task.Title, task.Description, task.DueDate.Format("2006-01-02"), task.Completed)
	}
}

// FilterTasks : Filters tasks by their completion status.
func FilterTasks(p *Person, completed bool) []Task {
	var filteredTasks []Task
	for _, task := range p.Task {
		if task.Completed == completed {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
