package model

import (
	"time"

	"gorm.io/gorm"
)

type Hospital struct {
	ID                       int       `json:"id"`
	Name                     string    `json:"name"`
	Address                  string    `json:"address"`
	City                     string    `json:"city"`
	State_Id                 int       `json:"state_id"`
	Country_Id               int       `json:"country_id"`
	Is_Active                int       `json:"is_active"`
	Contact_Number           string    `json:"contact_number"`
	Fax_Number               string    `json:"fax_number"`
	Emergency_Contact_Number string    `json:"emergency_contact_number"`
	Email                    string    `json:"email"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (h *Hospital) Create(db *gorm.DB) error {
	res := db.Create(&h)
	return res.Error
}
func (h *Hospital) Get(db *gorm.DB) error {
	res := db.Find(&h)
	return res.Error
}
func (h *Hospital) Update(db *gorm.DB, updatedFileds map[string]interface{}) error {

	res := db.Model(&h).Updates(updatedFileds)
	return res.Error
}
func (h *Hospital) Delete(db *gorm.DB) error {
	res := db.Delete(&h)
	return res.Error
}

func GetAllHospitals(db *gorm.DB) []Hospital {
	var hospitals []Hospital
	db.Find(&hospitals)
	return hospitals
}
