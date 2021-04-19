package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	urlPrefix    = "/tasks/"
	listenAddr   = ":8000"
)


func router(res http.ResponseWriter, req *http.Request) {
	path := strings.Trim(strings.TrimPrefix(req.URL.Path, urlPrefix), "/")
	id, idErr := strconv.ParseInt(path, 10, 64)

	isList := len(path) == 0
	hasId := !isList && idErr == nil && id > 0

	switch {
	case isList && req.Method == "GET":
		TaskListView(res)
	case isList && req.Method == "POST":
		TaskCreateView(res, req)

	case hasId && req.Method == "GET":
		TaskRetrieveView(res, id)
	case hasId && req.Method == "DELETE":
		TaskDeleteView(res, id)
	case hasId && req.Method == "PUT":
		TaskUpdateView(res, req, id)

	default:
		http.Error(res, "Page not found", http.StatusNotFound)
	}
}


func main() {
	var err error
	DB, err = sql.Open("mysql", "root:election@/todo?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	fmt.Println("Running HTTP on address", listenAddr)

	http.HandleFunc(urlPrefix, router)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
