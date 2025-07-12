package controllers

import (
	"encoding/json"
	"net/http"

	"beTestBookingToGo/pkg/database"
	"beTestBookingToGo/internal/customer/entities"
	"beTestBookingToGo/pkg/response"
)


func Getnationality(w http.ResponseWriter, r *http.Request) {
	var nationalities []entities.Nationality

	result := database.Instance.
		Find(&nationalities)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Error.Error()})
		return
	}

	response.Success(w, nationalities)
}