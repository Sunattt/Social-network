package handlers

import (
	"Sunat/internal/services"
	"Sunat/pkg/helpers"
	"Sunat/pkg/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Service *services.Services
	Logger  *zap.Logger
}

const (
	nicknameChangeInterval = 14 * 24 * time.Hour // 14 дней
	nameChangeInterval     = 14 * 24 * time.Hour // 14 дней
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var u *models.Users

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if errors.Is(err, io.EOF) {
			helpers.BadRequest(w, err, h.Logger)
			err = helpers.ResponseAnswer(w, "Incorrect data")
			if err != nil {
				helpers.InternalServerError(w, err, h.Logger)
			}
			return
		}
		log.Println("ERROR authorizate !!", err)
		helpers.InternalServerError(w, err, h.Logger)
		helpers.ResponseAnswer(w, "OPPSSS something went wrong")
		return
	}

	token, err := h.Service.GetLoginService(u)
	if err != nil {
		errAct := errors.New("User is bloced!")
		if errors.Is(err, errAct) {
			helpers.InternalServerError(w, err, h.Logger)
			err = helpers.ResponseAnswer(w, "User is not found!")
			if err != nil {
				helpers.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		helpers.BadRequest(w, err, h.Logger)
		return
	}

	sendToken := models.SendToken{
		Date:   time.Now(),
		Answer: "Authorization was successful!",
		Token:  token,
	}
	err = helpers.SendToken(w, &sendToken)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var user *models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		return
	}
	err = h.Service.CheckNikName(user.NikName, user.UpdatedAt, nicknameChangeInterval)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Username entered registration")
		return
	}

	err = h.Service.CheckName(user.Name, user.UpdatedAt, nameChangeInterval)

	err = h.Service.CheckPhone(user.Phone)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Phone entered registration")
		return
	}

	err = h.Service.CheckEmail(user.Email)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect E-mail entered registration")
		return
	}

	err = h.Service.CheckPassword(user.Password)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Password entered registration")
		return
	}

	token, err := h.Service.RegistrationUser(user)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		helpers.ResponseAnswer(w, "OPPSSS wrong while service")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(token)
	if err != nil {
		log.Printf("[ERROR] Can't Marshal Error JSON! Info: %v", err)
	}

	w.Write(respBytes)
	err = helpers.ResponseAnswer(w, "Registration complated successfully.")
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}
}
