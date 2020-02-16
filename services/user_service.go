package services

import (
	"net/http"

	"github.com/susilowibowo/my-microservices/domain/users"
	"github.com/susilowibowo/my-microservices/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// return nil, nil

	// return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
