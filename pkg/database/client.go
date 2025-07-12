package database

import (
	"beTestBookingToGo/internal/customer/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	// Instance.AutoMigrate(
	// 	&entities.Nationality{},
	// 	&entities.Customer{},
	// 	&entities.FamilyList{},
	// )

	if err := Instance.AutoMigrate(&entities.Customer{}); err != nil {
		log.Fatalf("Migrating Customer failed: %v", err)
	}

	// 2. Kemudian Nationality
	if err := Instance.AutoMigrate(&entities.Nationality{}); err != nil {
		log.Fatalf("Migrating Nationality failed: %v", err)
	}

	// 3. Lalu FamilyList
	if err := Instance.AutoMigrate(&entities.FamilyList{}); err != nil {
		log.Fatalf("Migrating FamilyList failed: %v", err)
	}

	log.Println("Database Migration Completed...")
	seedNationalities()
}
