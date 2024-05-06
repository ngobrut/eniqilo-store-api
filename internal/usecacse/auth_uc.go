package usecacse

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/jwt"
	"github.com/ngobrut/eniqlo-store-api/pkg/util"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error) {
	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	data := &model.User{
		Name:     req.Name,
		Email:    req.Email,
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
		Name:        data.Name,
		Email:       data.Email,
		CreatedAt:   data.CreatedAt,
		AccessToken: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error) {
	var res *response.AuthResponse

	// todo:

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	var res *model.User

	// todo:

	return res, nil
}
