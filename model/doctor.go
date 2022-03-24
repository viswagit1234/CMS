package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Doctor struct {
	ID            int    `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	HospitalId    int    `json:"hospital_id"`
	Address       string `json:"address"`
	IsActive      int    `json:"is_active"`
	CreatedAt     int    `json:"created_at"`
	UpdatedAt     int    `json:"updated_at"`
}

func (d *Doctor) Create(db *gorm.DB) error {
	res := db.Create(&d)
	return res.Error
}
func (d *Doctor) Get(db *gorm.DB) error {
	res := db.Find(&d, d.ID)
	return res.Error
}
func (d *Doctor) Update(db *gorm.DB, updatedFileds map[string]interface{}) error {
	fmt.Println("we are here to test")
	fmt.Println(d)
	res := db.Model(&d).Where("id=?", d.ID).Updates(updatedFileds)
	return res.Error
}
func (d *Doctor) Delete(db *gorm.DB) error {
	res := db.Delete(&d, d.ID)
	return res.Error
}
func GetAllDoctors(db *gorm.DB) []Doctor {
	var doctors []Doctor
	db.Find(&doctors)
	return doctors
}
