package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"beTestBookingToGo/pkg/database"
	"beTestBookingToGo/internal/customer/entities"
	"beTestBookingToGo/internal/customer/entities/dto"
	"beTestBookingToGo/pkg/response"
)


func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON: " + err.Error())
		return
	}

	cstDob, err := time.Parse("2006-01-02", req.CstDob)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid date format for CstDob")
		return
	}

	customer := dto.Customer{
		CstName:       req.CstName,
		CstDob:        cstDob,
		NationalityID: req.NationalityID,
	}

	// Tambah FamilyList
	for _, fam := range req.Family {
		customer.FamilyList = append(customer.FamilyList, entities.FamilyList{
			FlName: fam.FlName,
			FlDob:  fam.FlDob,
			FlRelation: "Other", // Default relation
		})
	}

	if err := database.Instance.Create(&customer).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, customer)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	// Default pagination params
	page := 1
	limit := 10

	// Ambil query parameter
	queryPage := r.URL.Query().Get("page")
	queryLimit := r.URL.Query().Get("limit")

	if queryPage != "" {
		if p, err := strconv.Atoi(queryPage); err == nil && p > 0 {
			page = p
		}
	}
	if queryLimit != "" {
		if l, err := strconv.Atoi(queryLimit); err == nil && l > 0 {
			limit = l
		}
	}

	// Hitung offset
	offset := (page - 1) * limit

	var customers []entities.Customer

	result := database.Instance.
		Preload("FamilyList").
		Preload("Nationality").
		Limit(limit).
		Offset(offset).
		Order("cst_id DESC").
		Find(&customers)

	if result.Error != nil {
		response.Error(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	response.Success(w, customers)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var customer dto.Customer
	if err := database.Instance.Preload("FamilyList").First(&customer, id).Error; err != nil {
		response.Error(w, http.StatusNotFound, "Customer not found")
		return
	}
	response.Success(w, customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var existing dto.Customer
	if err := database.Instance.Preload("FamilyList").First(&existing, id).Error; err != nil {
		response.Error(w, http.StatusNotFound, "Customer not found")
		return
	}

	var req dto.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON: " + err.Error())
		return
	}

	cstDob, err := time.Parse("2006-01-02", req.CstDob)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid date format for CstDob")
		return
	}

	existing.CstName = req.CstName
	existing.CstDob = cstDob
	existing.NationalityID = req.NationalityID

	// Hapus FamilyList lama
	database.Instance.Where("cst_id = ?", existing.CstID).Delete(&entities.FamilyList{})

	// Tambah FamilyList baru
	existing.FamilyList = []entities.FamilyList{}
	for _, fam := range req.Family {
		existing.FamilyList = append(existing.FamilyList, entities.FamilyList{
			FlName:     fam.FlName,
			FlDob:      fam.FlDob,
			FlRelation: "Other",
		})
	}

	if err := database.Instance.Session(&gorm.Session{FullSaveAssociations: true}).Save(&existing).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, existing)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var customer dto.Customer
	if err := database.Instance.First(&customer, id).Error; err != nil {
		response.Error(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Hapus FamilyList terkait
	database.Instance.Where("cst_id = ?", customer.CstID).Delete(&entities.FamilyList{})

	// Hapus Customer
	if err := database.Instance.Delete(&customer).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, map[string]string{"message": "Customer deleted successfully"})
}