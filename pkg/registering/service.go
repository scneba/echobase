package registering

import (
	"fmt"
	"net/mail"

	"github.com/google/uuid"
)

type RegisteringInterface interface {
	RegisterUser(user User) (id *uuid.UUID, errs []error)
}

type DatabaseRepository interface {
	RegisterUser(user User) (id uuid.UUID, err error)
	UserExists(user User) (unique bool, err error)
}

type service struct {
	db DatabaseRepository
}

func NewRegisteringService(db DatabaseRepository) RegisteringInterface {
	return &service{db: db}
}

func (s *service) RegisterUser(user User) (id *uuid.UUID, errs []error) {
	errs = s.validateUser(user)
	if errs != nil {
		fmt.Printf("Error while vaildating user %s: ", errs)
		return nil, errs
	}

	userID, err := s.db.RegisterUser(user)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	return &userID, nil
}

func (s *service) validateUser(user User) (errs []error) {
	if len(user.FirstName) == 0 {
		errs = append(errs, &Error{Field: "first_name", Code: EM_REQUIRED_FIELD, Data: user})
	}
	if len(user.LastName) == 0 {
		errs = append(errs, &Error{Field: "last_name", Code: EM_REQUIRED_FIELD, Data: user})
	}
	if len(user.Username) == 0 {
		errs = append(errs, &Error{Field: "user_name", Code: EM_REQUIRED_FIELD, Data: user})
	}
	if len(user.Email) == 0 {
		errs = append(errs, &Error{Field: "email", Code: EM_REQUIRED_FIELD, Data: user})
	}
	//TODO: what's the right phonenumber format
	if len(user.PhoneNumber) == 0 {
		errs = append(errs, &Error{Field: "phone_number", Code: EM_REQUIRED_FIELD, Data: user})
	}
	if !isEmailValid(user.Email) {
		errs = append(errs, &Error{Field: "email", Code: EMAIL_VALIDATION_ERROR, Data: user})
	}
	if unique, _ := s.db.UserExists(user); unique {
		errs = append(errs, &Error{Field: "user_name", Code: USER_VALIDATION_ERROR, Data: user})
	}
	return errs
}

func newUUID() uuid.UUID {
	id := uuid.New()
	return id
}
func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
