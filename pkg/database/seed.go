package database

import (
	"errors"
	"beTestBookingToGo/internal/customer/entities"
	"log"

	"gorm.io/gorm"
)

func seedNationalities() {
    nationalities := []entities.Nationality{
        {NationalityName: "Indonesia", NationalityCode: "ID"},
        {NationalityName: "Malaysia", NationalityCode: "MY"},
        {NationalityName: "Singapore", NationalityCode: "SG"},
        {NationalityName: "United States", NationalityCode: "US"},
    }

    var count int64
    if err := Instance.Model(&entities.Nationality{}).Count(&count).Error; err != nil {
        log.Fatalf("Gagal menghitung nationality: %v", err)
    }

    if count == 0 {
        // Tabel kosong, insert semua
        if err := Instance.Create(&nationalities).Error; err != nil {
            log.Fatalf("Seeding Nationality gagal: %v", err)
        }
        log.Printf("Nationality seed inserted successfully (%d records).", len(nationalities))
        return
    }

    // Tabel sudah ada isinya, cek per record
    var toInsert []entities.Nationality

    for _, nat := range nationalities {
        var existing entities.Nationality
        err := Instance.Where("nationality_code = ?", nat.NationalityCode).First(&existing).Error

        if errors.Is(err, gorm.ErrRecordNotFound) {
            toInsert = append(toInsert, nat)
        } else if err != nil {
            log.Fatalf("Gagal cek Nationality %s: %v", nat.NationalityCode, err)
        }
        // else: sudah ada, skip
    }

    if len(toInsert) == 0 {
        log.Println("Nationality seed skipped (all records already exist).")
        return
    }

    if err := Instance.Create(&toInsert).Error; err != nil {
        log.Fatalf("Seeding Nationality failed: %v", err)
    }

    log.Printf("Nationality seed inserted successfully (%d records).", len(toInsert))
}

