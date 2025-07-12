package entities

type Nationality struct {
	NationalityID   int    `gorm:"column:nationality_id;primaryKey;autoIncrement"`
	NationalityName string `gorm:"column:nationality_name;size:50;not null"`
	NationalityCode string `gorm:"column:nationality_code;size:2;not null"`
}