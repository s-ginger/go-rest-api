package auth

import (
	"net/http"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("super-secret-key")


type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Ok bool `json:"ok"`
	Token string `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	// TODO: сохранить пользователя в БД, захешировать пароль и т.д.
	// Для примера просто отдадим токен


	claims := jwt.MapClaims{
		"username": req.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate Token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RegisterResponse{Ok: true, Token: tokenString})
}