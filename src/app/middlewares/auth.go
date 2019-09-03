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
		_, err := decodeToken(r)
		if exceptions.HandlerErrors(
			err, w, "You must provide a valid authenticated access token", http.StatusUnauthorized,
		) {
			return
		}
		next(w, r)
	}
}

// GetUserAuth - s
func GetUserAuth(r *http.Request) (jwt.MapClaims, error) {
	// get claims
	claims, err := decodeToken(r)
	return claims, err
}

func decodeToken(r *http.Request) (jwt.MapClaims, error) {
	apiSecret := os.Getenv("API_SECRET")
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
