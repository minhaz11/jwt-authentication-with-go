package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"time"
)

var jwtSecret = []byte("my_secret")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		slog.Info("Error decoding credentials: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		slog.Info("Invalid credentials")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    creds.Username,
			Subject:   "my-subject",
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		slog.Error("Error signing token ", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "_token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
	})

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_token")

	if errors.Is(err, http.ErrNoCookie) {
		slog.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		slog.Error("Invalid token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := cookie.Value

	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		fmt.Println("Invalid token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = w.Write([]byte("welcome"))
	if err != nil {
		return
	}
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {

}
