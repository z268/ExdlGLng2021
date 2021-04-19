package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

type Task struct {
	Id          int64
	Name        string
	Description string
	DueDate     time.Time
	Status      uint8
}

func (task *Task) delete() error {
	_, err := DB.Exec("DELETE FROM tasks WHERE `id` = ?", task.Id)
	return err
}

func (task *Task) save() error {
	query := "UPDATE tasks SET `name` = ?, `description` = ?, `duedate` = ?, `status` = ? WHERE `id` = ?"
	_, err := DB.Exec(query, task.Name, task.Description, task.DueDate, task.Status, task.Id)
	return err
}
