package domain

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")

var ErrUserNotAuthorized = errors.New("user not authorized to this resource and or method")
