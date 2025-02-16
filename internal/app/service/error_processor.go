package service

import (
	"errors"

	"avito-shop-service/internal/app/service/service_errors"
)

func (s *Service) IsInvalidReq(err error) bool {
	return errors.Is(err, service_errors.ErrInvalidReq)
}

func (s *Service) IsInvalidAuth(err error) bool {
	return errors.Is(err, service_errors.ErrInvalidAuth)
}
