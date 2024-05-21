package services

import (
	"Sunat/internal/repositories"
	"Sunat/pkg/jwt_token"
	"Sunat/pkg/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var (
	ErrNikNameUsed         = errors.New("NikName is used")
	ErrRegistration        = errors.New("incorrect value entered for registration")
	ErrNikNameRegistration = errors.New("username is taken by another user")
)

const (
	JWTSecretKey = "everything"
)

type Services struct {
	Repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Services {
	return &Services{
		Repository: repository,
	}
}

func (s *Services) CheckEmail(email string) error {
	//Если пользователь не ввёл данные
	if email == "" || !IsEmailValid(email) {
		return ErrRegistration
	}
	return nil
}

func (s *Services) CheckNikName(nikName string, lastUpdate time.Time, interval time.Duration) error {
	if nikName == "" || !IsNikNameValid(nikName) {
		return ErrRegistration
	}
	// Проверка интервала для изменения никнейма
	if time.Since(lastUpdate) < interval {
		return errors.New("nickname cannot be changed within the specified interval")
	}

	isFree := s.Repository.IsNikNameFree(nikName)
	if isFree {
		return ErrNikNameRegistration
	}
	return nil
}

func (s *Services) CheckName(name string, lastUpdate time.Time, interval time.Duration) error {
	if len(name) < 4 {
		return errors.New("invalid nickname length")
	}
	// Проверка интервала для изменения никнейма
	if time.Since(lastUpdate) < interval {
		return errors.New("nickname cannot be changed within the specified interval")
	}

	return nil
}

func (s *Services) CheckPhone(phone string) error {
	if phone == "" || !IsPhoneValid(phone) {
		return ErrRegistration
	}
	return nil
}

func (s *Services) CheckPassword(password string) error {
	if password == "" || !IsPasswordValid(password) {
		return ErrRegistration
	}
	return nil
}

func (s *Services) RegistrationUser(u *models.Users) (token string, err error) {
	//хеширование пароля (скрыть)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = s.Repository.AddUserToDB(u, string(hash))
	if err != nil {
		log.Println("Error while add user to db")
		return "", err
	}

	token, eror := jwt_token.CreateToken(u.NikName, JWTSecretKey)
	if eror != nil {
		return "", err
	}
	err = s.Repository.AddTokenToDb(u.Id, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Services) GetTokenService(u *models.Users) (token string, err error) {
	active, err := s.Repository.DBCheckActiveById(u.Id)
	if err != nil {
		return "", err
	}

	if !active {
		err = errors.New("User is blocked!")
		return "", err
	}

	userCount, err := s.Repository.DBCheckNikName(u.NikName)
	if err != nil {
		return "", err
	}

	switch {
	case userCount <= 0:
		err = errors.New("user is not found")
		return "", err
	case userCount > 1:
		err = errors.New("another similar user has been found, please contact technical support")
		return "", err
	}
	password, err := s.Repository.GetPasswordFromDb(u.Id)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		return "", err
	}

	token, err = jwt_token.CreateToken(u.NikName, JWTSecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Services) GetLoginService(user *models.Users) (token string, err error) {
	active, err := s.Repository.DBCheckActiveById(user.Id)
	if err != nil {
		return "", err
	}

	if !active {
		err = errors.New("User is blocked!")
		return "", err
	}

	userCount, err := s.Repository.DBCheckNikName(user.NikName)
	if err != nil {
		err = errors.New("User is not fount!")
		return "", err
	}
	switch {
	case userCount <= 0:
		err = errors.New("user is not found")
		return "", err
	case userCount > 1:
		err = errors.New("another similar user has been found, please contact technical support")
		return "", err
	}

	password, err := s.Repository.GetPasswordFromDb(user.Id)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		err = errors.New("wrong password")
		return "", err
	}

	token, err = jwt_token.CreateToken(user.NikName, JWTSecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
