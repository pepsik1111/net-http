package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	IDuser   int    `json:"id_user"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	IDorders []int  `json:"id_orders,omitempty"`
}

var NextIDUser = 1

type Car struct {
	Brand  string `json:"brand"`
	Number string `json:"number"`
}

type Order struct {
}

var users []User
var cars []Car

func createCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метот не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newCar Car

	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Ошибка в json", http.StatusBadRequest)
		return
	}

	cars = append(cars, newCar)

	fmt.Printf("Добваленна новая машина %+v\n", newCar)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Машина упешно создана\n"))

}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "ошибка json", http.StatusBadRequest)
		return
	}
	users = append(users, newUser)

	fmt.Printf("Добавлен польователь %+v\n", newUser)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь создан\n"))

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/user", getUser)

	println("Сервер запущен на http://localhost:8080/car")

	http.HandleFunc("/car", getCar)
	http.HandleFunc("/car/create", createCar)
	http.HandleFunc("/user/create", createUser)
	http.ListenAndServe(":8080", nil)
}
