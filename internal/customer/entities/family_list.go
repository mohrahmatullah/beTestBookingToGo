package entities

type FamilyList struct {
	FlID      int    `gorm:"column:fl_id;primaryKey;autoIncrement"`
	CstID     int    `gorm:"column:cst_id;null"`
	FlRelation string `gorm:"column:fl_relation;size:50;null"`
	FlName    string `gorm:"column:fl_name;size:50;not null"`
	FlDob     string `gorm:"column:fl_dob;size:50;not null"`
}