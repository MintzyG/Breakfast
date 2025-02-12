package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"breakfast/config"
	"breakfast/internal/models"
	u "breakfast/internal/utilities"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var secretKey string = config.GetJWTSecret()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			u.Send(w, "mw-error: Authorization header is required", nil, http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			u.Send(w, "mw-error: Authorization header format must be Bearer {token}", nil, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				u.Send(w, "mw-error: Invalid signing method", nil, http.StatusUnauthorized)
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			u.Send(w, "mw-error: Invalid token:" + err.Error(), nil, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			u.Send(w, "mw-error: Token is not valid or has expired", nil, http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.UserClaims); ok {
			// Enhanced logging for claims
			_, err := json.Marshal(claims)
			if err != nil {
				u.Send(w, "mw-error: Error marshaling claims: " + err.Error(), nil, http.StatusUnauthorized)
			} // else {
			//   log.Println(claimsJSON)
			// }

			// Expiry check
			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
				u.Send(w, "mw-error: Token has expired", nil, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			u.Send(w, "mw-error: Invalid token claims", nil, http.StatusUnauthorized)
		}
	})
}
