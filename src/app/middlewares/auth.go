package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"wikilibras-core/src/app/exceptions"

	"github.com/dgrijalva/jwt-go"
)

// Authentication -
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tokenValid(r)
		if exceptions.HandlerErrors(
			err, w, "You must provide a valid authenticated access token", http.StatusUnauthorized,
		) {
			return
		}
		next(w, r)
	}
}

func tokenValid(r *http.Request) error {
	apiSecret := os.Getenv("API_SECRET")
	tokenString := extractToken(r)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
