package http

import (
	"beTestBookingToGo/internal/customer/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	customerRoutes := router.PathPrefix("/api/customers").Subrouter()
	customerRoutes.HandleFunc("", controllers.CreateCustomer).Methods("POST")
	customerRoutes.HandleFunc("", controllers.GetCustomers).Methods("GET")
	customerRoutes.HandleFunc("/{id}", controllers.GetCustomerByID).Methods("GET")
	customerRoutes.HandleFunc("/{id}", controllers.UpdateCustomer).Methods("PUT")
	customerRoutes.HandleFunc("/{id}", controllers.DeleteCustomer).Methods("DELETE")

	nationalityRoutes := router.PathPrefix("/api/nationalities").Subrouter()
	nationalityRoutes.HandleFunc("", controllers.Getnationality).Methods("GET")

	return router
}