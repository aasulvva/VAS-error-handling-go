package error_handling

type VASError struct {
	ErrorId          string  `json:"error_id"`
	ErrorName        string  `json:"error_name"`
	ErrorDescription *string `json:"error_description,omitempty"`
	StatusCode       uint    `json:"status_code"`
	GoError          error   `json:"go_error"`
}
