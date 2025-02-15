package postgresql

import (
	"avito-shop-service/internal/app/repository/rep_errors"
	"errors"
)

func (s *ServiceStorage) IsNotFound(err error) bool {
	return errors.Is(err, rep_errors.ErrNotFound)
}
