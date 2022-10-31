package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Esta es una api hecha con gorm y mux"))
  }