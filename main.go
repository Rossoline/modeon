package main

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
}

type Person struct {
	Name        string
	PhoneNumber string
	TaskManager TaskManager
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
func (tm *TaskManager) AddTask(t Task) {
	if slices.Contains(tm.Tasks, t) {
		return
	}
	tm.Tasks = append(tm.Tasks, t)
}

// EditTask : Edits an existing task based on its title
func (tm *TaskManager) EditTask(oldTitle string, newTask Task) bool {
	for i, task := range tm.Tasks {
		if task.Title == oldTitle {
			tm.Tasks[i] = newTask
			return true
		}
	}
	return false
}

// DeleteTask : Removes a task based on its title.
func (tm *TaskManager) DeleteTask(title string) bool {
	for i, task := range tm.Tasks {
		if task.Title == title {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

// MarkTaskAsCompleted : Marks a task as completed based on its title.
func (tm *TaskManager) MarkTaskAsCompleted(title string) bool {
	for i, task := range tm.Tasks {
		if task.Title == title {
			tm.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

// ListAllTasks : Lists all tasks of a person.
func (tm *TaskManager) ListAllTasks() {
	for _, task := range tm.Tasks {
		fmt.Printf("Title: %s, Description: %s, DueDate: %s, Completed: %t\n", task.Title, task.Description, task.DueDate.Format("2006-01-02"), task.Completed)
	}
}

// FilterTasks : Filters tasks by their completion status.
func (tm *TaskManager) FilterTasks(completed bool) []Task {
	var filteredTasks []Task
	for _, task := range tm.Tasks {
		if task.Completed == completed {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func (tm *TaskManager) FilterTasksByDate(dueDate time.Time) []Task {
	var filteredTasks []Task
	for _, task := range tm.Tasks {
		if task.DueDate.Equal(dueDate) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
