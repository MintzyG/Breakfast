package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Missing header", "HEADER_MISSING")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid header", "HEADER_MALFORMED")
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !token.Valid {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token", "UNAUTHORIZED_ACCESS")
			return
		}

		claims, ok := token.Claims.(*models.UserClaims)
		if !ok {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token claims", "UNAUTHORIZED_ACCESS")
			return
		}

		id, err := uuid.Parse(claims.UserID)
		if err != nil {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Invalid uuid: %v", err.Error()), "UNAUTHORIZED_ACCESS")
			return
		}

		ok, _ = DB.IsUserValid(id)
		if !ok {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid user", "UNAUTHORIZED_ACCESS")
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
