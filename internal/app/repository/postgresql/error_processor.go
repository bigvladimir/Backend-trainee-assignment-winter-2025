package postgresql

import (
	"errors"

	"avito-shop-service/internal/app/repository/rep_errors"
)

func (s *ServiceStorage) IsNotFound(err error) bool {
	return errors.Is(err, rep_errors.ErrNotFound)
}
