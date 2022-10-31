package main

import (
	"log"
	"net/http"

	"github.com/Thrashy190/backend-go-proyect/db"
	"github.com/Thrashy190/backend-go-proyect/models"
	"github.com/Thrashy190/backend-go-proyect/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main (){
  
  err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

  db.Dbconnection()

  db.DB.AutoMigrate(models.Task{})
  db.DB.AutoMigrate(models.User{})

  r := mux.NewRouter()
  // home route
  r.HandleFunc("/",routes.HomeHandler)


  s := r.PathPrefix("/api").Subrouter()

	// tasks routes
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

  // users routes
  r.HandleFunc("/users",routes.GetUsersHandler).Methods("GET")
  r.HandleFunc("/users/{id}",routes.GetUserHandler).Methods("GET")
  r.HandleFunc("/users",routes.PostUserHandler).Methods("POST")
  r.HandleFunc("/users/{id}",routes.DeleteUserHandler).Methods("DELETE")
  
  http.ListenAndServe(":3000",r)
}