package controller

import (
	service "food-reserve/logic/service/interface"
	"food-reserve/logic/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}
func (c *AuthController) RoleMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			ctx.Abort()
			return
		}

		claims, err := c.authService.CheckPermission(tokenStr, requiredPermission)

		if err != nil {
			if err.Error() == utils.Forbidden {
				ctx.JSON(http.StatusForbidden, gin.H{"error": utils.Forbidden})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)
		ctx.Next()
		return
	}
}
