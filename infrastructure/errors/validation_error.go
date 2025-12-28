package errors

type TValidationError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func ValidationError(key string, message string) *TValidationError {
	return &TValidationError{
		Key:     key,
		Message: message,
	}
}
