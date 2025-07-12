package entities

import "time"

type Customer struct {
	CstID         int       `gorm:"column:cst_id;primaryKey;autoIncrement"`
	NationalityID int       `gorm:"column:nationality_id;null"`
	CstName       string    `gorm:"column:cst_name;size:50;not null"`
	CstDob        time.Time `gorm:"column:cst_dob;not null"`
	CstPhoneNum   string    `gorm:"column:cst_phoneNum;size:20;not null"`
	CstEmail      string    `gorm:"column:cst_email;size:50;not null"`

	Nationality   Nationality `gorm:"references:NationalityID"`
	FamilyList    []FamilyList `gorm:"foreignKey:CstID"`
}