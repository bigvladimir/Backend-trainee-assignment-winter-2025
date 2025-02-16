package service

import (
	"bytes"
	"context"

	"avito-shop-service/internal/app/model"
	"avito-shop-service/internal/app/service/service_errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Auth(ctx context.Context, input model.AuthRequest) (model.AuthResponse, error) {
	if input.Username == "" || input.Password == "" {
		return model.AuthResponse{}, service_errors.ErrInvalidReq
	}

	inputHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.AuthResponse{}, err
	}

	var data model.UserAuthInfo
	data, err = s.stor.GetUserAuthInfoByUsername(ctx, input.Username)
	if err != nil && s.stor.IsNotFound(err) {
		if err := s.registerUser(ctx, model.UserCreation{
			Username:     input.Username,
			PasswordHash: inputHash,
			Balance:      s.settings.userStartBalance,
		}); err != nil {
			return model.AuthResponse{}, err
		}
	} else {
		return model.AuthResponse{}, err
	}

	if !bytes.Equal(inputHash, data.PasswordHash) {
		return model.AuthResponse{}, service_errors.ErrInvalidAuth
	}

	var token string
	token, err = s.tm.CreateToken(data.UserID)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{Token: token}, nil
}

// вообще по идее это должна быть отдельно используемая функция, но не хочется усложнять логику
func (s *Service) registerUser(ctx context.Context, input model.UserCreation) error {
	if err := s.stor.AddUser(ctx, input); err != nil {
		return err
	}
	return nil
}
