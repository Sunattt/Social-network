package repositories

import (
	"Sunat/pkg/models"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{Database: database}
}

// return true if Login is free, else - false
func (r *Repository) IsNikNameFree(nikName string) bool {
	var user models.Users
	amountOfChar := r.Database.Table("users").
		Where("nik_name = ?", nikName).First(&user).RowsAffected
	if amountOfChar != 0 {
		return false
	}
	return true
}

func (r *Repository) AddUserToDb(user *models.Users) error {
	err := r.Database.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddUserToDB(u *models.Users, hash string) error {
	u.Password = hash
	u.CreatedAt = time.Now()
	return r.Database.Create(u).Error
}

func (r *Repository) AddTokenToDb(userId uint, token string) error {
	var dbToken models.Tokens
	err := r.Database.First(&dbToken, "user_id = ?", userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	dbToken.Token = token
	dbToken.UpdatedAt = time.Now()
	dbToken.UserID = userId

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Database.Create(&dbToken).Error
	} else {
		return r.Database.Save(&dbToken).Error
	}
}

func (r *Repository) DBCheckNikName(nik string) (userCount int64, err error) {
	err = r.Database.Model(&models.Users{}).Where("nik_name = ? AND active = true", nik).Count(&userCount).Error
	if err != nil {
		log.Println("[ERROR] during check nik and pass !!!")
		return
	}
	return
}

func (r *Repository) DBCheckActiveById(id uint) (active bool, err error) {
	var user models.Users
	err = r.Database.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return user.Active, nil
}

func (r *Repository) GetPasswordFromDb(id uint) (password string, err error) {
	var user models.Users
	err = r.Database.Select("password").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}
	return user.Password, nil
}
