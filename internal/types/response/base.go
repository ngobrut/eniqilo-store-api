package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_validator"
)

type JsonResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Error   *ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

const (
	CONTENT_TYPE_HEADER string = "Content-Type"
	CONTENT_TYPE_JSON   string = "application/json"
)

func OK(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JsonResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func Error(w http.ResponseWriter, err error) {
	v, isValidationErr := err.(custom_validator.ValidatorError)
	if isValidationErr {
		w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(JsonResponse{
			Message: "ValidationError",
			Error: &ErrorResponse{
				Code:    v.Code,
				Message: v.Message,
				Details: v.Details,
			},
		})

		return
	}

	e, isCustomErr := err.(*custom_error.CustomError)
	if !isCustomErr {
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Println(err.Error(), "[unhandled-error]")
		}

		w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JsonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Error: &ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: constant.HTTPStatusText(http.StatusInternalServerError),
			},
		})

		return
	}

	httpCode := http.StatusInternalServerError
	msg := constant.HTTPStatusText(httpCode)

	if e.ErrorContext != nil && e.ErrorContext.HTTPCode > 0 {
		httpCode = e.ErrorContext.HTTPCode
		msg = constant.HTTPStatusText(httpCode)

		if e.ErrorContext.Message != "" {
			msg = e.ErrorContext.Message
		}
	}

	w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(JsonResponse{
		Message: http.StatusText(httpCode),
		Error: &ErrorResponse{
			Code:    httpCode,
			Message: msg,
		},
	})
}

func UnauthorizedError(w http.ResponseWriter) {
	w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(JsonResponse{
		Message: "Error",
		Error: &ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constant.HTTPStatusText(http.StatusUnauthorized),
		},
	})
}
