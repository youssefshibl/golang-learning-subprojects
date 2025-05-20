package src

import (
	"errors"
	"fmt"
	"time"
)

func MakeNewTask(id int64, desc string) *Task {

	return &Task{
		Id:        id,
		Desc:      desc,
		Status:    Todo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func AddTask(desc string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if Some(tasks, func(t Task) bool {
		return t.Desc == desc
	}) {
		return TaskNameExist
	}
	var NewTaskId int64 = 1
	if len((tasks)) > 0 {
		NewTaskId = int64(tasks[len(tasks)-1].Id) + 1
	}
	task := MakeNewTask(NewTaskId, desc)
	tasks = append(tasks, *task)
	s, err := WriteTasksToFile(tasks)
	if !s {
		return err
	}

	fmt.Println("Task added successfully wiht id: ", task.Id)

	return nil

}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	exist := find(&tasks, (func(t Task) bool {
		return t.Id == id
	}))
	if exist == nil {
		return errors.New("this id not exist")
	}

	tasks = filter(tasks, func(t Task) bool {
		return t.Id != id
	})
	s, err := WriteTasksToFile(tasks)
	if !s {
		return err
	}
	fmt.Println("Task deleted successfully wiht id: ", id)
	return nil
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	task := find(&tasks, func(t Task) bool {
		return t.Id == id
	})
	if task == nil {
		return errors.New("task with this id not exsit ")
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	s, err := WriteTasksToFile(tasks)
	if !s {
		return err
	}
	fmt.Println("Task updated successfully wiht id: ", id)
	return nil

}

func UpdateTaskSDescription(id int64, desc string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	task := find(&tasks, func(t Task) bool {
		return t.Id != id
	})
	if task == nil {
		return errors.New("task with this id not exsit ")
	}

	task.Desc = desc
	task.UpdatedAt = time.Now()

	s, err := WriteTasksToFile(tasks)
	if !s {
		return err
	}
	fmt.Println("Task updated successfully wiht id: ", id)
	return nil

}

func ListAllTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if status != "" && status != "All" {
		tasks = filter(tasks, func(t Task) bool {
			return t.Status == status
		})
	}

	for _, task := range tasks {
		fmt.Println("Tasks")
		fmt.Printf("Id: %d , Des: %s , Status: %s , CreatedAt: %s , UpdatedAt: %s", task.Id, task.Desc, task.Status, task.CreatedAt, task.UpdatedAt)
	}
	return nil

}
