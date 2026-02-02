package main

import (
	"github/prac-soc/db/conn"
	"github/prac-soc/db/store"
	api "github/prac-soc/handle"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conn.ConnectDB()

	queries := db.New(conn.Pool)

	handler := &api.Handler{
		Queries: queries,
	}

	r := mux.NewRouter()
	r.HandleFunc("/teacher", handler.RegisterTeacher).Methods("POST")
	r.HandleFunc("/teacher/login", handler.LoginTeacher).Methods("POST")
	r.HandleFunc("/teacher/quizzes", handler.CreateRoom).Methods("POST")

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
