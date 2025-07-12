package dto

import (
	"time"
	"beTestBookingToGo/internal/customer/entities"
)

type Customer struct {
	CstID         int       `gorm:"column:cst_id;primaryKey;autoIncrement"`
	NationalityID int       `gorm:"column:nationality_id;null"`
	CstName       string    `gorm:"column:cst_name;size:50;not null"`
	CstDob        time.Time `gorm:"column:cst_dob;not null"`
	CstPhoneNum   string    `gorm:"column:cst_phoneNum;size:20;not null"`
	CstEmail      string    `gorm:"column:cst_email;size:50;not null"`

	Nationality *entities.Nationality `gorm:"foreignKey:NationalityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	FamilyList  []entities.FamilyList `gorm:"foreignKey:CstID"`
}