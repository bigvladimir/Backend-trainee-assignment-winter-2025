package service_errors

import (
	"errors"
)

var ErrInvalidReq = errors.New("invalid request")
var ErrInvalidAuth = errors.New("invalid authentication")
