package router

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
