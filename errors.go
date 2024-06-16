package error_handling

import (
	"errors"
	"fmt"
	"net/http"
)

// Error IDs
const ERRID_RATE_LIMIT = "RATE_LIMITED"
const ERRID_RL_COOLDOWN = "RATE_LIMITED_COOLDOWN"
const ERRID_RL_LOGIN = "RATE_LIMITED_LOGIN"
const ERRID_PROCESSING = "PROCESSING_ERROR"
const ERRID_DECODING = "JSON_DECODING_ERROR"
const ERRID_UNAUTHORIZED = "UNAUTHORIZED"
const ERRID_UNSUPPORTED_METHOD = "UNSUPPORTED_METHOD"
const ERRID_INVALID_CREDENTIALS = "INVALID_CREDENTIALS"
const ERRID_MISSING_FIELD = "MISSING_FIELD"
const ERRID_INVALID_DATA = "INVALID_DATA"

// Error types
func RateLimitError(ip string) *VASError {
	err := VASError{
		ErrorId:          ERRID_RATE_LIMIT,
		ErrorName:        "You have been sending too many requests recently!",
		ErrorDescription: nil,
		StatusCode:       http.StatusTooManyRequests,
		GoError:          errors.New(fmt.Sprintf("ip %s sending too many requests", ip)),
	}
	return &err
}

func RateLimitCooldownError(ip string, cooldown uint, limiterName string) *VASError {
	desc := fmt.Sprintf("Please wait %d seconds before trying again", cooldown)
	err := VASError{
		ErrorId:          ERRID_RL_COOLDOWN,
		ErrorName:        "You have been sending too many requests!",
		ErrorDescription: &desc,
		StatusCode:       http.StatusTooManyRequests,
		GoError:          errors.New(fmt.Sprintf("ip %s sending too many requests, limited by %s for %d second(s)", ip, limiterName, cooldown)),
	}
	return &err
}

func RateLimitLoginError(ip string, userId uint, cooldown uint) *VASError {
	desc := fmt.Sprintf("Please wait %d seconds before trying again", cooldown)
	return &VASError{
		ErrorId:          ERRID_RL_LOGIN,
		ErrorName:        "You have been sending too many login requests!",
		ErrorDescription: &desc,
		StatusCode:       http.StatusTooManyRequests,
		GoError:          errors.New(fmt.Sprintf("too many login requests for user with ID %d coming from %s", userId, ip)),
	}
}

func ProcessingError(errType string, err error) *VASError {
	desc := "Please try again. If the error persists, please contact the administrators!"
	return &VASError{
		ErrorId:          ERRID_PROCESSING,
		ErrorName:        "An error occurred processing your request!",
		ErrorDescription: &desc,
		StatusCode:       http.StatusInternalServerError,
		GoError:          errors.New(fmt.Sprintf("[%s] %s", errType, err)),
	}
}

func DecodingError(err error) *VASError {
	desc := "Please check documentation for the correct JSON schema for the request"
	return &VASError{
		ErrorId:          ERRID_DECODING,
		ErrorName:        "An error occurred decoding JSON request data",
		ErrorDescription: &desc,
		StatusCode:       http.StatusBadRequest,
		GoError:          errors.New(fmt.Sprintf("JSON decoding: %s", err)),
	}
}

func UnauthorizedError(err error) *VASError {
	desc := "You lack the permission required to perform this action!"
	return &VASError{
		ErrorId:          ERRID_UNAUTHORIZED,
		ErrorName:        "Unauthorized access!",
		ErrorDescription: &desc,
		StatusCode:       http.StatusUnauthorized,
		GoError:          err,
	}
}

func UnsupportedMethodError(methodUsed string, allowedMethods []string) *VASError {
	desc := "Allowed methods: "
	for idx, method := range allowedMethods {
		if idx == len(allowedMethods)-1 {
			desc += method
		} else {
			desc += method + ", "
		}
	}
	return &VASError{
		ErrorId:          ERRID_UNSUPPORTED_METHOD,
		ErrorName:        fmt.Sprintf("Unsupported method: %s", methodUsed),
		ErrorDescription: &desc,
		StatusCode:       http.StatusMethodNotAllowed,
		GoError:          errors.New(fmt.Sprintf("client used an unsupported method, %s", methodUsed)),
	}
}

func InvalidCredentialsError(err error) *VASError {
	desc := "Wrong username/password combination"
	return &VASError{
		ErrorId:          ERRID_INVALID_CREDENTIALS,
		ErrorName:        "Invalid credentials",
		ErrorDescription: &desc,
		StatusCode:       http.StatusUnauthorized,
		GoError:          err,
	}
}
func MissingFieldError(field string) *VASError {
	desc := fmt.Sprintf("Applies to field '%s'", field)
	return &VASError{
		ErrorId:          ERRID_MISSING_FIELD,
		ErrorName:        "Missing field in input data found while processing your request",
		ErrorDescription: &desc,
		StatusCode:       http.StatusBadRequest,
		GoError:          nil,
	}
}
func InvalidDataError(field string, err error) *VASError {
	desc := fmt.Sprintf("Applies to field '%s'", field)
	return &VASError{
		ErrorId:          ERRID_INVALID_DATA,
		ErrorName:        "Invalid input data found while processing your request",
		ErrorDescription: &desc,
		StatusCode:       http.StatusNotAcceptable,
		GoError:          err,
	}
}
