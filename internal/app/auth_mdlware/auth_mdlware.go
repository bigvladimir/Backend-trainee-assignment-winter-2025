package auth_mdlware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"avito-shop-service/internal/app/model"
	tm "avito-shop-service/internal/pkg/token_manager"
)

func AuthCheck(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		if len(authHeader) < 8 {
			unauthError(w, "invalid Authorization header")
			return
		}

		userID, err := tm.TokenManager().VerifyToken(authHeader[7:])
		if err != nil {
			unauthError(w, "invalid token")
			return
		}

		req.Header.Set("X-User-ID", fmt.Sprintf("%d", userID))

		handler.ServeHTTP(w, req)
	}
}

func unauthError(w http.ResponseWriter, errText string) {
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(model.ErrorResponse{
		Errors: errText,
	}); err != nil {
		log.Println("response error:", err)
	}
	return
}
