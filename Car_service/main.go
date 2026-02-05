package main

import (
	"encoding/json"
	"net/http"
)

type CarResponse struct {
	Brand     string `json:"brand"`
	OwnerName string `json:"owner_name"` // Эти данные мы получим от другого сервиса
}

func getCarInfo(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://user-service:8081/internal/user?id=1")
	if err != nil {
		http.Error(w, "User service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// 2. Распаковываем ответ от User Service
	var user struct {
		Name string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&user)

	// 3. Формируем общий ответ
	result := CarResponse{
		Brand:     "Tesla",
		OwnerName: user.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/car-details", getCarInfo)
	http.ListenAndServe(":8082", nil)
}
