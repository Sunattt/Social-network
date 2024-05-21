package handlers

import (
	"Sunat/internal/services"
	"Sunat/pkg/helpers"
	"Sunat/pkg/jwt_token"
	"context"
	"log"
	"net/http"
)

const (
	KeyUserId = "userID"
)

// и возвращает обработчик, который будет выполнять проверку аутентификации клиента перед вызовом следующего обработчика.
func (h *Handler) Authentication(hand http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//заголовка запросов
		token := r.Header.Get("Authorization")

		nikName, valid, err := jwt_token.ValidToken(token, services.JWTSecretKey)
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			helpers.ResponseAnswer(w, "OOPPSS something went wrong!!!")
			return
		}

		if !valid {
			log.Println(valid)
			helpers.Unauthorized(w, h.Logger)
			helpers.ResponseAnswer(w, "Token has expired!!!")
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserId, nikName)
		r = r.WithContext(ctx)

		r.Header.Set("nik_name", nikName)
		hand.ServeHTTP(w, r)
		log.Println("User Authentication was successfully!")
	})
}
