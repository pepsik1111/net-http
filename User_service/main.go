// services/user-service/main.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	// Имитируем базу данных
	user := User{ID: id, Name: "Ivan Go-developer"}

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/internal/user", getUser)
	http.ListenAndServe(":8081", nil)
}
