package model

type AuthRequest struct {
	// Имя пользователя для аутентификации.
	Username string `json:"username"`
	// Пароль для аутентификации.
	Password string `json:"password"`
}

type AuthResponse struct {
	// JWT-токен для доступа к защищенным ресурсам.
	Token string `json:"token,omitempty"`
}

type UserCreation struct {
	Username     string
	PasswordHash []byte
	Balance      int
}

type UserAuthInfo struct {
	UserID       int
	PasswordHash []byte
}
