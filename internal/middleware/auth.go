package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	custom_jwt "github.com/ngobrut/eniqlo-store-api/pkg/jwt"
)

func Authorize(secret string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := GetTokenFromHeader(r)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			res, err := jwt.ParseWithClaims(token, &custom_jwt.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			claims, ok := res.Claims.(*custom_jwt.CustomClaims)
			if !ok && !res.Valid {
				response.UnauthorizedError(w)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, claims.UserID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("token is empty")
	}

	token := strings.Split(header, " ")
	if len(token) < 2 {
		return "", errors.New("token is invalid")
	}

	return token[1], nil
}

func ParseWithoutVerified(token string) *custom_jwt.CustomClaims {
	res, _, err := new(jwt.Parser).ParseUnverified(token, &custom_jwt.CustomClaims{})
	if err != nil {
		return nil
	}

	claims, ok := res.Claims.(*custom_jwt.CustomClaims)
	if ok && claims.ID != "" {
		return claims
	}

	return nil
}
