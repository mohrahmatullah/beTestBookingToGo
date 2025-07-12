package dto

type CreateCustomerRequest struct {
	CstName       string              `json:"CstName"`
	CstDob        string              `json:"CstDob"`
	NationalityID int                 `json:"NationalityID"`
	Family        []CreateFamilyInput `json:"Family"`
}

type CreateFamilyInput struct {
	FlName string `json:"FlName"`
	FlDob  string `json:"FlDob"`
}