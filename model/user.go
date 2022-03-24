package model

import(
	 "time"
      "gorm.io/gorm"
)
type User struct {
	ID                  int           `json:"id"`
	ChaildName          string        `json:"chaild_name"`
	FatherName          string        `json:"fathername"`
	MotherName          string        `json:"mothername"`
	DateOfBirth         time.Time     `json:"date_of_birth"`
	Gender              int           `json:"gender"`
	MobleNumber         string        `json:"mobile_number"`
	Email               string        `json:"email"`
	HospitalId          int           `json:"hospital_id"`
	City                string        `json:"city"`
	Addess              string        `json:"address"`
	StateId             int           `json:"state_id"`
	Country_id          int           `json:"country_id"`
	AccessToken         string        `json:"access_token"`
	IsActive            int           `json:"is_active"`
	IsVerify            int           `json:"is_verify"`
    LastLoginIp         string        `json:"last_login_ip"`
	LastLoginTime       int           `json:"last_login_time"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
}    

func (u *User) Create(db *gorm.DB) error {
	res := db.Create(&u)
	return res.Error
}
func (u *User) Get(db *gorm.DB) error {
	res := db.Find(&u)
	return res.Error
}
func (u *User) Update(db *gorm.DB, updatedFileds map[string]interface{}) error {

	res := db.Model(&u).Updates(updatedFileds)
	return res.Error
}
func (u *User) Delete(db *gorm.DB) error {
	res := db.Delete(&u)
	return res.Error
}

func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}
