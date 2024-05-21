package services

import (
	"log"
	"regexp"
)

// проверка на правильное написание E-mail адреса
func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		log.Println("Error valid email")
	}
	return match
}

// правильно написание ника
func IsNikNameValid(nikName string) bool {
	pattern := `^[a-z0-9._]*$`
	match, err := regexp.MatchString(pattern, nikName)
	if err != nil {
		log.Println("Error validtion nikName")
	}
	return match
}

func IsNameValid(nikName string) bool {
	pattern := `^[a-zA-Z0-9._]*$`
	match, err := regexp.MatchString(pattern, nikName)
	if err != nil {
		log.Println("Error validtion nikName")
	}
	return match
}

func IsAgeValid(age string) bool {
	pattern := "^[0-9]{2}.[0-9]{2}.[0-9]{4}$"
	match, err := regexp.MatchString(pattern, age)
	if err != nil {
		log.Println("Error validating age")
	}
	return match
}

func IsPhoneValid(phone string) bool {
	pattern := `^\+[992]-[0-9]{3}-[0-9]{7}$`

	match, err := regexp.MatchString(pattern, phone)
	if err != nil {
		log.Println("Ошибка при написании телефона", err)
		return false
	}

	return match
}

func IsPasswordValid(password string) bool {
	pattern := `^[a-zA-Z0-9._]*$`
	match, err := regexp.MatchString(pattern, password)
	if err != nil {
		log.Println("Error valid email")
	}
	return match
}
