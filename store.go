package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(32)"`
	Gender   int    `gorm:"column:gender;default:1"` // 1:male 2:female
	Age      int    `gorm:"column:age;default:18"`
	Nickname string `gorm:"column:nickname;"`
	Phone    string `gorm:"column:phone;"`
	Email    string `gorm:"column:email;"`
}

func (u *User) Create()(err error){
	err = GORM.Create(u).Error
	return err
}


func (u *User) Update()(err error) {
	err = GORM.Model(&User{}).Updates(u).Error
	return err
}

func (u *User)Get() (err error) {
	err = GORM.Where(u).First(u).Error
	return err
}

func (u *User)GetList() (users []User,err error) {
	err = GORM.Find(&users, u).Error
	return users,err
}

func (u *User)Delete() (err error) {
	err = GORM.Delete(u).Error
	return err
}
