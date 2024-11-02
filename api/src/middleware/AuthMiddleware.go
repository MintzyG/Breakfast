package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
			BFE.HandleError(w, BFE.New(BFE.ErrUnauthorized, err))
			return
		}

		claims, ok := token.Claims.(*models.UserClaims)
		if !ok {
			BFE.HandleError(w, BFE.New(BFE.ErrClaims, errors.New("MW: Invalid token claims")))
			return
		}

		id, err := uuid.Parse(claims.UserID)
		if err != nil {
			BFE.HandleError(w, BFE.New(BFE.ErrUnauthorized, fmt.Errorf("MW: Invalid uuid: %v", err.Error())))
			return
		}

		ok, _ = DB.IsUserValid(id)
		if !ok {
			BFE.HandleError(w, BFE.New(BFE.ErrUnprocessable, fmt.Errorf("MW: Invalid User uuid: %v", id)))
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
