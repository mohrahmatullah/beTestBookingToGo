package controllers

import (
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
		response.Error(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	response.Success(w, nationalities)
}