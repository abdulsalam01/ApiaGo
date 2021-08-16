package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey;autoIncrement:true" form:"id" json:"id"`
	Username  string `gorm:"NOT NULL" json:"username"`
	Password  string `gorm:"NOT NULL" json:"password"`
	CreateAt  time.Time
	UpdatedAt time.Time
	Contact   []Contact `gorm:"foreignKey:id"`
}

func (u *User) FindAll(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Find(&users).Limit(100)

	if err.Error != nil {
		return &[]User{}, err.Error
	}

	return &users, nil
}

func (u *User) FindById(db gorm.DB, id int) (*User, error) {
	err := db.Debug().First(&User{}, id).Take(&u)

	if err.Error != nil {
		return &User{}, err.Error
	}

	return u, nil
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) Update(db *gorm.DB, id int) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Updates(User{
		Username:  u.Username,
		Password:  u.Password,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	})

	if err.Error != nil {
		return &User{}, err.Error
	}

	res, _ := u.FindById(*db, id)
	return res, nil
}

func (u *User) Delete(db *gorm.DB, id int) (int, error) {
	res := db.Debug().Take(&User{}).Delete(&User{}, id)

	if res.Error != nil {
		return -1, res.Error
	}

	return int(res.RowsAffected), nil
}
