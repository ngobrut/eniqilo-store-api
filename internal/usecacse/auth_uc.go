package usecacse

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
	"github.com/ngobrut/eniqlo-store-api/pkg/jwt"
	"github.com/ngobrut/eniqlo-store-api/pkg/util"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error) {
	if !strings.HasPrefix(req.Phone, "+") {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "invalid phone number format",
		})

		return nil, err
	}

	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	data := &model.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: pwd,
	}

	err = u.repo.CreateUser(ctx, data)
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: data.UserID.String(),
	}

	token, err := jwt.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{
		UserID:      data.UserID,
		Name:        data.Name,
		Phone:       data.Phone,
		AccessToken: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error) {
	if !strings.HasPrefix(req.Phone, "+") {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "invalid phone number format",
		})

		return nil, err
	}

	user, err := u.repo.GetOneUserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	err = util.ComparePwd([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "wrong password",
		})

		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: user.UserID.String(),
	}

	token, err := jwt.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		Phone:       user.Phone,
		AccessToken: token,
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	res, err := u.repo.GetOneUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
