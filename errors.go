package error_handling

import (
	"errors"
	"fmt"
	"net/http"
)

// Error IDs
const ERRID_RATE_LIMIT = "RATE_LIMITED"
const ERRID_PROCESSING = "PROCESSING_ERROR"
const ERRID_DECODING = "JSON_DECODING_ERROR"
const ERRID_UNAUTHORIZED = "UNAUTHORIZED"

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
