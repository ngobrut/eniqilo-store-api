package util

import (
	"net/http"
	"strconv"

	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
	"golang.org/x/crypto/bcrypt"
)

func ComparePwd(hashed []byte, plain []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hashed, plain)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "wrong password",
		})

		return
	}

	return
}

func HashPwd(cost string, pwd []byte) (string, error) {
	salt, err := strconv.Atoi(cost)
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword(pwd, salt)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
