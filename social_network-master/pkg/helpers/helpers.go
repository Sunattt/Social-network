package helpers

import (
	"Sunat/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	lifetime   = 3 * time.Hour
	signingKey = "grkjk#4#%35FSFJLja#4353KSFjH"
	KeyUserId  = "userID"
)

func BadRequest(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info("the user entered incorrect data", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusBadRequest), 400)
}

func InternalServerError(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Error(http.StatusText(http.StatusInternalServerError), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
}

func Unauthorized(w http.ResponseWriter, logger *zap.Logger) {
	logger.Error(http.StatusText(http.StatusUnauthorized), zap.Error(errors.New("User isn't auth")))
	http.Error(w, http.StatusText(http.StatusUnauthorized), 401)
}

//func Forbidden(w http.ResponseWriter, err error, logger *zap.Logger) {
//	logger.Info(http.StatusText(http.StatusForbidden), zap.Error(err))
//	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
//}
//
//func NotFoundErr(w http.ResponseWriter, err error, logger *zap.Logger) {
//	logger.Info(http.StatusText(http.StatusNotFound), zap.Error(err))
//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//}

func ResponseAnswer(w http.ResponseWriter, report string) (err error) {
	answer := models.Answer{
		Date:   time.Now(),
		Answer: report,
	}
	myAnswer, err := json.MarshalIndent(answer, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(myAnswer)
	if err != nil {
		return err
	}
	return
}

func SendToken(w http.ResponseWriter, sendToken *models.SendToken) error {
	answer, err := json.MarshalIndent(sendToken, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(answer)
	if err != nil {
		return err
	}

	return nil
}

//
//func ResetContentServerError(w http.ResponseWriter, err error, logger *zap.Logger) {
//	logger.Info(http.StatusText(http.StatusResetContent), zap.Error(err))
//	http.Error(w, http.StatusText(http.StatusResetContent), http.StatusResetContent)
//}

func GetUserIdFromContext(r *http.Request) (id int, err error) {
	id, ok := r.Context().Value(KeyUserId).(int)
	if !ok {
		err = errors.New("User is not found!")
		return 0, err
	}
	return id, nil
}

func SendAnswer(w http.ResponseWriter, msg string) error {
	answer := models.Answer{Answer: msg}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(answer)
	if err != nil {
		return err
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
