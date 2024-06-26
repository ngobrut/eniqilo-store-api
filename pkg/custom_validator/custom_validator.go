package custom_validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/go-playground/locales/en"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
)

type ValidatorError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ValidatorError) Error() string {
	return e.Message
}

func ValidateStruct(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println("[error-parse-body]", err.Error())
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "please check your body request",
		})

		return err
	}

	validate := validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation("category", validateCategory)
	validate.RegisterValidation("phoneCode", validatePhoneNumber)
	validate.RegisterValidation("validUrl", validateURL)

	err = validate.Struct(data)
	if err == nil {
		return nil
	}

	var message string
	var details = make([]string, 0)
	for _, field := range err.(validator.ValidationErrors) {
		message = field.Translate(trans)

		switch field.Tag() {
		case "category":
			message = fmt.Sprintf("%s must be one of [%s]", field.Field(), strings.Join(constant.Categories, ", "))
		case "phoneCode":
			message = "should start with `+` and international calling codes"
		case "validUrl":
			message = "should be url"
		}

		details = append(details, message)
	}

	err = ValidatorError{
		Code:    http.StatusBadRequest,
		Message: message,
		Details: details,
	}

	return err
}

func validateCategory(fl validator.FieldLevel) bool {
	return constant.ValidCategory[fl.Field().String()]
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	for _, code := range constant.CountryCode {
		match, _ := regexp.MatchString("^"+regexp.QuoteMeta(code)+`\d+$`, fl.Field().String())
		if match {
			return true
		}
	}
	return false
}

func validateURL(fl validator.FieldLevel) bool {
	parsedURL, err := url.Parse(fl.Field().String())
	if err != nil {
		return false
	}

	// Check if the scheme is present and it's http or https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Check if the host is present and it has a valid format
	if parsedURL.Host == "" {
		return false
	}

	// Check if the host has a valid domain format
	parts := strings.Split(parsedURL.Host, ".")
	if len(parts) < 2 {
		return false
	}

	// Check if the path, if present, is in a valid format
	if parsedURL.Path != "" && !strings.HasPrefix(parsedURL.Path, "/") {
		return false
	}

	// All checks passed, URL is valid
	return true
}
