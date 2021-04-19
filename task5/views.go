package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TaskListView(res http.ResponseWriter) {
	taskList, err := getTasks()
	if err != nil {
		http.Error(res, fmt.Sprintf("Getting task error: %v", err), http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(taskList)
}

func TaskCreateView(res http.ResponseWriter, req *http.Request) {
	task := Task{}
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(res, fmt.Sprintf("Format error: %v", err), http.StatusBadRequest)
		return
	}

	task, err = createTask(task)
	if err != nil {
		http.Error(res, fmt.Sprintf("Insert error: %v", err), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(task)
}

func TaskRetrieveView(res http.ResponseWriter, id int64) {
	task, err := getTask(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&task)
}

func TaskDeleteView(res http.ResponseWriter, id int64) {
	task, err := getTask(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	err = task.delete()
	if err != nil {
		http.Error(res, fmt.Sprintf("Delete error: %v", err), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func TaskUpdateView(res http.ResponseWriter, req *http.Request, id int64) {
	task, err := getTask(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(res, fmt.Sprintf("Format error: %v", err), http.StatusBadRequest)
		return
	}

	err = task.save()
	if err != nil {
		http.Error(res, fmt.Sprintf("Save error: %v", err), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(task)
}
