package utilities

import (
	"context"

	"breakfast/internal/models"
)

func GetUserFromContext(ctx context.Context) *models.UserClaims {
	claims, ok := ctx.Value("user").(*models.UserClaims)
	if !ok {
		return nil
	}
	return claims
}
