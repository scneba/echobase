package registering

type Error struct {
	Code  ErrorCode   //unique code for the error
	Field string      //the field which has the error
	Data  interface{} //for bulk registration, this is important to know the data with error.
}

func (e *Error) Error() string {
	return string(e.Code)
}

type ErrorCode string

const (
	EM_REQUIRED_FIELD      ErrorCode = "REQUIRED_FIELD"
	EM_VALIDATION_ERROR    ErrorCode = "VALIDATION_ERROR"
	EMAIL_VALIDATION_ERROR ErrorCode = "EMAIL_VALIDATION_ERROR"
	USER_VALIDATION_ERROR  ErrorCode = "USER_ALREADY_EXIST"
)
