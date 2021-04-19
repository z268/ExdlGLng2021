package main

type TaskNotFound struct{}

func (t *TaskNotFound) Error() string {
	return "Task not found"
}

func createTask(task Task) (Task, error) {
	var query = "INSERT INTO tasks (`name`, `description`, `duedate`, `status`) VALUES (?, ?, ?, ?)"
	res, err := DB.Exec(query, task.Name, task.Description, task.DueDate, task.Status)
	if err != nil {
		return task, err
	}

	task.Id, err = res.LastInsertId()
	if err != nil {
		return task, err
	}

	return task, nil
}

func getTask(id int64) (Task, error) {
	var task Task

	query := "SELECT `id`, `name`, `description`, `duedate`, `status` FROM `tasks` WHERE id = ?"
	rows, err := DB.Query(query, id)
	if err != nil {
		return task, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.DueDate, &task.Status)
		return task, err
	}
	return task, &TaskNotFound{}
}

func getTasks() ([]Task, error) {
	var query = "SELECT `id`, `name`, `description`, `duedate`, `status` FROM `tasks`"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
