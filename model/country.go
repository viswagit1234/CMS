package model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	IsActive  int       `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (c *Country) Create(db *gorm.DB) error {
	res := db.Create(&c)
	return res.Error
}
func (c *Country) Get(db *gorm.DB) error {
	res := db.Find(&c)
	return res.Error
}
func (c *Country) Update(db *gorm.DB, updatedFileds map[string]interface{}) error {

	res := db.Model(&c).Updates(updatedFileds)
	return res.Error
}
func (c *Country) Delete(db *gorm.DB) error {
	res := db.Delete(&c)
	return res.Error
}

func GetAllCountries(db *gorm.DB) []Country {
	var countries []Country
	db.Find(&countries)
	return countries
}
