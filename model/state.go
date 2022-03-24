package model

import (
	"time"

	"gorm.io/gorm"
)

type State struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	CountryId int       `json:"country_id"`
	IsActive  int       `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (s *State) Create(db *gorm.DB) error {
	res := db.Create(&s)
	return res.Error
}
func (s *State) Get(db *gorm.DB) error {
	res := db.Find(&s)
	return res.Error
}
func (s *State) Update(db *gorm.DB, updatedFileds map[string]interface{}) error {

	res := db.Model(&s).Updates(updatedFileds)
	return res.Error
}
func (s *State) Delete(db *gorm.DB) error {
	res := db.Delete(&s)
	return res.Error
}

func GetAllStates(db *gorm.DB) []State {
	var states []State
	db.Find(&states)
	return states
}
