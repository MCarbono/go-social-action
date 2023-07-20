package appError

type FieldsValidationError struct {
	Message string `json:"err"`
}

func (e FieldsValidationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string `json:"err"`
}

func (e NotFoundError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string `json:"err"`
}

func (e InternalServerError) Error() string {
	return e.Message
}

func NewInternalServerError() error {
	return InternalServerError{
		Message: `We could not process your request because of some internal problem. 
			Please try again in a few minutes and if the problem persists, contact the technical support`,
	}
}
