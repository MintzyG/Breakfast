package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"breakfast/config"
	"breakfast/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var secretKey string = config.GetJWTSecret()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Authorization header is required"}`, http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error": "Authorization header format must be Bearer {token}"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Invalid signing method")
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			http.Error(w, `{"error": "Invalid token: `+err.Error()+`"}`, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, `{"error": "Token is not valid or has expired"}`, http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.UserClaims); ok {
			// Enhanced logging for claims
			claimsJSON, err := json.Marshal(claims)
			if err != nil {
				log.Println("Error marshaling claims:", err)
			} else {
				log.Println("Token Claims (JSON):", string(claimsJSON))
			}

			// Expiry check
			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
				http.Error(w, `{"error": "Token has expired"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, `{"error": "Invalid token claims"}`, http.StatusUnauthorized)
		}
	})
}

