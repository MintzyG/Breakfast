package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	BFE "breakfast/_internal/errors"
	"breakfast/models"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			BFE.HandleError(w, BFE.New(BFE.ErrHeaderMissing, errors.New("MW: Missing Authorization header")))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			BFE.HandleError(w, BFE.New(BFE.ErrHeaderMalformed, errors.New("MW: Invalid/Malformed Authorization header, expected format: Bearer {token}")))
			return
		}

		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !token.Valid {
			BFE.HandleError(w, BFE.New(BFE.ErrUnauthorized, errors.New("MW: "+err.Error())))
			return
		}

		claims, ok := token.Claims.(*models.UserClaims)
		if !ok {
			BFE.HandleError(w, BFE.New(BFE.ErrClaims, errors.New("MW: Invalid token claims")))
			return
		}

		ctx := context.WithValue(r.Context(), models.ClaimsContext, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
