package token_manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var errUninitializedTokenManager = errors.New("token manager not initialized")
var errInvatilTTL = errors.New("invalid time to live param")

var tm tokenManagerSettings

type tokenManagerSettings struct {
	initialized bool

	secretKey []byte
	ttl       time.Duration
}

// Init необходимая инициализация токен менеджера
func Init(secretKey string, minutesTTL int) error {
	if minutesTTL < 1 {
		return errInvatilTTL
	}

	tm.secretKey = []byte(secretKey)
	tm.ttl = time.Duration(minutesTTL) * time.Minute

	tm.initialized = true

	return nil
}

func TokenManager() *tokenManagerSettings {
	if !tm.initialized {
		panic(errUninitializedTokenManager)
	}
	return &tm
}

func (t *tokenManagerSettings) CreateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(t.ttl).Unix(),
	})

	return token.SignedString(t.secretKey)
}

func (t *tokenManagerSettings) VerifyToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid crypto method")
		}
		return t.secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return int(userID), nil
		}
	}

	return 0, fmt.Errorf("invalid token")
}
