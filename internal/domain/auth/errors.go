package auth

import "errors"

var (
	PassIncorrect = errors.New("password mismatch")
)
